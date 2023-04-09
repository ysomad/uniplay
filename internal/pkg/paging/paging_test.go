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

	type test[T constraints.Signed] struct {
		name string
		args args[T]
		want IntSeek[T]
	}

	tests := []test[int32]{
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

func TestNewInfList(t *testing.T) {
	type args[T any] struct {
		items    []T
		pageSize int32
	}

	type test[T any] struct {
		name    string
		args    args[T]
		want    InfList[T]
		wantErr bool
	}

	tests := []test[string]{
		{
			name: "has next page 1",
			args: args[string]{
				items: []string{
					"item 1",
					"item 2",
					"item 3",
					"item 4",
					"item 5",
					"item 6",
				},
				pageSize: 5,
			},
			want: InfList[string]{
				Items: []string{
					"item 1",
					"item 2",
					"item 3",
					"item 4",
					"item 5",
				},
				HasNext: true,
			},
			wantErr: false,
		},
		{
			name: "has next page 2",
			args: args[string]{
				items: []string{
					"item 1",
					"item 2",
					"item 3",
					"item 4",
					"item 5",
					"item 6",
					"item 7",
					"item 8",
					"item 9",
					"item 10",
					"item 11",
				},
				pageSize: 10,
			},
			want: InfList[string]{
				Items: []string{
					"item 1",
					"item 2",
					"item 3",
					"item 4",
					"item 5",
					"item 6",
					"item 7",
					"item 8",
					"item 9",
					"item 10",
				},
				HasNext: true,
			},
			wantErr: false,
		},
		{
			name: "does not have next page 1",
			args: args[string]{
				items: []string{
					"item 1",
					"item 2",
					"item 3",
					"item 4",
					"item 5",
				},
				pageSize: 5,
			},
			want: InfList[string]{
				Items: []string{
					"item 1",
					"item 2",
					"item 3",
					"item 4",
					"item 5",
				},
				HasNext: false,
			},
			wantErr: false,
		},
		{
			name: "does not have next page 2",
			args: args[string]{
				items: []string{
					"item 1",
					"item 2",
					"item 3",
				},
				pageSize: 5,
			},
			want: InfList[string]{
				Items: []string{
					"item 1",
					"item 2",
					"item 3",
				},
				HasNext: false,
			},
			wantErr: false,
		},
		{
			name: "does not have next page 2",
			args: args[string]{
				items: []string{
					"item 1",
					"item 2",
					"item 3",
				},
				pageSize: 5,
			},
			want: InfList[string]{
				Items: []string{
					"item 1",
					"item 2",
					"item 3",
				},
				HasNext: false,
			},
			wantErr: false,
		},
		{
			name: "items has length more than pageSize + 1",
			args: args[string]{
				items: []string{
					"item 1",
					"item 2",
					"item 3",
					"item 4",
					"item 5",
					"item 6",
					"item 7",
					"item 8",
					"item 9",
					"item 10",
					"item 11",
					"item 12",
				},
				pageSize: 10,
			},
			want:    InfList[string]{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInfList(tt.args.items, tt.args.pageSize)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, (err != nil))
		})
	}
}
