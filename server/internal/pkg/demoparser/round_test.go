package demoparser

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	"github.com/stretchr/testify/assert"
)

func Test_newRoundHistory(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want roundHistory
	}{
		{
			name: "success",
			want: roundHistory{rounds: make([]*round, 0, 50)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newRoundHistory()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_roundHistory_startRound(t *testing.T) {
	type args struct {
		ts *common.TeamState
	}
	tests := []struct {
		name string
		rh   roundHistory
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rh.start(tt.args.ts)
		})
	}
}

func Test_roundHistory_endRoundCurrRound(t *testing.T) {
	type args struct {
		e events.RoundEnd
	}
	tests := []struct {
		name    string
		rh      roundHistory
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rh.endCurrent(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("roundHistory.endRoundCurrRound() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_roundHistory_killCount(t *testing.T) {
	type args struct {
		kill events.Kill
	}
	tests := []struct {
		name    string
		rh      roundHistory
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rh.killCount(tt.args.kill); (err != nil) != tt.wantErr {
				t.Errorf("roundHistory.killCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newRound(t *testing.T) {
	type args struct {
		ts *common.TeamState
	}
	tests := []struct {
		name string
		args args
		want *round
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRound(tt.args.ts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_round_end(t *testing.T) {
	type fields struct {
		StartedAt time.Time
		TeamA     *roundTeam
		TeamB     *roundTeam
		KillFeed  []*roundKill
		Reason    events.RoundEndReason
	}
	type args struct {
		winner *common.TeamState
		loser  *common.TeamState
		reason events.RoundEndReason
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
			r := &round{
				StartedAt: tt.fields.StartedAt,
				TeamA:     tt.fields.TeamA,
				TeamB:     tt.fields.TeamB,
				KillFeed:  tt.fields.KillFeed,
				Reason:    tt.fields.Reason,
			}
			r.end(tt.args.winner, tt.args.loser, tt.args.reason)
		})
	}
}

func Test_newRoundTeam(t *testing.T) {
	type args struct {
		members []*common.Player
		side    common.Team
	}
	tests := []struct {
		name string
		args args
		want *roundTeam
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRoundTeam(tt.args.members, tt.args.side); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRoundTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundTeam_onRoundEnd(t *testing.T) {
	type fields struct {
		Survivors map[uint64]struct{}
		Cash      int
		CashSpend int
		EqValue   int
		Side      common.Team
	}
	type args struct {
		ts *common.TeamState
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
			rt := &roundTeam{
				Survivors: tt.fields.Survivors,
				Cash:      tt.fields.Cash,
				CashSpend: tt.fields.CashSpend,
				EqValue:   tt.fields.EqValue,
				Side:      tt.fields.Side,
			}
			rt.onRoundEnd(tt.args.ts)
		})
	}
}

func Test_newRoundKill(t *testing.T) {
	t.Parallel()

	tplayer1 := &common.Player{SteamID64: 123, Team: common.TeamTerrorists, IsConnected: true, Name: "tplayer1", UserID: 1}
	tplayer2 := &common.Player{SteamID64: 102, Team: common.TeamTerrorists, IsConnected: true, Name: "tplayer2", UserID: 2}
	tplayer3 := &common.Player{SteamID64: 777, Team: common.TeamTerrorists, IsConnected: true, Name: "tplayer3", UserID: 3}

	ctplayer1 := &common.Player{SteamID64: 456, Team: common.TeamCounterTerrorists, IsConnected: true, Name: "ctplayer1", UserID: 4}
	ctplayer2 := &common.Player{SteamID64: 103, Team: common.TeamCounterTerrorists, IsConnected: true, Name: "ctplayer2", UserID: 5}
	ctplayer3 := &common.Player{SteamID64: 322, Team: common.TeamCounterTerrorists, IsConnected: true, Name: "ctplayer3", UserID: 6}

	roundStartedAt := time.Now()

	type args struct {
		roundStartedAt time.Time
		kill           events.Kill
	}

	tests := []struct {
		want *roundKill
		name string
		args args
	}{
		{
			name: "1. Killer on T side, Victim on CT side, without Assister",
			args: args{
				kill: events.Kill{
					Killer: tplayer1,
					Victim: ctplayer1,
					Weapon: &common.Equipment{Type: common.EqAK47},
				},
				roundStartedAt: roundStartedAt,
			},
			want: &roundKill{
				Killer:      123,
				KillerSide:  common.TeamTerrorists,
				Victim:      456,
				Headshot:    false,
				Wallbang:    false,
				KillerBlind: false,
				SinceStart:  uint16(time.Since(roundStartedAt).Seconds()),
				Weapon:      common.EqAK47,
			},
		},
		{
			name: "2. Killer on CT side, Victim on T side, without Assister",
			args: args{
				kill: events.Kill{
					Killer:            ctplayer1,
					Victim:            tplayer1,
					IsHeadshot:        true,
					PenetratedObjects: 2,
					Weapon:            &common.Equipment{Type: common.EqM4A4},
				},
				roundStartedAt: roundStartedAt,
			},
			want: &roundKill{
				Killer:      456,
				KillerSide:  common.TeamCounterTerrorists,
				Victim:      123,
				Headshot:    true,
				Wallbang:    true,
				KillerBlind: false,
				SinceStart:  uint16(time.Since(roundStartedAt).Seconds()),
				Weapon:      common.EqM4A4,
			},
		},
		{
			name: "3. Killer on T side, Victim on CT side, Assister on T side",
			args: args{
				kill: events.Kill{
					Killer:        tplayer1,
					Victim:        ctplayer1,
					Assister:      tplayer2,
					AttackerBlind: true,
					Weapon:        &common.Equipment{Type: common.EqAWP},
				},
				roundStartedAt: roundStartedAt,
			},
			want: &roundKill{
				Killer:        123,
				KillerSide:    common.TeamTerrorists,
				Victim:        456,
				Assister:      102,
				AssisterSide:  common.TeamTerrorists,
				AssistedFlash: false,
				Headshot:      false,
				Wallbang:      false,
				KillerBlind:   true,
				SinceStart:    uint16(time.Since(roundStartedAt).Seconds()),
				Weapon:        common.EqAWP,
			},
		},
		{
			name: "4. Killer on T side, Victim on CT side, Assister on CT side",
			args: args{
				kill: events.Kill{
					Killer:        tplayer1,
					Victim:        ctplayer1,
					Assister:      ctplayer2,
					AssistedFlash: true,
					Weapon:        &common.Equipment{Type: common.EqAWP},
				},
				roundStartedAt: roundStartedAt,
			},
			want: &roundKill{
				Killer:        123,
				KillerSide:    common.TeamTerrorists,
				Victim:        456,
				Assister:      103,
				AssisterSide:  common.TeamCounterTerrorists,
				AssistedFlash: true,
				Headshot:      false,
				Wallbang:      false,
				KillerBlind:   false,
				SinceStart:    uint16(time.Since(roundStartedAt).Seconds()),
				Weapon:        common.EqAWP,
			},
		},
		{
			name: "5. Killer on CT side, Victim on T side, Assister on T side",
			args: args{
				kill: events.Kill{
					Killer:   ctplayer1,
					Victim:   tplayer1,
					Assister: tplayer2,
					Weapon:   &common.Equipment{Type: common.EqAWP},
				},
				roundStartedAt: roundStartedAt,
			},
			want: &roundKill{
				Killer:        456,
				KillerSide:    common.TeamCounterTerrorists,
				Victim:        123,
				Assister:      102,
				AssisterSide:  common.TeamTerrorists,
				AssistedFlash: false,
				Headshot:      false,
				Wallbang:      false,
				KillerBlind:   false,
				SinceStart:    uint16(time.Since(roundStartedAt).Seconds()),
				Weapon:        common.EqAWP,
			},
		},
		{
			name: "6. Killer on CT side, Victim on T side, Assister on CT side",
			args: args{
				kill: events.Kill{
					Killer:     ctplayer1,
					Victim:     tplayer1,
					Assister:   ctplayer2,
					IsHeadshot: true,
					Weapon:     &common.Equipment{Type: common.EqAWP},
				},
				roundStartedAt: roundStartedAt,
			},
			want: &roundKill{
				Killer:        456,
				KillerSide:    common.TeamCounterTerrorists,
				Victim:        123,
				Assister:      103,
				AssisterSide:  common.TeamCounterTerrorists,
				AssistedFlash: false,
				Headshot:      true,
				Wallbang:      false,
				KillerBlind:   false,
				SinceStart:    uint16(time.Since(roundStartedAt).Seconds()),
				Weapon:        common.EqAWP,
			},
		},
		{
			name: "7. Killer on T side, Victim on T side, Assister on T side",
			args: args{
				kill: events.Kill{
					Killer:            tplayer1,
					Victim:            tplayer2,
					Assister:          tplayer3,
					PenetratedObjects: 5,
					Weapon:            &common.Equipment{Type: common.EqAWP},
				},
				roundStartedAt: roundStartedAt,
			},
			want: &roundKill{
				Killer:        123,
				KillerSide:    common.TeamTerrorists,
				Victim:        102,
				Assister:      777,
				AssisterSide:  common.TeamTerrorists,
				AssistedFlash: false,
				Headshot:      false,
				Wallbang:      true,
				KillerBlind:   false,
				SinceStart:    uint16(time.Since(roundStartedAt).Seconds()),
				Weapon:        common.EqAWP,
			},
		},
		{
			name: "8. Killer on CT side, Victim on CT side, Assister on CT side",
			args: args{
				kill: events.Kill{
					Killer:        ctplayer1,
					Victim:        ctplayer2,
					Assister:      ctplayer3,
					AttackerBlind: true,
					Weapon:        &common.Equipment{Type: common.EqAWP},
				},
				roundStartedAt: roundStartedAt,
			},
			want: &roundKill{
				Killer:        456,
				KillerSide:    common.TeamCounterTerrorists,
				Victim:        103,
				Assister:      322,
				AssisterSide:  common.TeamCounterTerrorists,
				AssistedFlash: false,
				Headshot:      false,
				Wallbang:      false,
				KillerBlind:   true,
				SinceStart:    uint16(time.Since(roundStartedAt).Seconds()),
				Weapon:        common.EqAWP,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newRoundKill(tt.args.kill, tt.args.roundStartedAt)
			fmt.Printf("Test Case: %s\n", tt.name)
			fmt.Printf("Expected: %+v\n", tt.want)
			fmt.Printf("Actual  : %+v\n", got)
			assert.Equal(t, tt.want, got)
		})
	}
}
