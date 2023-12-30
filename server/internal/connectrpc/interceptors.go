package connectrpc

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	kratos "github.com/ory/kratos-client-go"

	"github.com/ysomad/uniplay/server/internal/kratosx"
)

var errKratosUnsucessfulResponse = errors.New("identity service unsuccessful response")

// sessionCookie returns kratos session cookie from header Cookie.
// TODO: WRITE TESTS.
func sessionCookie(h http.Header) (string, error) {
	req := http.Request{Header: h}

	cookie, err := req.Cookie(kratosx.SessionCookie)
	if err != nil {
		return "", err
	}

	return cookie.String(), nil
}

// newOrganizerInterceptor creates connect interceptor which checking current user against organizer schema id.
func newOrganizerInterceptor(client *kratos.APIClient, orgSchemaID string) connect.UnaryInterceptorFunc {
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
				slog.Error("kratos request error", "error", err)
				return nil, connect.NewError(connect.CodeUnauthenticated, errKratosUnsucessfulResponse)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				slog.Error("kratos unsuccessful status", "status", resp.StatusCode)
				return nil, connect.NewError(connect.CodeUnauthenticated, errKratosUnsucessfulResponse)
			}

			if !session.GetActive() {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("inactive session"))
			}

			identity := session.GetIdentity()

			if identity.SchemaId != orgSchemaID {
				slog.Info("attempt to perform organizer action",
					"curr_schema_id", identity.SchemaId,
					"want_schema_id", orgSchemaID)
				return nil, connect.NewError(connect.CodePermissionDenied, errors.New("permission denied"))
			}

			ctx = kratosx.WithIdentityID(ctx, identity.Id)
			return next(ctx, req)
		})
	})
}
