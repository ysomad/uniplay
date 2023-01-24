package domain

import (
	"reflect"
	"testing"
)

func TestNewWeaponStats(t *testing.T) {
	type args struct {
		total []WeaponTotalStat
	}
	tests := []struct {
		name string
		args args
		want []WeaponStat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeaponStats(tt.args.total); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeaponStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_round(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := round(tt.args.n); got != tt.want {
				t.Errorf("round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcAccuracy(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcAccuracy(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("calcAccuracy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newWeaponAccuracyStat(t *testing.T) {
	type args struct {
		shots       int32
		headHits    int32
		chestHits   int32
		stomachHits int32
		lArmHits    int32
		rArmHits    int32
		lLegHits    int32
		rLegHits    int32
	}
	tests := []struct {
		name string
		args args
		want WeaponAccuracyStat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newWeaponAccuracyStat(tt.args.shots, tt.args.headHits, tt.args.chestHits, tt.args.stomachHits, tt.args.lArmHits, tt.args.rArmHits, tt.args.lLegHits, tt.args.rLegHits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWeaponAccuracyStat() = %v, want %v", got, tt.want)
			}
		})
	}
}
