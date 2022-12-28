package v1

import (
	"errors"
	"net/http"

	"github.com/ssssargsian/uniplay/internal/domain"
	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
	"github.com/ssssargsian/uniplay/internal/pkg/apperror"
	"go.uber.org/zap"
)

func (h *handler) GetPlayerStats(w http.ResponseWriter, r *http.Request, steamID uint64) {
	stats, err := h.player.GetStats(r.Context(), steamID)
	if err != nil {
		h.log.Error("http - v1 - handler.GetPlayerStats", zap.Error(err))

		if errors.Is(err, domain.ErrPlayerNotFound) {
			apperror.Write(w, http.StatusNotFound, err)
			return
		}

		apperror.Write(w, http.StatusInternalServerError, err)
		return
	}

	err = writeBody(w, http.StatusOK, v1.PlayerStats{
		TotalStats: v1.TotalStats{
			Assists:            stats.Total.Assists,
			BlindKills:         stats.Total.BlindKills,
			BlindedPlayers:     stats.Total.BlindedPlayers,
			BlindedTimes:       stats.Total.BlindedTimes,
			BombsDefused:       stats.Total.BombsDefused,
			BombsPlanted:       stats.Total.BombsPlanted,
			DamageDealt:        stats.Total.DamageDealt,
			DamageTaken:        stats.Total.DamageTaken,
			Deaths:             stats.Total.Deaths,
			Draws:              stats.Total.Draws,
			FlashbangAssists:   stats.Total.FlashbangAssists,
			GrenadeDamageDealt: stats.Total.GrenadeDamageDealt,
			HeadshotKills:      stats.Total.HeadshotKills,
			Kills:              stats.Total.Kills,
			Loses:              stats.Total.Loses,
			MatchesPlayed:      stats.Total.MatchesPlayed,
			MVPCount:           stats.Total.MVPCount,
			NoScopeKills:       stats.Total.NoScopeKills,
			RoundsPlayed:       stats.Total.RoundsPlayed,
			ThroughSmokeKills:  stats.Total.ThroughSmokeKills,
			TimePlayed:         stats.Total.TimePlayed,
			WallbangKills:      stats.Total.WallbangKills,
			Wins:               stats.Total.Wins,
		},
		CalculatedStats: v1.CalculatedStats{
			HeadshotPercentage: stats.Calc.HeadshotPercentage,
			KillDeathRatio:     stats.Calc.KillDeathRatio,
			WinRate:            stats.Calc.WinRate,
		},
		RoundStats: v1.RoundStats{
			Assists:            stats.Round.Assists,
			BlindedPlayers:     stats.Round.BlindedPlayers,
			BlindedTimes:       stats.Round.BlindedTimes,
			DamageDealt:        stats.Round.DamageDealt,
			Deaths:             stats.Round.Deaths,
			GrenadeDamageDealt: stats.Round.GrenadeDamageDealt,
			Kills:              stats.Round.Kills,
		},
	})
	if err != nil {
		apperror.Write(w, http.StatusInternalServerError, err)
	}
}
