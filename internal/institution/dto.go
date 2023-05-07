package institution

import (
	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"
)

type listParams struct {
	searchQuery string
	filter      domain.InstitutionFilter
	paging      paging.IntSeek[int32]
}

func newListParams(s *string, f domain.InstitutionFilter, id, psize *int32) listParams {
	p := listParams{
		filter: f,
	}

	if s != nil {
		p.searchQuery = *s
	}

	var lastID, pageSize int32

	if id != nil {
		lastID = *id
	}

	if psize != nil {
		pageSize = *psize
	}

	p.paging = paging.NewIntSeek(lastID, pageSize)

	return p
}
