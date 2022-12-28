package v1

import (
	"errors"
	"net/http"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/pkg/apperror"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) GetWeaponStats(w http.ResponseWriter, r *http.Request, steamID uint64, params v1.GetWeaponStatsParams) {
	s, err := h.player.GetWeaponStats(r.Context(), steamID, domain.WeaponStatsFilter{
		WeaponID: params.WeaponId,
		ClassID:  params.ClassId,
	})
	if err != nil {
		h.log.Error("http - v1 - handler.GetWeaponStats")

		if errors.Is(err, domain.ErrPlayerNotFound) {
			apperror.Write(w, http.StatusNotFound, err)
			return
		}

		apperror.Write(w, http.StatusInternalServerError, err)
		return
	}

	if err = writeBody(w, http.StatusOK, s); err != nil {
		apperror.Write(w, http.StatusInternalServerError, err)
	}
}
