package v1

import (
	"net/http"

	"go.uber.org/zap"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) GetWeaponClassCompendium(w http.ResponseWriter, r *http.Request) {
	wc, err := h.compendium.GetWeaponClassList(r.Context())
	if err != nil {
		h.log.Error("http - v1 - handler.GetWeaponClassCompendium", zap.Error(err))
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	res := make(v1.WeaponClassList, len(wc))
	for i, c := range wc {
		res[i] = v1.WeaponClass{
			ID:   uint8(c.ID),
			Name: c.Name,
		}
	}

	writeBody(w, http.StatusOK, res)
}
