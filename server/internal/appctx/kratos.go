package appctx

import (
	"context"
)

type identityIDKey struct{}

func IdentityID(ctx context.Context) (string, bool) {
	e, ok := ctx.Value(identityIDKey{}).(string)
	return e, ok
}

func WithIdentityID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, identityIDKey{}, id)
}
