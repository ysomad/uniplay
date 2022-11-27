package replayparser

import (
	"errors"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type replayStats struct {
	ps *playerStats
	ws *weaponStats
	m  *match
}

func (s *replayStats) PlayerStats() ([]dto.PlayerStat, error) { return s.ps.toDTO() }
func (s *replayStats) WeaponStats() ([]dto.WeaponStat, error) { return s.ws.toDTO() }

// Match returns dto.Match without ID.
func (s *replayStats) Match() *dto.Match {
	return s.m.toDTO()
}

// PlayerSteamIDs returns list of player steam ids or error if list is empty.
func (s *replayStats) PlayerSteamIDs() ([]uint64, error) {
	if len(s.m.team1.playerSteamIDs)+len(s.m.team2.playerSteamIDs) == 0 {
		return nil, errors.New("empty list of player steam ids")
	}
	return append(s.m.team1.playerSteamIDs, s.m.team2.playerSteamIDs...), nil
}

func (s *replayStats) TeamPlayers() []dto.TeamPlayer {
	team1PlayersAmount := len(s.m.team1.playerSteamIDs)
	players := make([]dto.TeamPlayer, team1PlayersAmount+len(s.m.team2.playerSteamIDs))

	for i, steamID := range s.m.team1.playerSteamIDs {
		players[i] = dto.TeamPlayer{
			TeamName:   s.m.team1.clanName,
			SteamID:    steamID,
			MatchState: domain.NewPlayerMatchState(uint8(s.m.team1.score), uint8(s.m.team2.score)),
		}
	}

	for i, steamID := range s.m.team2.playerSteamIDs {
		players[team1PlayersAmount+i] = dto.TeamPlayer{
			TeamName:   s.m.team2.clanName,
			SteamID:    steamID,
			MatchState: domain.NewPlayerMatchState(uint8(s.m.team2.score), uint8(s.m.team1.score)),
		}
	}

	return players
}
