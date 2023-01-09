package v1

import (
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/pkg/apperror"
	"github.com/ssssargsian/uniplay/internal/replay"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) UploadReplay(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 160<<20) // max body size is 160 mb

	if err := r.ParseMultipartForm(64 << 20); err != nil { // 64 mb
		apperror.Write(w, http.StatusBadRequest, fmt.Errorf("r.ParseMultipartForm: %w", err))
		return
	}

	file, header, err := r.FormFile("replay")
	if err != nil {
		apperror.Write(w, http.StatusBadRequest, fmt.Errorf("r.FormFile: %w", err))
		return
	}

	replay, err := replay.New(file, header.Filename)
	if err != nil {
		apperror.Write(w, http.StatusBadRequest, fmt.Errorf("replay.New: %w", err))
		return
	}
	defer replay.Close()

	match, err := h.replay.CollectStats(r.Context(), replay)
	if err != nil {
		h.log.Error("http - v1 - handler.UploadReplay", zap.Error(err))

		switch {
		case errors.Is(err, domain.ErrMatchAlreadyExist):
			apperror.Write(w, http.StatusConflict, err)
			return
		}

		apperror.Write(w, http.StatusInternalServerError, err)
		return
	}

	writeBody(w, http.StatusOK, v1.Match{
		MatchID:       match.ID,
		MapName:       match.MapName,
		MatchDuration: match.Duration,
		Team1: v1.MatchTeam{
			ClanName:       match.Team1.ClanName,
			FlagCode:       match.Team1.FlagCode,
			PlayerSteamIds: match.Team1.Players,
			Score:          match.Team1.Score,
		},
		Team2: v1.MatchTeam{
			ClanName:       match.Team2.ClanName,
			FlagCode:       match.Team2.FlagCode,
			PlayerSteamIds: match.Team2.Players,
			Score:          match.Team2.Score,
		},
		UploadedAt: match.UploadedAt,
	})
}
