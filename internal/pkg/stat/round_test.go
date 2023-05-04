package stat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_round(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1",
			args: args{63.312312123},
			want: 63.31,
		},
		{
			name: "2",
			args: args{0},
			want: 00.00,
		},
		{
			name: "3",
			args: args{999.9999999},
			want: 1000,
		},
		{
			name: "4",
			args: args{123.567},
			want: 123.57,
		},
		{
			name: "5",
			args: args{0.99999},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := round(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
