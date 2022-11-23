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
			AssistsPerRound:       rand.Float64(),
			BlindPerRound:         rand.Float64(),
			BlindedPerRound:       rand.Float64(),
			DamagePerRound:        rand.Float64(),
			DeathsPerRound:        rand.Float64(),
			GrenadeDamagePerRound: rand.Float64(),
			HeadshotPercentage:    rand.Float64(),
			KillDeathRatio:        rand.Float64(),
			KillsPerRound:         rand.Float64(),
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
