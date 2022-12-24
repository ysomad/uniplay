package replayparser

import (
	"sync"
	"time"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/dto"
)

type parseResult struct {
	match match
	stats stats
}

func (r *parseResult) Match() *dto.ReplayMatch {
	return &dto.ReplayMatch{
		ID: r.match.id,
		Team1: dto.ReplayTeam{
			ClanName:       r.match.team1.clanName,
			FlagCode:       r.match.team1.flagCode,
			Score:          int8(r.match.team1.score),
			PlayerSteamIDs: r.match.team1.playerSteamIDs,
			State:          domain.NewMatchState(int8(r.match.team1.score), int8(r.match.team2.score)),
		},
		Team2: dto.ReplayTeam{
			ClanName:       r.match.team2.clanName,
			FlagCode:       r.match.team2.flagCode,
			Score:          int8(r.match.team2.score),
			PlayerSteamIDs: r.match.team2.playerSteamIDs,
			State:          domain.NewMatchState(int8(r.match.team2.score), int8(r.match.team1.score)),
		},
		MapName:    r.match.mapName,
		Duration:   r.match.duration,
		UploadedAt: time.Now(),
	}
}

func (r *parseResult) Stats() ([]*dto.PlayerStat, []*dto.PlayerWeaponStat) {
	var (
		wg          sync.WaitGroup
		playerStats []*dto.PlayerStat
		weaponStats []*dto.PlayerWeaponStat
	)

	wg.Add(2)

	go func() {
		for _, ps := range r.stats.playerStats {
			playerStats = append(playerStats, ps)
		}

		wg.Done()
	}()

	go func() {
		for _, weapons := range r.stats.weaponStats {
			for _, ws := range weapons {
				weaponStats = append(weaponStats, ws)
			}
		}

		wg.Done()
	}()

	wg.Wait()
	return playerStats, weaponStats
}
