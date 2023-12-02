package middleware

import (
	"errors"
	"fmt"
	"net/http"

	ory "github.com/ory/client-go"

	"github.com/ysomad/uniplay/internal/appctx"
	"github.com/ysomad/uniplay/internal/config"
	"github.com/ysomad/uniplay/internal/httpapi/writer"
)

type kratos struct {
	client            *ory.APIClient
	organizerSchemaID string
}

func NewKratos(client *ory.APIClient, conf config.Kratos) kratos {
	return kratos{
		client:            client,
		organizerSchemaID: conf.OrganizerSchemaID,
	}
}

var (
	errIdentityNotMatch = errors.New("session identity not match")
)

// SessionAuth returns middleware which is authenticates request.
// returns 401 if received session from kratos has identity schema id different from given schema id.
// Must be used only for requests from browser.
func (k kratos) SessionAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("ory_kratos_session")
		if err != nil {
			writer.Error(w, http.StatusUnauthorized, err)
			return
		}

		ctx := r.Context()

		session, resp, err := k.client.FrontendAPI.
			ToSession(ctx).
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

		identity := session.GetIdentity()

		if identity.SchemaId != k.organizerSchemaID {
			writer.Error(w, http.StatusForbidden,
				fmt.Errorf("%w, must be %s", errIdentityNotMatch, k.organizerSchemaID))
			return
		}

		ctx = appctx.WithIdentityID(ctx, identity.Id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
