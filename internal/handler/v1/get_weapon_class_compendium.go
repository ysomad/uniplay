package v1

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *handler) GetWeaponClassCompendium(w http.ResponseWriter, r *http.Request) {
	wc, err := h.compendium.GetWeaponClassList(r.Context())
	if err != nil {
		h.log.Error("http - v1 - handler.GetWeaponClassCompendium", zap.Error(err))
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeBody(w, http.StatusOK, wc)
}
