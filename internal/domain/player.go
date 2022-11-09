package domain

import "time"

type Player struct {
	SteamID      uint64
	TeamName     string
	TeamFlagCode string
	CreateTime   time.Time
	UpdateTime   time.Time
}

type PlayerStats struct {
	TotalKills             int32
	TotalDeaths            int32
	KillDeathRatio         float32
	DamagePerRound         float32 // ADR
	GrenadeDamangePerRound float32
	KillsPerRound          float32 // KPR
	AssistsPerRound        float32
	DeathsPerRound         float32 // DPR
	HeadshotPercentage     float32
	MatchesPlayed          uint16
	RoundsPlayed           uint32
}

type PlayerWeaponStats []struct {
	WeaponName string
	TotalKills int32
}

/* TODO:
https://www.hltv.org/stats/players/individual/7998/s1mple
1. Total opening kills
2. Total opening deaths
3. Opening kill ration
4. Team win percent after first kill
5. First kill in won rounds
6. Rifle kills
7. Sniper kills
8. SMG kills
9. Pistol kills
10. Grenade kills
11. Other kill
12. 0 kill rounds
13. 1 kill rounds
14. 2 kill rounds
15. 3 kill rounds
16. 4 kill rounds
17. 5 kill rounds
*/
