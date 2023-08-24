package account

import (
	"errors"
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	gen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/account"
)

type Controller struct {
	account *service
}

func NewController(s *service) *Controller {
	return &Controller{account: s}
}

func (c *Controller) CreateAccount(p gen.CreateAccountParams) gen.CreateAccountResponder {
	account, err := c.account.Create(p.HTTPRequest.Context(), p.Payload.Email.String(), p.Payload.Password)
	if err != nil {
		if errors.Is(err, domain.ErrAccountEmailTaken) {
			return gen.NewCreateAccountConflict().WithPayload(&models.Error{
				Code:    domain.CodeAccountEmailTaken,
				Message: err.Error(),
			})
		}

		return gen.NewCreateAccountInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return gen.NewCreateAccountOK().WithPayload(&models.Account{
		ID:        account.ID.String(),
		Email:     account.Email,
		CreatedAt: strfmt.DateTime(account.CreatedAt),
	})
}
