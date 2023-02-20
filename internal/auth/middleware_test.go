package auth

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	tokenHeader      = "token"
	unaryHandlerFunc = func(_ context.Context, _ interface{}) (interface{}, error) {
		return nil, nil
	}
)

func TestAuthMiddleware_Authorize(t *testing.T) {
	r := require.New(t)

	testCases := []struct {
		name                 string
		context              context.Context
		verifyHandleParamsFn func(ctx context.Context, req interface{}) (interface{}, error)
		secretTokenMap       map[string]string
		expErr               string
		expErrCode           codes.Code
	}{
		{
			name: "success: request authorized",
			context: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string]string{
					serviceNameHeader: "service1",
					tokenHeader:       "123456",
				})),
			verifyHandleParamsFn: func(ctx context.Context, req interface{}) (interface{}, error) {
				wantMD := metadata.New(map[string]string{
					serviceNameHeader: "service1",
					tokenHeader:       "123456",
				})
				gotMD, mdFound := metadata.FromIncomingContext(ctx)
				r.True(mdFound)

				r.Equal(wantMD, gotMD)
				return nil, nil
			},
			secretTokenMap: map[string]string{
				"service1": "123456",
				"service2": "qwerty",
			},
		},
		{
			name: "success: empty secret token map",
			context: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string]string{
					serviceNameHeader: "service1",
					tokenHeader:       "123456",
				})),
			verifyHandleParamsFn: func(ctx context.Context, req interface{}) (interface{}, error) {
				wantMD := metadata.New(map[string]string{
					serviceNameHeader: "service1",
					tokenHeader:       "123456",
				})
				gotMD, mdFound := metadata.FromIncomingContext(ctx)
				r.True(mdFound)

				r.Equal(wantMD, gotMD)
				return nil, nil
			},
			secretTokenMap: map[string]string{},
		},
		{
			name: "success: registered service without token",
			context: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string]string{
					serviceNameHeader: "service1",
					tokenHeader:       "",
				})),
			verifyHandleParamsFn: func(ctx context.Context, req interface{}) (interface{}, error) {
				wantMD := metadata.New(map[string]string{
					serviceNameHeader: "service1",
					tokenHeader:       "",
				})
				gotMD, mdFound := metadata.FromIncomingContext(ctx)
				r.True(mdFound)

				r.Equal(wantMD, gotMD)
				return nil, nil
			},
			secretTokenMap: map[string]string{
				"service1": "",
				"service2": "",
			},
		},
		{
			name: "failure: service name is not registered",
			context: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string]string{
					serviceNameHeader: "service3",
					tokenHeader:       "123456",
				})),
			expErr:     "rpc error: code = PermissionDenied desc = unknown service",
			expErrCode: codes.PermissionDenied,
			secretTokenMap: map[string]string{
				"service1": "123456",
				"service2": "qwerty",
			},
		},
		{
			name: "failure: invalid token",
			context: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string]string{
					serviceNameHeader: "service1",
					tokenHeader:       "qwerty",
				})),
			expErr:     "rpc error: code = Unauthenticated desc = invalid auth token",
			expErrCode: codes.Unauthenticated,
			secretTokenMap: map[string]string{
				"service1": "123456",
				"service2": "qwerty",
			},
		},
		{
			name: "failure: invalid token header",
			context: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string]string{
					serviceNameHeader:         "service1",
					"incorrect" + tokenHeader: "123456",
				})),
			expErr:     "rpc error: code = InvalidArgument desc = cannot get token header (token): header could not be found, incorrect or multiple",
			expErrCode: codes.InvalidArgument,
			secretTokenMap: map[string]string{
				"service1": "123456",
				"service2": "qwerty",
			},
		},
		{
			name:       "failure: empty metadata(incorrect IncomingContext)",
			context:    context.Background(),
			expErr:     "rpc error: code = InvalidArgument desc = cannot get app name header (x-service-name): metadata doesn't exist",
			expErrCode: codes.InvalidArgument,
			secretTokenMap: map[string]string{
				"service1": "123456",
				"service2": "qwerty",
			},
		},
		{
			name: "failure: invalid service name header",
			context: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string]string{
					"incorrect" + serviceNameHeader: "service1",
					tokenHeader:                     "123456",
				})),
			expErr:     "rpc error: code = InvalidArgument desc = cannot get app name header (x-service-name): header could not be found, incorrect or multiple",
			expErrCode: codes.InvalidArgument,
			secretTokenMap: map[string]string{
				"service1": "123456",
				"service2": "qwerty",
			},
		},
		{
			name: "failure: multiple values of service name header",
			context: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string]string{
					serviceNameHeader:                  "service1",
					strings.ToUpper(serviceNameHeader): "service1",
					tokenHeader:                        "123456",
				})),
			expErr:     "rpc error: code = InvalidArgument desc = cannot get app name header (x-service-name): header could not be found, incorrect or multiple",
			expErrCode: codes.InvalidArgument,
			secretTokenMap: map[string]string{
				"service1": "123456",
				"service2": "qwerty",
			},
		},
		{
			name: "failure: multiple values of token header",
			context: metadata.NewIncomingContext(context.Background(),
				metadata.New(map[string]string{
					serviceNameHeader:            "service1",
					tokenHeader:                  "123456",
					strings.ToUpper(tokenHeader): "123456",
				})),
			expErr:     "rpc error: code = InvalidArgument desc = cannot get token header (token): header could not be found, incorrect or multiple",
			expErrCode: codes.InvalidArgument,
			secretTokenMap: map[string]string{
				"service1": "123456",
				"service2": "qwerty",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			handleFn := unaryHandlerFunc
			if tc.verifyHandleParamsFn != nil {
				handleFn = tc.verifyHandleParamsFn
			}

			var actualErr error
			_, actualErr = NewTokenAuthServerInterceptor(tokenHeader, tc.secretTokenMap)(tc.context, struct{}{}, &grpc.UnaryServerInfo{}, handleFn)
			if tc.expErr != "" {
				r.EqualError(actualErr, tc.expErr)
				r.Equal(tc.expErrCode, status.Code(actualErr))
			} else {
				r.NoError(actualErr)
			}
		})
	}
}
