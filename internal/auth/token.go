package auth

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
)

// HeaderTokenDefault default token header name
const HeaderTokenDefault = "x-internal-token"

// Errors returned from meta information
var (
	ErrHeader        = errors.New("header could not be found, incorrect or multiple")
	ErrEmptyMetadata = errors.New("metadata doesn't exist")
)

// AddNamedTokenToContext puts the token in a custom header
func AddNamedTokenToContext(ctx context.Context, headerName, token string) context.Context {
	md, mdFound := metadata.FromOutgoingContext(ctx)
	if !mdFound {
		return metadata.AppendToOutgoingContext(ctx, headerName, token)
	}

	headerName = strings.ToLower(headerName)

	mdCp := md.Copy()
	mdCp.Set(headerName, token)

	return metadata.NewOutgoingContext(ctx, mdCp)
}

// DeleteHeaderValue removes the headerName from the input metadata in the context
func DeleteHeaderValue(ctx context.Context, headerName string) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	mdCp := md.Copy()
	headerName = strings.ToLower(headerName)
	delete(mdCp, headerName)

	return metadata.NewIncomingContext(ctx, mdCp)
}

// getHeaderValue retrieves the value of headerName from the context
func getHeaderValue(ctx context.Context, headerName string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrEmptyMetadata
	}
	headers := md.Get(headerName)
	if len(headers) != 1 {
		return "", ErrHeader
	}

	return headers[0], nil
}
