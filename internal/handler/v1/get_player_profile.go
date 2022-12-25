package v1

import (
	"errors"
	"net/http"

	"github.com/ssssargsian/uniplay/internal/domain"
	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
	"github.com/ssssargsian/uniplay/internal/pkg/apperror"
	"go.uber.org/zap"
)

func (h *handler) GetPlayerProfile(w http.ResponseWriter, r *http.Request, steamID uint64) {
	p, err := h.player.Get(r.Context(), steamID)
	if err != nil {
		h.log.Error("http - v1 - handler.GetPlayerProfile", zap.Error(err))

		if errors.Is(err, domain.ErrPlayerNotFound) {
			apperror.Write(w, http.StatusBadRequest, err)
			return
		}

		apperror.Write(w, http.StatusInternalServerError, err)
		return
	}

	writeBody(w, http.StatusOK, v1.Player{
		SteamID:      p.SteamID,
		TeamName:     p.TeamName,
		TeamFlagCode: p.TeamFlagCode,
		CreateTime:   p.CreateTime,
		UpdateTime:   p.UpdateTime,
	})
}
