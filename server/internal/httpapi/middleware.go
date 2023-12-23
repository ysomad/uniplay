package httpapi

import (
	"errors"
	"fmt"
	"net/http"

	kratos "github.com/ory/kratos-client-go"

	"github.com/ysomad/uniplay/server/internal/kratosctx"
)

var errIdentityNotMatch = errors.New("session identity not match")

func newAuthMiddleware(client *kratos.APIClient, orgSchemaID string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("ory_kratos_session")
			if err != nil {
				writerError(w, http.StatusUnauthorized, err)
				return
			}

			ctx := r.Context()

			session, resp, err := client.FrontendApi.
				ToSession(ctx).
				Cookie(cookie.String()).
				Execute()
			if err != nil {
				writerError(w, http.StatusUnauthorized, err)
				return
			}

			if resp.StatusCode != http.StatusOK {
				writeStatus(w, http.StatusUnauthorized)
				return
			}

			if !session.GetActive() {
				writeStatus(w, http.StatusUnauthorized)
				return
			}

			identity := session.GetIdentity()

			if identity.SchemaId != orgSchemaID {
				writerError(w, http.StatusForbidden,
					fmt.Errorf("%w, must be %s", errIdentityNotMatch, orgSchemaID))
				return
			}

			ctx = kratosctx.WithIdentityID(ctx, identity.Id)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
