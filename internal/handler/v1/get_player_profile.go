package v1

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) GetPlayerProfile(w http.ResponseWriter, r *http.Request, steamID uint64) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: implement
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(v1.PlayerProfile{
		MainTeamName: "Na`Vi",
		SteamID:      76561197984560929,
		Stats: v1.PlayerStats{
			AssistsPerRound:       rand.Float32(),
			BlindPerRound:         rand.Float32(),
			BlindedPerRound:       rand.Float32(),
			DamagePerRound:        rand.Float32(),
			DeathsPerRound:        rand.Float32(),
			GrenadeDamagePerRound: rand.Float32(),
			HeadshotPercentage:    rand.Float32(),
			KillDeathRatio:        rand.Float32(),
			KillsPerRound:         rand.Float32(),
			MatchesPlayed:         uint16(rand.Uint32()),
			RoundsPlayed:          rand.Uint32(),
			TotalDeaths:           rand.Uint32(),
			TotalKills:            rand.Uint32(),
		},
		WeaponStats: v1.WeaponStats{
			{
				TotalKills: rand.Uint32(),
				WeaponName: "AK-47",
			},
			{
				TotalKills: rand.Uint32(),
				WeaponName: "Deagle",
			},
			{
				TotalKills: rand.Uint32(),
				WeaponName: "M4A1",
			},
			{
				TotalKills: rand.Uint32(),
				WeaponName: "AWP",
			},
			{
				TotalKills: rand.Uint32(),
				WeaponName: "USP-S",
			},
		},
		WeaponClassStats: v1.WeaponClassStats{
			{
				TotalKills:  rand.Uint32(),
				WeaponClass: "rifle",
			},
			{
				TotalKills:  rand.Uint32(),
				WeaponClass: "pistols",
			},
			{
				TotalKills:  rand.Uint32(),
				WeaponClass: "smg",
			},
		},
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
