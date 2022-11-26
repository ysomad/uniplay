package v1

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) GetWeaponCompendium(w http.ResponseWriter, r *http.Request) {
	wl, err := h.compendium.GetWeaponList(r.Context())
	if err != nil {
		h.log.Error("http - v1 - handler.GetWeaponCompendium", zap.Error(err))
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeBody(w, http.StatusOK, wl)
}
