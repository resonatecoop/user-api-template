package authorization

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	uuid "github.com/google/uuid"
	"github.com/resonatecoop/user-api-template/model"
	"github.com/resonatecoop/user-api-template/pkg/access"
	uuidpkg "github.com/resonatecoop/user-api-template/pkg/uuid"
	pbUser "github.com/resonatecoop/user-api-template/proto/user"
	grpclog "google.golang.org/grpc/grpclog"
)

var (
	// ErrAccessTokenNotFound ...
	ErrAccessTokenNotFound = errors.New("Access token not found")
	// ErrAccessTokenExpired ...
	ErrAccessTokenExpired = errors.New("Access token expired")
)

type AuthInterceptor struct {
	db  *bun.DB
	rf  int
	acc *access.AccessConfig
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func NewAuthInterceptor(db *bun.DB, rf int, acc *access.AccessConfig) *AuthInterceptor {
	return &AuthInterceptor{db, rf, acc}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		var err error

		start := time.Now()

		// Logging with grpclog (grpclog.LoggerV2)

		args := []string{
			"[user-api-auth]",
			//formatOperation(event),
			fmt.Sprintf("Request - Method:%s\tDuration:0\tError:%v\n", info.FullMethod, err),
		}

		// if event.Err != nil {
		// 	typ := reflect.TypeOf(event.Err).String()
		// 	args = append(args,
		// 		"\t",
		// 		color.New(color.BgRed).Sprintf(" %s ", typ+": "+event.Err.Error()),
		// 	)
		// }

		//fmt.Println(args...)

		grpclog.Infof(strings.Join(args, ""))
		// grpclog.Infof("Request - Method:%s\tDuration:0\tError:%v\n",
		// info.FullMethod,
		// err)

		NoTokenMethods := strings.Split(interceptor.acc.NoTokenMethods, ",")

		TokenRequired := !stringInSlice(info.FullMethod, NoTokenMethods)

		// Skip authorize when configured methods are requested
		//	eg if requesting token
		if TokenRequired {
			grpclog.Infof("Expecting AccessToken, let's check ...")
			err := interceptor.authorize(ctx, req, info.FullMethod)
			if err != nil {
				grpclog.Infof("Request Denied - Method:%s\tDuration:%s\tError:%v\n",
					info.FullMethod,
					time.Since(start),
					err)
				return nil, err
			}
		}

		// Calls the handler
		h, err := handler(ctx, req)

		grpclog.Infof("Request Authorised - Method:%s\tDuration:%s\tError:%v\n",
			info.FullMethod,
			time.Since(start),
			err)

		return h, err
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, req interface{}, method string) error {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	PublicMethods := strings.Split(interceptor.acc.PublicMethods, ",")
	WriteMethods := strings.Split(interceptor.acc.WriteMethods, ",")

	isPublicAccessMethod := stringInSlice(method, PublicMethods)

	accessTokenSource := strings.Split(values[0], " ")

	if len(accessTokenSource) != 2 {
		return status.Errorf(codes.PermissionDenied, "incorrect authorization header format")
	}

	accessToken := accessTokenSource[1]

	accessTokenRecord, err := interceptor.Authenticate(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	scopes := strings.Split(accessTokenRecord.Scope, " ")

	// determine if the request has write permission
	_, read_write := interceptor.find(scopes, "read_write")

	// leave now if no write permission
	if !read_write && stringInSlice(method, WriteMethods) {
		return status.Errorf(codes.PermissionDenied, "attempt to write to user-api without write scope")
	}

	scopes = interceptor.delete(scopes, "read_write")

	scopes = interceptor.delete(scopes, "read")

	// assume role is remaining scope element
	tokenRole := scopes[0]

	tokenRoleRow := new(model.Role)

	err = interceptor.db.NewSelect().
		Model(tokenRoleRow).
		Where("name = ?", tokenRole).
		Scan(ctx)

	if err != nil {
		return status.Errorf(codes.PermissionDenied, "problem determining role from token")
	}

	tokenRoleValue := tokenRoleRow.ID

	user := new(model.User)

	err = interceptor.db.NewSelect().
		Model(user).
		Where("id = ?", accessTokenRecord.UserID).
		Scan(ctx)

	if err != nil {
		return status.Errorf(codes.PermissionDenied, "problem determining user role")
	}

	userRoleValue := user.RoleID

	var activeRole int32

	if userRoleValue > tokenRoleValue {
		activeRole = userRoleValue
	} else {
		activeRole = tokenRoleValue
	}

	if isPublicAccessMethod {
		// everyone can access but check it's against their own ID
		if activeRole > int32(model.LabelRole) {
			id, err := interceptor.extractUserIdFromReq(ctx, req, accessTokenRecord)

			if err != nil {
				return err
			}

			ID, err := uuid.Parse(id)

			if err != nil {
				return status.Errorf(codes.PermissionDenied, "UUID in request is not valid")
			}

			if ID != user.ID {
				return status.Errorf(codes.PermissionDenied, "requestor is not authorized to take action on another user record")
			}
			// must be working on their own record
		}
		return nil
	}

	// If not an admin, you can't access the remaining non-public methods
	if activeRole > int32(model.TenantAdminRole) {
		return status.Errorf(codes.PermissionDenied, "requestor is not authorized for this method")
	}

	// else all is fine at this gate at least, go ahead
	return nil
}

// Authenticate checks the access token is valid
func (interceptor *AuthInterceptor) Authenticate(token string) (*model.AccessToken, error) {
	// Fetch the access token from the database
	ctx := context.Background()
	accessToken := new(model.AccessToken)

	err := interceptor.db.NewSelect().
		Model(accessToken).
		Where("token = ?", token).
		Limit(1).
		Scan(ctx)

	// Not found
	if err != nil {
		return nil, ErrAccessTokenNotFound
	}

	// Check the access token hasn't expired
	if time.Now().UTC().After(accessToken.ExpiresAt) {
		return nil, ErrAccessTokenExpired
	}

	// Extend refresh token expiration database

	increasedExpiresAt := time.Now().Add(
		time.Duration(interceptor.rf) * time.Second,
	)

	//var res sql.Result

	//	err = GetOrCreateRefreshToken

	if uuidpkg.IsValidUUID(accessToken.UserID.String()) && accessToken.UserID != uuid.Nil {
		_, err = interceptor.db.NewUpdate().
			Model(new(model.RefreshToken)).
			Set("expires_at = ?", increasedExpiresAt).
			Set("updated_at = ?", time.Now().UTC()).
			Where("client_id = ?", accessToken.ClientID.String()).
			Where("user_id = ?", accessToken.UserID.String()).
			Exec(ctx)
	} else {
		_, err = interceptor.db.NewUpdate().
			Model(new(model.RefreshToken)).
			Set("expires_at = ?", increasedExpiresAt).
			Set("updated_at = ?", time.Now().UTC()).
			Where("client_id = ?", accessToken.ClientID.String()).
			Where("user_id = uuid_nil()").
			Exec(ctx)
	}

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (interceptor *AuthInterceptor) extractUserIdFromReq(ctx context.Context, req interface{}, accessTokenRecord *model.AccessToken) (string, error) {
	// Dealing with normal users, Label admins can maintain own artist content.
	// attempt to extract the Id from all the possible request types dealing with failure
	userReq, ok := req.(*pbUser.UserRequest)

	if ok {
		return userReq.Id, nil
	}

	userUpdateReq, ok := req.(*pbUser.UserUpdateRequest)

	if ok {
		return userUpdateReq.Id, nil
	}

	userGroupCreateReq, ok := req.(*pbUser.UserGroupCreateRequest)

	if ok {
		return userGroupCreateReq.Id, nil
	}

	userGroupUpdateReq, ok := req.(*pbUser.UserGroupUpdateRequest)

	if ok {
		newUserGroup := new(model.UserGroup)

		err := interceptor.db.NewSelect().
			Model(newUserGroup).
			Where("owner_id = ?", accessTokenRecord.UserID).
			Where("id = ?", userGroupUpdateReq.Id).
			Scan(ctx)

		if err != nil {
			return "", status.Errorf(codes.PermissionDenied, "Supplied UUID for User Group is not valid or logged in User doesn't own Group")
		}

		return accessTokenRecord.UserID.String(), nil
	}

	return "", status.Errorf(codes.PermissionDenied, "UUID in request is not valid")
}

func (interceptor *AuthInterceptor) find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func (interceptor *AuthInterceptor) delete(slice []string, val string) []string {
	var location int

	for i, item := range slice {
		if item == val {
			location = i
			slice = append(slice[:location], slice[location+1:]...)
		}
	}

	return slice
}
