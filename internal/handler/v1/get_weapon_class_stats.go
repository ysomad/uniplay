package v1

import (
	"net/http"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) GetWeaponClassStats(w http.ResponseWriter, r *http.Request, steamID uint64, params v1.GetWeaponClassStatsParams) {
	// s, err := h.statistic.GetWeaponClassStats(r.Context(), steamID, params.ClassId)
	// if err != nil {
	// 	h.log.Error("http - v1 - handler.GetWeaponClassStats")

	// 	if errors.Is(err, domain.ErrWeaponClassStatsNotFound) {
	// 		writeError(w, http.StatusNotFound, err)
	// 		return
	// 	}

	// 	writeError(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// writeBody(w, http.StatusOK, s)
}
