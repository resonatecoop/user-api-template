package iamserver

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/merefield/grpc-user-api/pkg/access"
	jwt "github.com/merefield/grpc-user-api/pkg/jwt"
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
		//	var i interface{}
		var err error

		start := time.Now()

		// Logging with grpclog (grpclog.LoggerV2)
		grpclog.Infof("Request - Method:%s\tDuration:0\tError:%v\n",
			info.FullMethod,
			err)

		no_token_methods_array := strings.Split(interceptor.acc.NoTokenMethods, "\\|")

		isPresent := stringInSlice(info.FullMethod, no_token_methods_array)

		// Skip authorize when GetJWT is requested
		//	if info.FullMethod != "/iam.ResonateIAM/Auth" {
		if !isPresent {
			grpclog.Infof("Expecting JWT, let's check ...\tError:%v\n",
				err)
			if err := interceptor.authorize(ctx, info.FullMethod); err != nil {
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

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	//accessibleRoles, ok := interceptor.accessibleRoles[method]
	// if !ok {
	// 	// everyone can access
	// 	return nil
	// }

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	//claims, err := interceptor.jwt.ParseToken(accessToken)
	_, err := interceptor.jwt.ParseToken(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	// for _, role := range accessibleRoles {
	// 	if role == claims.Role {
	// 		return nil
	// 	}
	// }
	return nil
	//return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
