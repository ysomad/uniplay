package paging

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func testInt32(i int32) *int32 { return &i }

func TestNewIntSeek(t *testing.T) {
	type args[T constraints.Signed] struct {
		lastID   *T
		pageSize *int32
	}

	type testCase[T constraints.Signed] struct {
		name string
		args args[T]
		want IntSeek[T]
	}

	tests := []testCase[int32]{
		{
			name: "success",
			args: args[int32]{
				lastID:   testInt32(69),
				pageSize: testInt32(50),
			},
			want: IntSeek[int32]{
				PageSize: 50,
				LastID:   69,
			},
		},
		{
			name: "empty lastID",
			args: args[int32]{
				pageSize: testInt32(50),
			},
			want: IntSeek[int32]{
				PageSize: 50,
				LastID:   0,
			},
		},
		{
			name: "empty page size",
			args: args[int32]{
				lastID: testInt32(69),
			},
			want: IntSeek[int32]{
				LastID:   69,
				PageSize: defaultPageSize,
			},
		},
		{
			name: "page size less than min page size",
			args: args[int32]{
				lastID:   testInt32(69),
				pageSize: testInt32(0),
			},
			want: IntSeek[int32]{
				LastID:   69,
				PageSize: defaultPageSize,
			},
		},
		{
			name: "page size greater than max page size",
			args: args[int32]{
				lastID:   testInt32(69),
				pageSize: testInt32(1111),
			},
			want: IntSeek[int32]{
				LastID:   69,
				PageSize: maxPageSize,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIntSeek(tt.args.lastID, tt.args.pageSize)
			assert.Equal(t, tt.want, got)
		})
	}
}
