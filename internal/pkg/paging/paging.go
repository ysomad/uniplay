package paging

import (
	"time"

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

// Seek is set of params for seek(keyset) pagination using uuids
// or other non-sortable primary keys.
type Seek struct {
	PageToken token
	PageSize  int32
}

func (s Seek) DecodedToken() (string, time.Time, error) { return s.PageToken.Decode() }

// InfList is a list of objects for infinity scroll pagination.
type InfList[T any] struct {
	Items   []T
	HasNext bool
}

func NewInfList[T any](items []T, pageSize int32) InfList[T] {
	itemCount := len(items)
	hasNext := itemCount == int(pageSize+1)

	if hasNext {
		items = items[:itemCount-1]
	}

	return InfList[T]{
		Items:   items,
		HasNext: hasNext,
	}
}
