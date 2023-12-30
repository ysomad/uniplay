package kratosx

import (
	"context"
)

type identityIDKey struct{}

// IdentityID returns kratos identity id from context, returns empty string if not found.
func IdentityID(ctx context.Context) string {
	e, _ := ctx.Value(identityIDKey{}).(string) //nolint:errcheck // anyway have to check empty string
	return e
}

// WithIdentityID adds identity id to context.
func WithIdentityID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, identityIDKey{}, id)
}
