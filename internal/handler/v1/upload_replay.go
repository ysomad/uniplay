package v1

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/pkg/apperror"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) UploadReplay(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 160<<20) // max body size is 300 mb

	if err := r.ParseMultipartForm(64 << 20); err != nil { // 64 mb
		apperror.Write(w, http.StatusBadRequest, fmt.Errorf("r.ParseMultipartForm: %w", err))
		return
	}

	file, fileHeader, err := r.FormFile("replay")
	if err != nil {
		apperror.Write(w, http.StatusBadRequest, fmt.Errorf("r.FormFile: %w", err))
		return
	}
	defer file.Close()

	filenameParts := strings.Split(fileHeader.Filename, ".")
	if filenameParts[len(filenameParts)-1] != "dem" {
		apperror.Write(w, http.StatusBadRequest, errInvalidReplayFileExtension)
		return
	}

	if err = os.MkdirAll("./tmp", os.ModePerm); err != nil {
		apperror.Write(w, http.StatusInternalServerError, fmt.Errorf("os.MkdirAll: %w", err))
		return
	}

	replayFilename := fmt.Sprintf("./tmp/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))

	dest, err := os.Create(replayFilename)
	if err != nil {
		apperror.Write(w, http.StatusInternalServerError, fmt.Errorf("os.Create: %w", err))
		return
	}
	defer dest.Close()

	defer func() {
		if err := os.Remove(replayFilename); err != nil {
			apperror.Write(w, http.StatusInternalServerError, fmt.Errorf("os.Remove: %w", err))
			return
		}
	}()

	if _, err = io.Copy(dest, file); err != nil {
		apperror.Write(w, http.StatusInternalServerError, fmt.Errorf("io.Copy: %w", err))
		return
	}

	match, err := h.replay.CollectStats(r.Context(), replayFilename)
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
			PlayerSteamIds: match.Team1.PlayerSteamIDs,
			Score:          match.Team1.Score,
		},
		Team2: v1.MatchTeam{
			ClanName:       match.Team2.ClanName,
			FlagCode:       match.Team2.FlagCode,
			PlayerSteamIds: match.Team2.PlayerSteamIDs,
			Score:          match.Team2.Score,
		},
		UploadedAt: match.UploadedAt,
	})
}
