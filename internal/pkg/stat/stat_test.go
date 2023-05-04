package stat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadshotPercentage(t *testing.T) {
	type args struct {
		hsKills int32
		kills   int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 kills",
			args: args{
				hsKills: 231,
				kills:   0,
			},
			want: 0,
		},
		{
			name: "0 hs kills",
			args: args{
				hsKills: 0,
				kills:   231,
			},
			want: 0,
		},
		{
			name: "success",
			args: args{
				hsKills: 35,
				kills:   35,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HeadshotPercentage(tt.args.hsKills, tt.args.kills)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestKD(t *testing.T) {
	type args struct {
		kills  int32
		deaths int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 kills",
			args: args{
				kills:  0,
				deaths: 223,
			},
			want: 0,
		},
		{
			name: "0 deaths",
			args: args{
				kills:  13,
				deaths: 0,
			},
			want: 0,
		},
		{
			name: "success",
			args: args{
				kills:  35,
				deaths: 23,
			},
			want: 1.52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KD(tt.args.kills, tt.args.deaths)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestADR(t *testing.T) {
	type args struct {
		dmg    int32
		rounds int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 rounds played",
			args: args{
				dmg:    231,
				rounds: 0,
			},
			want: 0,
		},
		{
			name: "0 damage dealt",
			args: args{
				dmg:    0,
				rounds: 23,
			},
			want: 0,
		},
		{
			name: "success",
			args: args{
				dmg:    1567,
				rounds: 5,
			},
			want: 313.4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ADR(tt.args.dmg, tt.args.rounds)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAVG(t *testing.T) {
	type args struct {
		sum int32
		num int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 num",
			args: args{
				sum: 231,
				num: 0,
			},
			want: 0,
		},
		{
			name: "0 sum",
			args: args{
				sum: 0,
				num: 23,
			},
			want: 0,
		},
		{
			name: "2 decimal place",
			args: args{
				sum: 563,
				num: 96,
			},
			want: 5.86,
		},
		{
			name: "1 decimal places",
			args: args{
				sum: 1567,
				num: 5,
			},
			want: 313.4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AVG(tt.args.sum, tt.args.num)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWinRate(t *testing.T) {
	type args struct {
		wins    int32
		matches int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 matches played",
			args: args{
				wins:    345,
				matches: 0,
			},
			want: 0,
		},
		{
			name: "0 wins",
			args: args{
				wins:    0,
				matches: 23,
			},
			want: 0,
		},
		{
			name: "50%",
			args: args{
				wins:    50,
				matches: 100,
			},
			want: 50,
		},
		{
			name: "33.33%",
			args: args{
				wins:    33,
				matches: 99,
			},
			want: 33.33,
		},
		{
			name: "58.54%",
			args: args{
				wins:    48,
				matches: 82,
			},
			want: 58.54,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WinRate(tt.args.wins, tt.args.matches)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAccuracy(t *testing.T) {
	type args struct {
		targetHits int32
		totalHits  int32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "100%",
			args: args{
				targetHits: 100,
				totalHits:  100,
			},
			want: 100,
		},
		{
			name: "50%",
			args: args{
				targetHits: 50,
				totalHits:  100,
			},
			want: 50,
		},
		{
			name: "33.33%",
			args: args{
				targetHits: 33,
				totalHits:  99,
			},
			want: 33.33,
		},
		{
			name: "100% 2",
			args: args{
				targetHits: 9999,
				totalHits:  9999,
			},
			want: 100,
		},
		{
			name: "0%",
			args: args{
				targetHits: 0,
				totalHits:  0,
			},
			want: 0,
		},
		{
			name: "0%",
			args: args{
				targetHits: -5,
				totalHits:  -5,
			},
			want: 0,
		},
		{
			name: "0 target hits",
			args: args{
				targetHits: 0,
				totalHits:  12345,
			},
			want: 0,
		},
		{
			name: "0 total hits",
			args: args{
				targetHits: 543,
				totalHits:  0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Accuracy(tt.args.targetHits, tt.args.totalHits)
			assert.Equal(t, tt.want, got)
		})
	}
}
