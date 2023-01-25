package replay

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/events"
	"github.com/stretchr/testify/assert"
)

func Test_replayMatch_swapTeamSides(t *testing.T) {
	type fields struct {
		id         uuid.UUID
		team1      replayTeam
		team2      replayTeam
		mapName    string
		duration   time.Duration
		uploadedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &replayMatch{
				id:         tt.fields.id,
				team1:      tt.fields.team1,
				team2:      tt.fields.team2,
				mapName:    tt.fields.mapName,
				duration:   tt.fields.duration,
				uploadedAt: tt.fields.uploadedAt,
			}
			m.swapTeamSides()
		})
	}
}

func Test_replayMatch_updateTeamsScore(t *testing.T) {
	type fields struct {
		id         uuid.UUID
		team1      replayTeam
		team2      replayTeam
		mapName    string
		duration   time.Duration
		uploadedAt time.Time
	}
	type args struct {
		e events.ScoreUpdated
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &replayMatch{
				id:         tt.fields.id,
				team1:      tt.fields.team1,
				team2:      tt.fields.team2,
				mapName:    tt.fields.mapName,
				duration:   tt.fields.duration,
				uploadedAt: tt.fields.uploadedAt,
			}
			m.updateTeamsScore(tt.args.e)
		})
	}
}

func Test_replayMatch_setTeamStates(t *testing.T) {
	type fields struct {
		id         uuid.UUID
		team1      replayTeam
		team2      replayTeam
		mapName    string
		duration   time.Duration
		uploadedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &replayMatch{
				id:         tt.fields.id,
				team1:      tt.fields.team1,
				team2:      tt.fields.team2,
				mapName:    tt.fields.mapName,
				duration:   tt.fields.duration,
				uploadedAt: tt.fields.uploadedAt,
			}
			m.setTeamStates()
		})
	}
}

func Test_replayMatch_teamPlayers(t *testing.T) {
	type fields struct {
		id         uuid.UUID
		team1      replayTeam
		team2      replayTeam
		mapName    string
		duration   time.Duration
		uploadedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []teamPlayer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &replayMatch{
				id:         tt.fields.id,
				team1:      tt.fields.team1,
				team2:      tt.fields.team2,
				mapName:    tt.fields.mapName,
				duration:   tt.fields.duration,
				uploadedAt: tt.fields.uploadedAt,
			}
			if got := m.teamPlayers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replayMatch.teamPlayers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newReplayTeam(t *testing.T) {
	type args struct {
		name    string
		flag    string
		side    common.Team
		players []*common.Player
	}
	tests := []struct {
		name string
		args args
		want replayTeam
	}{
		{
			name: "success",
			args: args{
				name: "Virtus PRO",
				flag: "RU",
				side: common.TeamCounterTerrorists,
				players: []*common.Player{
					{
						SteamID64: 1,
					},
					{
						SteamID64: 2,
					},
					{
						SteamID64: 3,
					},
					{
						SteamID64: 4,
					},
					{
						SteamID64: 5,
					},
					{
						SteamID64: 6,
					},
					{
						SteamID64: 7,
					},
					{
						SteamID64: 8,
					},
					{
						SteamID64: 9,
					},
					{
						SteamID64: 10,
					},
				},
			},
			want: replayTeam{
				clanName: "Virtus PRO",
				flagCode: "RU",
				players:  []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				_side:    common.TeamCounterTerrorists,
			},
		},
		{
			name: "nil players in slice",
			args: args{
				name: "Na`Vi",
				flag: "UA",
				side: common.TeamTerrorists,
				players: []*common.Player{
					nil,
					nil,
					{
						SteamID64: 3,
					},
					{
						SteamID64: 4,
					},
					{
						SteamID64: 5,
					},
					{
						SteamID64: 6,
					},
					{
						SteamID64: 7,
					},
					nil,
					{
						SteamID64: 9,
					},
					{
						SteamID64: 10,
					},
				},
			},
			want: replayTeam{
				clanName: "Na`Vi",
				flagCode: "UA",
				players:  []uint64{3, 4, 5, 6, 7, 9, 10},
				_side:    common.TeamTerrorists,
			},
		},
		{
			name: "invalid steamid64 in slice",
			args: args{
				name: "Mousesports",
				flag: "FR",
				side: common.TeamTerrorists,
				players: []*common.Player{
					{
						SteamID64: 0,
					},
					{
						SteamID64: 2,
					},
					{
						SteamID64: 3,
					},
					{
						SteamID64: 4,
					},
					{
						SteamID64: 0,
					},
					{
						SteamID64: 6,
					},
					{
						SteamID64: 7,
					},
					{
						SteamID64: 8,
					},
					{
						SteamID64: 0,
					},
					{
						SteamID64: 10,
					},
				},
			},
			want: replayTeam{
				clanName: "Mousesports",
				flagCode: "FR",
				players:  []uint64{2, 3, 4, 6, 7, 8, 10},
				_side:    common.TeamTerrorists,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newReplayTeam(tt.args.name, tt.args.flag, tt.args.side, tt.args.players)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_replayTeam_swapSide(t *testing.T) {
	type fields struct {
		_side common.Team
	}
	tests := []struct {
		name     string
		teamSide common.Team
		wantSide common.Team
	}{
		{
			name:     "swap t to ct",
			teamSide: common.TeamTerrorists,
			wantSide: common.TeamCounterTerrorists,
		},
		{
			name:     "swap ct to t",
			teamSide: common.TeamCounterTerrorists,
			wantSide: common.TeamTerrorists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &replayTeam{_side: tt.teamSide}
			tr.swapSide()

			assert.Equal(t, tt.wantSide, tr._side)
		})
	}
}
