package institution

import (
	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"
)

type getListParams struct {
	searchQuery string
	filter      domain.InstitutionFilter
	paging      paging.IntSeek[int32]
}

func newGetListParams(s *string, f domain.InstitutionFilter, p paging.IntSeek[int32]) getListParams {
	if s != nil {
		return getListParams{
			searchQuery: *s,
			filter:      f,
			paging:      p,
		}
	}

	return getListParams{
		filter: f,
		paging: p,
	}
}
