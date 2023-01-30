package compendium

import (
	"context"
	"net/http"

	"github.com/ysomad/uniplay/internal/domain"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/compendium"
)

type compendiumService interface {
	GetWeaponList(ctx context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error)
}

type Controller struct {
	compendium compendiumService
}

func NewController(c compendiumService) *Controller {
	return &Controller{
		compendium: c,
	}
}

func (c *Controller) GetWeapons(p compendium.GetWeaponsParams) compendium.GetWeaponsResponder {
	weapons, err := c.compendium.GetWeaponList(p.HTTPRequest.Context())
	if err != nil {
		return compendium.NewGetWeaponsInternalServerError().
			WithPayload(&models.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	payload := make(models.WeaponList, len(weapons))

	for i, w := range weapons {
		payload[i] = models.WeaponListInner{
			WeaponID: w.WeaponID,
			Weapon:   w.Weapon,
			ClassID:  w.ClassID,
			Class:    w.Class,
		}
	}

	return compendium.NewGetWeaponsOK().WithPayload(payload)
}

func (c *Controller) GetWeaponClasses(p compendium.GetWeaponClassesParams) compendium.GetWeaponClassesResponder {
	classes, err := c.compendium.GetWeaponClassList(p.HTTPRequest.Context())
	if err != nil {
		return compendium.NewGetWeaponClassesInternalServerError().
			WithPayload(&models.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	payload := make(models.WeaponClassList, len(classes))

	for i, wc := range classes {
		payload[i] = models.WeaponClassListInner{
			ID:    wc.ID,
			Class: wc.Class,
		}
	}

	return compendium.NewGetWeaponClassesOK().WithPayload(payload)
}
