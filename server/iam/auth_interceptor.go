package iamserver

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	uuid "github.com/google/uuid"
	"github.com/resonatecoop/user-api/model"
	"github.com/resonatecoop/user-api/pkg/access"
	jwt "github.com/resonatecoop/user-api/pkg/jwt"
	pbUser "github.com/resonatecoop/user-api/proto/user"
	grpclog "google.golang.org/grpc/grpclog"

	"github.com/fatih/color"
)

type AuthInterceptor struct {
	jwt *jwt.JWT
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

func NewAuthInterceptor(jwt *jwt.JWT, acc *access.AccessConfig) *AuthInterceptor {
	return &AuthInterceptor{jwt, acc}
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
			grpclog.Infof("Expecting JWT, let's check ...")
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

	PublicUserMethods := strings.Split(interceptor.acc.PublicUserMethods, ",")
	isPublicAccessMethod := stringInSlice(method, PublicUserMethods)

	accessToken := values[0]

	claims, err := interceptor.jwt.ParseToken(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if isPublicAccessMethod {
		// everyone can access but check it's against their own ID
		if claims.Role > model.LabelRole {
			// dealing with normal users, Label admins can maintain own artist content.
			// attempt to extract the Id from all the possible request types dealing with failure
			// TODO a bit nasty, can we make this more elegant and less opinionated?
			var id string
			userReq, ok := req.(*pbUser.UserRequest)
			if !ok {
				userUpdateReq, ok := req.(*pbUser.UpdateUserRequest)
				if ok {
					id = userUpdateReq.Id
				} else {
					return status.Errorf(codes.PermissionDenied, "UUID in request is not valid")
				}
			} else {
				id = userReq.Id
			}

			ID, err := uuid.Parse(id)

			if err != nil {
				return status.Errorf(codes.PermissionDenied, "UUID in request is not valid")
			}

			if ID != claims.ID {
				return status.Errorf(codes.PermissionDenied, "requestor is not authorized to take action on another user record")
			}
			// must be working on their own record
		}
		return nil
	}

	// If not an admin, you can't access the remaining non-public methods
	if claims.Role > model.TenantAdminRole {
		return status.Errorf(codes.PermissionDenied, "requestor is not authorized for this method")
	}

	// else all is fine at this gate at least, go ahead
	return nil
}

// func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
// 	if !h.verbose {
// 		switch event.Err {
// 		case nil, sql.ErrNoRows:
// 			return
// 		}
// 	}

// 	now := time.Now()
// 	dur := now.Sub(event.StartTime)

// 	args := []interface{}{
// 		"[bun]",
// 		now.Format(" 15:04:05.000 "),
// 		formatOperation(event),
// 		fmt.Sprintf(" %10s ", dur.Round(time.Microsecond)),
// 		string(event.Query),
// 	}

// 	if event.Err != nil {
// 		typ := reflect.TypeOf(event.Err).String()
// 		args = append(args,
// 			"\t",
// 			color.New(color.BgRed).Sprintf(" %s ", typ+": "+event.Err.Error()),
// 		)
// 	}

// 	fmt.Println(args...)
// }

// func formatOperation(event *bun.QueryEvent) string {
// 	operation := eventOperation(event)
// 	return operationColor(operation).Sprintf(" %-16s ", operation)
// }

// func eventOperation(event *bun.QueryEvent) string {
// 	switch event.QueryAppender.(type) {
// 	case *bun.SelectQuery:
// 		return "SELECT"
// 	case *bun.InsertQuery:
// 		return "INSERT"
// 	case *bun.UpdateQuery:
// 		return "UPDATE"
// 	case *bun.DeleteQuery:
// 		return "DELETE"
// 	case *bun.CreateTableQuery:
// 		return "CREATE TABLE"
// 	case *bun.DropTableQuery:
// 		return "DROP TABLE"
// 	}
// 	return queryOperation(event.Query)
// }

// func queryOperation(name []byte) string {
// 	if idx := bytes.IndexByte(name, ' '); idx > 0 {
// 		name = name[:idx]
// 	}
// 	if len(name) > 16 {
// 		name = name[:16]
// 	}
// 	return string(name)
// }

func operationColor(operation string) *color.Color {
	switch operation {
	case "SELECT":
		return color.New(color.BgGreen)
	case "INSERT":
		return color.New(color.BgBlue)
	case "UPDATE":
		return color.New(color.BgYellow)
	case "DELETE":
		return color.New(color.BgRed)
	default:
		return color.New(color.FgBlack, color.BgWhite)
	}
}
