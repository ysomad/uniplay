package interceptor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"connectrpc.com/connect"
	ory "github.com/ory/client-go"
	"github.com/ysomad/uniplay/server/internal/appctx"
	"github.com/ysomad/uniplay/server/internal/config"
)

var (
	errIdentityNotMatch = errors.New("session identity not match")
	errUnauthenticated  = errors.New("unauthenticated")
)

// sessionCookie returns kratos session cookie from header Cookie
func sessionCookie(h http.Header) (string, error) {
	cookieHdr := h.Get("Cookie")
	if cookieHdr == "" {
		return "", errors.New("not found cookie header")
	}

	for _, cookie := range strings.Split(cookieHdr, "; ") {
		parts := strings.SplitN(cookie, "=", 2)
		if len(parts) == 2 && parts[0] == "ory_kratos_session" {
			if cookie == "" {
				return "", errors.New("empty session cookie")
			}
			return cookie, nil
		}
	}

	return "", errors.New("session cookie not found")
}

func NewAuth(kratos *ory.APIClient, conf config.Kratos) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			cookie, err := sessionCookie(req.Header())
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}

			session, resp, err := kratos.FrontendAPI.
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

			if identity.SchemaId != conf.OrganizerSchemaID {
				return nil, connect.NewError(connect.CodePermissionDenied,
					fmt.Errorf("%w, must be %s", errIdentityNotMatch, conf.OrganizerSchemaID))
			}

			ctx = appctx.WithIdentityID(ctx, identity.Id)

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}
