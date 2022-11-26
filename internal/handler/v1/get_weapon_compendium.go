package v1

import (
	"net/http"

	"go.uber.org/zap"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) GetWeaponCompendium(w http.ResponseWriter, r *http.Request) {
	wl, err := h.compendium.GetWeaponList(r.Context())
	if err != nil {
		h.log.Error("http - v1 - handler.GetWeaponCompendium", zap.Error(err))
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	res := make(v1.WeaponList, len(wl))
	for i, w := range wl {
		res[i] = v1.Weapon{
			Name:      w.Name,
			ClassID:   uint8(w.ClassID),
			ClassName: w.ClassName,
		}
	}

	writeBody(w, http.StatusOK, res)
}
