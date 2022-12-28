package v1

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

func (h *handler) GetPlayerMatches(w http.ResponseWriter, r *http.Request, steamID uint64) {
	w.Header().Set("Content-Type", "application/json")
	// TODO: implement

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(v1.MatchList{
		NextPageToken: base64.RawURLEncoding.EncodeToString([]byte("f96d60c8-3338-4d14-aef1-a3583ff2914c,testdata")),
		Matches: []v1.Match{
			{
				MapName:       "de_dust2",
				MatchDuration: time.Hour,
				MatchID:       uuid.NewMD5(uuid.UUID{}, []byte("yeeeeeeeeeeeet")),
				Team1: v1.MatchTeam{
					ClanName: "Na`Vi",
					FlagCode: "UA",
					Score:    7,
					PlayerSteamIds: []uint64{
						steamID,
						76561197989430253,
						76561198039986599,
						76561197989430253,
						76561197989430253,
					},
				},
				Team2: v1.MatchTeam{
					ClanName: "yeet",
					FlagCode: "RU",
					Score:    16,
					PlayerSteamIds: []uint64{
						76561198039986599,
						76561197989430253,
						76561198039986599,
						76561197989430253,
						76561197989430253,
					},
				},
			},
			{
				MapName:       "de_tuscan",
				MatchDuration: time.Hour,
				MatchID:       uuid.NewMD5(uuid.UUID{}, []byte("yeeeeeeeeeeeet2")),
				Team1: v1.MatchTeam{
					ClanName: "Na`Vi",
					FlagCode: "UA",
					Score:    5,
					PlayerSteamIds: []uint64{
						steamID,
						76561197989430253,
						76561198039986599,
						76561197989430253,
						76561197989430253,
					},
				},
				Team2: v1.MatchTeam{
					ClanName: "fnatic",
					FlagCode: "GR",
					Score:    16,
					PlayerSteamIds: []uint64{
						76561198039986599,
						76561197989430253,
						76561198039986599,
						76561197989430253,
						76561197989430253,
					},
				},
			},
			{
				MapName:       "de_overpass",
				MatchDuration: time.Hour,
				MatchID:       uuid.NewMD5(uuid.UUID{}, []byte("yeeeeeeeeeeeet3")),
				Team1: v1.MatchTeam{
					ClanName: "Na`Vi",
					FlagCode: "UA",
					Score:    11,
					PlayerSteamIds: []uint64{
						steamID,
						76561197989430253,
						76561198039986599,
						76561197989430253,
						76561197989430253,
					},
				},
				Team2: v1.MatchTeam{
					ClanName: "virtus.pro",
					FlagCode: "RU",
					Score:    20,
					PlayerSteamIds: []uint64{
						76561198039986599,
						76561197989430253,
						76561198039986599,
						76561197989430253,
						76561197989430253,
					},
				},
			},
			{
				MapName:       "de_inferno",
				MatchDuration: time.Hour,
				MatchID:       uuid.NewMD5(uuid.UUID{}, []byte("yeeeeeeeeeeeet4")),
				Team1: v1.MatchTeam{
					ClanName: "Na`Vi",
					FlagCode: "UA",
					Score:    0,
					PlayerSteamIds: []uint64{
						steamID,
						76561197989430253,
						76561198039986599,
						76561197989430253,
						76561197989430253,
					},
				},
				Team2: v1.MatchTeam{
					ClanName: "cloud9",
					FlagCode: "NA",
					Score:    16,
					PlayerSteamIds: []uint64{
						76561198039986599,
						76561197989430253,
						76561198039986599,
						76561197989430253,
						76561197989430253,
					},
				},
			},
		},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
