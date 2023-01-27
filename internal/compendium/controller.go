package compendium

import (
	"context"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"

	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/models"
	"github.com/ssssargsian/uniplay/internal/gen/swagger2/v1/restapi/operations/compendium"
)

type compendiumService interface {
	GetWeaponList(ctx context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error)
}

type Controller struct {
	log        *zap.Logger
	compendium compendiumService
}

func NewController(l *zap.Logger, c compendiumService) *Controller {
	return &Controller{
		log:        l,
		compendium: c,
	}
}

func (c *Controller) GetWeapons(p compendium.GetWeaponsParams) compendium.GetWeaponsResponder {
	weapons, err := c.compendium.GetWeaponList(p.HTTPRequest.Context())
	if err != nil {
		return compendium.NewGetWeaponsInternalServerError()
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
		return compendium.NewGetWeaponClassesInternalServerError()
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
