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

		NoTokenMethods := strings.Split(interceptor.acc.NoTokenMethods, "\\|")

		TokenRequired := !stringInSlice(info.FullMethod, NoTokenMethods)

		// Skip authorize when configured methods are requested
		//	eg if requesting token
		if TokenRequired {
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

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	PublicUserMethods := strings.Split(interceptor.acc.PublicUserMethods, "\\|")
	isPublicAccessMethod := stringInSlice(method, PublicUserMethods)

	if isPublicAccessMethod {
		// everyone can access
		return nil
	}

	accessToken := values[0]

	claims, err := interceptor.jwt.ParseToken(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	// If not an admin, you can't access the remaining non-public methods
	if claims.Role > 3 {
		return status.Errorf(codes.PermissionDenied, "requestor is not authorized for this method")
	}

	// else all is fine at this gate at least, go ahead
	return nil
}
