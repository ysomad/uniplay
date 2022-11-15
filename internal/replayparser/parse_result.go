package replayparser

import (
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type parseResult struct {
	metrics       *playerMetrics
	weaponMetrics *weaponMetrics
	match         *match
}

func (r *parseResult) MetricList(matchID domain.MatchID) []dto.Metric {
	return r.metrics.toDTO(matchID)
}

func (r *parseResult) WeaponMetricList(matchID domain.MatchID) []dto.WeaponMetric {
	return r.weaponMetrics.toDTO(matchID)
}

// Match returns dto.Match without ID.
func (r *parseResult) Match() *dto.Match {
	return r.match.toDTO()
}

func (r *parseResult) PlayerSteamIDs() []uint64 {
	return append(r.match._team1.playerSteamIDs, r.match._team2.playerSteamIDs...)
}

func (r *parseResult) TeamPlayers() []dto.TeamPlayer {
	team1len := len(r.match._team1.playerSteamIDs)
	tp := make([]dto.TeamPlayer, team1len+len(r.match._team2.playerSteamIDs))

	for i, steamID := range r.match._team1.playerSteamIDs {
		tp[i] = dto.TeamPlayer{
			TeamName:      r.match._team1.clanName,
			PlayerSteamID: steamID,
		}
	}
	for i, steamID := range r.match._team2.playerSteamIDs {
		tp[team1len+i] = dto.TeamPlayer{
			TeamName:      r.match._team2.clanName,
			PlayerSteamID: steamID,
		}
	}

	return tp
}
