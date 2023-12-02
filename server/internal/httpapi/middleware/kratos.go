package middleware

import (
	"errors"
	"fmt"
	"net/http"

	ory "github.com/ory/client-go"

	"github.com/ysomad/uniplay/internal/httpapi/writer"
)

var errIdentityNotMatch = errors.New("session identity not match")

// NewSessionAuth returns middleware which is authenticates request for identity,
// returns 401 if received session from kratos has identity schema id different from given schema id.
// Must be used only for requests from browser.
func NewSessionAuth(client *ory.APIClient, schemaID string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("ory_kratos_session")
			if err != nil {
				writer.Error(w, http.StatusUnauthorized, err)
				return
			}

			session, resp, err := client.FrontendAPI.
				ToSession(r.Context()).
				Cookie(cookie.String()).
				Execute()
			if err != nil {
				writer.Error(w, http.StatusUnauthorized, err)
				return
			}

			if resp.StatusCode != http.StatusOK {
				writer.Status(w, http.StatusUnauthorized)
				return
			}

			if !session.GetActive() {
				writer.Status(w, http.StatusUnauthorized)
				return
			}

			if session.GetIdentity().SchemaId != schemaID {
				writer.Error(w, http.StatusForbidden,
					fmt.Errorf("%w, must be %s", errIdentityNotMatch, schemaID))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
