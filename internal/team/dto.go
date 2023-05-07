package team

import "github.com/ysomad/uniplay/internal/pkg/paging"

type listParams struct {
	searchQuery string
	paging      paging.IntSeek[int32]
}

func newListParams(s *string, id, psize *int32) listParams {
	p := listParams{}

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

type updateParams struct {
	clanName      string
	flagCode      string
	institutionID int32
}
