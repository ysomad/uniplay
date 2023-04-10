package compendium

import (
	"net/http"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	"github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/compendium"
)

type Controller struct {
	compendium *service
}

func NewController(s *service) *Controller {
	return &Controller{
		compendium: s,
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

func (c *Controller) GetMaps(p compendium.GetMapsParams) compendium.GetMapsResponder {
	maps, err := c.compendium.GetMapList(p.HTTPRequest.Context())
	if err != nil {
		return compendium.NewGetMapsInternalServerError().
			WithPayload(&models.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
	}

	payload := make(models.MapList, len(maps))

	for i, m := range maps {
		payload[i] = models.Map{
			Name:    m.Name,
			IconURL: m.IconURL,
		}
	}

	return compendium.NewGetMapsOK().WithPayload(payload)
}
