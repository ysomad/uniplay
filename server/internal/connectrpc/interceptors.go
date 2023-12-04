package connectrpc

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	kratos "github.com/ory/kratos-client-go"

	"github.com/ysomad/uniplay/server/internal/kratosctx"
)

var errIdentityNotMatch = errors.New("session identity not match")

// sessionCookie returns kratos session cookie from header Cookie
// TODO: WRITE TESTS
func sessionCookie(h http.Header) (string, error) {
	cookieHdr := h.Get("Cookie")
	if cookieHdr == "" {
		return "", errors.New("not found cookie header")
	}

	for _, cookie := range strings.Split(cookieHdr, "; ") {
		parts := strings.SplitN(cookie, "=", 2)

		slog.Debug("cookie parts", "val", parts)

		if len(parts) == 2 && parts[0] == "ory_kratos_session" {
			if cookie == "" {
				return "", errors.New("empty session cookie")
			}
			return cookie, nil
		}
	}

	return "", errors.New("session cookie not found")
}

func newAuthInterceptor(client *kratos.APIClient, orgSchemaID string) connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			cookie, err := sessionCookie(req.Header())
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			session, resp, err := client.FrontendApi.
				ToSession(ctx).
				Cookie(cookie).
				Execute()
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			if resp.StatusCode != http.StatusOK {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			if !session.GetActive() {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			identity := session.GetIdentity()

			if identity.SchemaId != orgSchemaID {
				return nil, connect.NewError(connect.CodePermissionDenied,
					fmt.Errorf("%w, must be %s", errIdentityNotMatch, orgSchemaID))
			}

			ctx = kratosctx.WithIdentityID(ctx, identity.Id)

			return next(ctx, req)
		})
	})
}
