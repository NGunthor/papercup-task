package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func TestAddNamedTokenToContext(t *testing.T) {
	const (
		validHeaderName = "x-token"
		validTokenName  = "x-secret"
	)

	var tests = []struct {
		name       string
		inMetadata func() context.Context
		want       []string
		inHeader   string
		inToken    string
	}{
		{
			name: "regular case - pass 1 arg, expects to receive a token",
			inMetadata: func() context.Context {
				return metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{}))
			},
			want:     []string{validTokenName},
			inHeader: validHeaderName,
			inToken:  validTokenName,
		},
		{
			name: "2",
			inMetadata: func() context.Context {
				return metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
					validHeaderName: "fake-token",
				}))
			},
			want:     []string{validTokenName},
			inHeader: validHeaderName,
			inToken:  validTokenName,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			// ACT
			ctx := tc.inMetadata()
			ctx = AddNamedTokenToContext(ctx, tc.inHeader, tc.inToken)

			// ASSERT
			md, mdFound := metadata.FromOutgoingContext(ctx)
			r.True(mdFound)

			r.Greater(md.Len(), 0)
			got := md.Get(tc.inHeader)
			r.Equal(tc.want, got)
		})
	}
}
