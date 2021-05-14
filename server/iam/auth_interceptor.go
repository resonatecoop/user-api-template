package iamserver

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/resonatecoop/user-api/pkg/access"
	jwt "github.com/resonatecoop/user-api/pkg/jwt"
	uuid "github.com/resonatecoop/user-api/pkg/uuid"
	pbUser "github.com/resonatecoop/user-api/proto/user"
	grpclog "google.golang.org/grpc/grpclog"
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
		grpclog.Infof("Request - Method:%s\tDuration:0\tError:%v\n",
			info.FullMethod,
			err)

		NoTokenMethods := strings.Split(interceptor.acc.NoTokenMethods, "\\|")

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
		if claims.Role > 3 {
			// dealing with normal users
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

			ID, err := uuid.GetUUIDFromString(id)

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
	if claims.Role > 3 {
		return status.Errorf(codes.PermissionDenied, "requestor is not authorized for this method")
	}

	// else all is fine at this gate at least, go ahead
	return nil
}
