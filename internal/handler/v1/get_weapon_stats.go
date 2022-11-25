package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ssssargsian/uniplay/internal/domain"
	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) GetWeaponStats(w http.ResponseWriter, r *http.Request, steamID uint64) {
	var rbody v1.WeaponStatsRequest
	if err := json.NewDecoder(r.Body).Decode(&rbody); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	s, err := h.statistic.GetWeaponStats(r.Context(), steamID, domain.WeaponStatsFilter{
		WeaponName:  rbody.WeaponName,
		WeaponClass: domain.NewWeaponClass(rbody.WeaponClass),
	})
	if err != nil {
		h.log.Error("http - v1 - handler.GetWeaponStats")

		if errors.Is(err, domain.ErrWeaponStatsNotFound) {
			writeError(w, http.StatusNotFound, err)
			return
		}

		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeBody(w, http.StatusOK, s)
}
