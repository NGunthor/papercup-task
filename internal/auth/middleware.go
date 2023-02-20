package auth

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	serviceNameHeader = "x-service-name"
)

type (
	Tokens map[string]string
)

// NewTokenAuthServerInterceptor server-side interceptor that performs service authorization by IncomingContext(ctx)
// returns 403(http) if the serviceName from the header is not registered in the token map
// returns 401(http) if the header token does not match the one in the token map
// returns 400(http) if any of the headers could not be read
func NewTokenAuthServerInterceptor(tokenHeaderName string, authTokens Tokens) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		if len(authTokens) == 0 {
			return handler(ctx, req)
		}

		serviceName, err := getHeaderValue(ctx, serviceNameHeader)
		if err != nil {
			return resp, status.Errorf(codes.InvalidArgument, "cannot get app name header (%s): %v", serviceNameHeader, err)
		}

		secretToken, isPermittedService := authTokens[serviceName]
		if !isPermittedService {
			return resp, status.Error(codes.PermissionDenied, "unknown service")
		}

		serviceToken, err := getHeaderValue(ctx, tokenHeaderName)
		if err != nil {
			return resp, status.Errorf(codes.InvalidArgument, "cannot get token header (%s): %v", tokenHeaderName, err)
		}
		// to prevent the transfer of secrets to 3rd services
		ctx = DeleteHeaderValue(ctx, HeaderTokenDefault)

		if serviceToken != secretToken {
			return resp, status.Error(codes.Unauthenticated, "invalid auth token")
		}

		return handler(ctx, req)
	}
}

// NewTokenAuthClientInterceptor provides a client-side interceptor
// puts authToken in tokenHeaderName header
func NewTokenAuthClientInterceptor(tokenHeaderName, authToken string) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx = AddNamedTokenToContext(ctx, tokenHeaderName, authToken)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
