package paging

import (
	"errors"

	"golang.org/x/exp/constraints"
)

const (
	minPageSize     = 1
	maxPageSize     = 500
	defaultPageSize = 100
)

// IntSeek is set of params for seek(keyset) pagination using integer as primary keys.
type IntSeek[T constraints.Signed] struct {
	LastID   T // id of last item in list of items
	PageSize int32
}

func NewIntSeek[T constraints.Signed](lastID *T, psize *int32) IntSeek[T] {
	var s IntSeek[T]

	if lastID != nil {
		s.LastID = *lastID
	}

	if psize == nil || *psize < minPageSize {
		s.PageSize = defaultPageSize
		return s
	}

	if *psize > maxPageSize {
		s.PageSize = maxPageSize
		return s
	}

	s.PageSize = *psize

	return s
}

// InfList is a list of objects for infinity scroll pagination.
type InfList[T any] struct {
	Items   []T
	HasNext bool
}

var (
	errInvalidArgs = errors.New("paging: length of items should not equal more than pageSize + 1")
)

func NewInfList[T any](items []T, pageSize int32) (InfList[T], error) {
	if len(items) > int(pageSize) && len(items)-int(pageSize) != 1 {
		return InfList[T]{}, errInvalidArgs
	}

	if len(items) != int(pageSize+1) {
		return InfList[T]{
			Items:   items,
			HasNext: false,
		}, nil
	}

	return InfList[T]{
		Items:   items[:len(items)-1],
		HasNext: true,
	}, nil
}

// Seek is set of params for seek(keyset) pagination using uuids
// or other non-sortable primary keys.
type Seek struct {
	PageToken token
	PageSize  int32
}
