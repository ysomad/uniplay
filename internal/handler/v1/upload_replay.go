package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) UploadReplay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TODO: implement
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(v1.ReplayUploadResponse{
		MapName:       "de_dust2",
		MatchDuration: time.Hour,
		MatchID:       uuid.NewMD5(uuid.UUID{}, []byte("YEEET")),
		Team1: v1.MatchTeam{
			ClanName: "Na`Vi",
			FlagCode: "UA",
			Score:    7,
			PlayerSteamIds: []int{
				76561198039986599,
				76561197989430253,
				76561198039986599,
				76561197989430253,
				76561197989430253,
			},
		},
		Team2: v1.MatchTeam{
			ClanName: "Virtus.PRO",
			FlagCode: "RU",
			Score:    19,
			PlayerSteamIds: []int{
				76561198039986599,
				76561197989430253,
				76561198039986599,
				76561197989430253,
				76561197989430253,
			},
		},
		UploadTime: time.Now(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
