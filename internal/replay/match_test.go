package replay

import (
	"testing"

	"github.com/google/uuid"
	"github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs/common"
	"github.com/stretchr/testify/assert"
	"github.com/ysomad/uniplay/internal/domain"
)

func Test_replayMatch_swapTeamSides(t *testing.T) {
	t.Parallel()

	type fields struct {
		team1 replayTeam
		team2 replayTeam
	}
	tests := []struct {
		name          string
		fields        fields
		wantTeam1Side common.Team
		wantTeam2Side common.Team
	}{
		{
			name: "success",
			fields: fields{
				team1: replayTeam{_side: common.TeamCounterTerrorists},
				team2: replayTeam{_side: common.TeamTerrorists},
			},
			wantTeam1Side: common.TeamTerrorists,
			wantTeam2Side: common.TeamCounterTerrorists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &replayMatch{
				team1: tt.fields.team1,
				team2: tt.fields.team2,
			}
			m.swapTeamSides()

			assert.Equal(t, tt.wantTeam1Side, m.team1._side)
			assert.Equal(t, tt.wantTeam2Side, m.team2._side)
		})
	}
}

func Test_replayMatch_setTeamStates(t *testing.T) {
	t.Parallel()

	type fields struct {
		team1 replayTeam
		team2 replayTeam
	}
	tests := []struct {
		name           string
		fields         fields
		wantTeam1State domain.MatchState
		wantTeam2State domain.MatchState
	}{
		{
			name: "team2 win",
			fields: fields{
				team1: replayTeam{
					score:      12,
					matchState: 0,
				},
				team2: replayTeam{
					score:      0,
					matchState: 16,
				},
			},
			wantTeam1State: domain.MatchStateLose,
			wantTeam2State: domain.MatchStateWin,
		},
		{
			name: "team1 win",
			fields: fields{
				team1: replayTeam{
					score:      22,
					matchState: 0,
				},
				team2: replayTeam{
					score:      5,
					matchState: 0,
				},
			},
			wantTeam1State: domain.MatchStateWin,
			wantTeam2State: domain.MatchStateLose,
		},
		{
			name: "draw",
			fields: fields{
				team1: replayTeam{
					score:      15,
					matchState: 0,
				},
				team2: replayTeam{
					score:      15,
					matchState: 0,
				},
			},
			wantTeam1State: domain.MatchStateDraw,
			wantTeam2State: domain.MatchStateDraw,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &replayMatch{
				team1: tt.fields.team1,
				team2: tt.fields.team2,
			}
			m.setTeamStates()
		})
	}
}

func Test_replayMatch_teamPlayers(t *testing.T) {
	t.Parallel()

	type fields struct {
		id    uuid.UUID
		team1 replayTeam
		team2 replayTeam
	}

	matchID, _ := uuid.NewRandom()

	tests := []struct {
		name   string
		fields fields
		want   []teamPlayer
	}{
		{
			name: "success",
			fields: fields{
				id: matchID,
				team1: replayTeam{
					id:         1,
					clanName:   "test",
					flagCode:   "ru",
					score:      16,
					matchState: domain.MatchStateWin,
					players:    []uint64{1, 2, 3, 4, 5},
					_side:      common.TeamTerrorists,
				},
				team2: replayTeam{
					id:         2,
					clanName:   "test2",
					flagCode:   "uk",
					score:      12,
					matchState: domain.MatchStateLose,
					players:    []uint64{6, 7, 8, 9, 10},
					_side:      common.TeamCounterTerrorists,
				},
			},
			want: []teamPlayer{
				{
					steamID:    1,
					teamID:     1,
					matchID:    matchID,
					matchState: domain.MatchStateWin,
				},
				{
					steamID:    2,
					teamID:     1,
					matchID:    matchID,
					matchState: domain.MatchStateWin,
				},
				{
					steamID:    3,
					teamID:     1,
					matchID:    matchID,
					matchState: domain.MatchStateWin,
				},
				{
					steamID:    4,
					teamID:     1,
					matchID:    matchID,
					matchState: domain.MatchStateWin,
				},
				{
					steamID:    5,
					teamID:     1,
					matchID:    matchID,
					matchState: domain.MatchStateWin,
				},
				{
					steamID:    6,
					teamID:     2,
					matchID:    matchID,
					matchState: domain.MatchStateLose,
				},
				{
					steamID:    7,
					teamID:     2,
					matchID:    matchID,
					matchState: domain.MatchStateLose,
				},
				{
					steamID:    8,
					teamID:     2,
					matchID:    matchID,
					matchState: domain.MatchStateLose,
				},
				{
					steamID:    9,
					teamID:     2,
					matchID:    matchID,
					matchState: domain.MatchStateLose,
				},
				{
					steamID:    10,
					teamID:     2,
					matchID:    matchID,
					matchState: domain.MatchStateLose,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &replayMatch{
				id:    tt.fields.id,
				team1: tt.fields.team1,
				team2: tt.fields.team2,
			}
			got := m.teamPlayers()

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_newReplayTeam(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
