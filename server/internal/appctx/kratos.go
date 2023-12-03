package appctx

import (
	"context"
)

type identityIDKey struct{}

func IdentityID(ctx context.Context) string {
	e, _ := ctx.Value(identityIDKey{}).(string)
	return e
}

func WithIdentityID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, identityIDKey{}, id)
}
