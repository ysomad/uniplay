package middleware

import (
	"fmt"
	"net/http"

	ory "github.com/ory/client-go"

	"github.com/ysomad/uniplay/internal/domain/identity"
	"github.com/ysomad/uniplay/internal/errorx"
)

// NewSessionAuth returns middleware which is authenticates request for identity,
// returns 401 if received session from kratos has identity schema id different from given schema.
// Must be used only for requests from browser.
func NewSessionAuth(client *ory.APIClient, identity identity.SchemaID) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("ory_kratos_session")
			if err != nil {
				errorx.WriteError(w, http.StatusUnauthorized, err)
				return
			}

			session, resp, err := client.FrontendAPI.
				ToSession(r.Context()).
				Cookie(cookie.String()).
				Execute()
			if err != nil {
				errorx.WriteError(w, http.StatusUnauthorized, err)
				return
			}

			if resp.StatusCode != http.StatusOK {
				errorx.WriteStatus(w, http.StatusUnauthorized)
				return
			}

			if !session.GetActive() {
				errorx.WriteStatus(w, http.StatusUnauthorized)
				return
			}

			if session.GetIdentity().SchemaId != string(identity) {
				errorx.WriteMessage(w, http.StatusForbidden, fmt.Sprintf("must have '%s' identity to use this endpoint", identity))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
