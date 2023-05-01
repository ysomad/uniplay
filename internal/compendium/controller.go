package compendium

import (
	"net/http"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	gen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/compendium"
)

type Controller struct {
	compendium *service
}

func NewController(s *service) *Controller {
	return &Controller{
		compendium: s,
	}
}

func (c *Controller) GetWeapons(p gen.GetWeaponsParams) gen.GetWeaponsResponder {
	weapons, err := c.compendium.GetWeaponList(p.HTTPRequest.Context())
	if err != nil {
		return gen.NewGetWeaponsInternalServerError().
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

	return gen.NewGetWeaponsOK().WithPayload(payload)
}

func (c *Controller) GetWeaponClasses(p gen.GetWeaponClassesParams) gen.GetWeaponClassesResponder {
	classes, err := c.compendium.GetWeaponClassList(p.HTTPRequest.Context())
	if err != nil {
		return gen.NewGetWeaponClassesInternalServerError().
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

	return gen.NewGetWeaponClassesOK().WithPayload(payload)
}

func (c *Controller) GetMaps(p gen.GetMapsParams) gen.GetMapsResponder {
	maps, err := c.compendium.GetMapList(p.HTTPRequest.Context())
	if err != nil {
		return gen.NewGetMapsInternalServerError().
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

	return gen.NewGetMapsOK().WithPayload(payload)
}

func (c *Controller) GetCities(p gen.GetCitiesParams) gen.GetCitiesResponder {
	var searchQuery string

	if p.Search != nil {
		searchQuery = *p.Search
	}

	cities, err := c.compendium.GetCityList(p.HTTPRequest.Context(), searchQuery)
	if err != nil {
		return gen.NewGetCitiesInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := make(models.CityList, len(cities))

	for i, c := range cities {
		payload[i] = models.City(c)
	}

	return gen.NewGetCitiesOK().WithPayload(payload)
}
