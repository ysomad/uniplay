package replayparser

import (
	"errors"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type parseResult struct {
	metrics       *playerMetrics
	weaponMetrics *weaponMetrics
	match         *match
}

func (r *parseResult) MetricList(matchID domain.MatchID) ([]dto.Metric, error) {
	return r.metrics.toDTO(matchID)
}

func (r *parseResult) WeaponMetricList(matchID domain.MatchID) ([]dto.WeaponMetric, error) {
	return r.weaponMetrics.toDTO(matchID)
}

// Match returns dto.Match without ID.
func (r *parseResult) Match() *dto.Match {
	return r.match.toDTO()
}

// PlayerSteamIDs returns list of player steam ids or error if list is empty.
func (r *parseResult) PlayerSteamIDs() ([]uint64, error) {
	if len(r.match.team1.playerSteamIDs)+len(r.match.team2.playerSteamIDs) == 0 {
		return nil, errors.New("empty list of player steam ids")
	}
	return append(r.match.team1.playerSteamIDs, r.match.team2.playerSteamIDs...), nil
}

func (r *parseResult) TeamPlayers() []dto.TeamPlayer {
	team1PlayersAmount := len(r.match.team1.playerSteamIDs)
	players := make([]dto.TeamPlayer, team1PlayersAmount+len(r.match.team2.playerSteamIDs))

	for i, steamID := range r.match.team1.playerSteamIDs {
		players[i] = dto.TeamPlayer{
			TeamName:   r.match.team1.clanName,
			SteamID:    steamID,
			MatchState: domain.NewPlayerMatchState(uint8(r.match.team1.score), uint8(r.match.team2.score)),
		}
	}

	for i, steamID := range r.match.team2.playerSteamIDs {
		players[team1PlayersAmount+i] = dto.TeamPlayer{
			TeamName:   r.match.team2.clanName,
			SteamID:    steamID,
			MatchState: domain.NewPlayerMatchState(uint8(r.match.team2.score), uint8(r.match.team1.score)),
		}
	}

	return players
}
