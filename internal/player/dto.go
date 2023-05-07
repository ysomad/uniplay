package player

import (
	"strconv"

	"github.com/ysomad/uniplay/internal/pkg/paging"
)

type updateParams struct {
	lastName  string
	firstName string
	avatarURL string
}

type listParams struct {
	searchQuery string
	paging      paging.IntSeek[uint64]
}

func newListParams(search, steamID *string, psize *int32) (lp listParams, err error) {
	if search != nil {
		lp.searchQuery = *search
	}

	var (
		steamID64 uint64
		pageSize  int32
	)

	if steamID != nil {
		steamID64, err = strconv.ParseUint(*steamID, 10, 64)
		if err != nil {
			return listParams{}, err
		}
	}

	if psize != nil {
		pageSize = *psize
	}

	lp.paging = paging.NewIntSeek(steamID64, pageSize)

	return lp, nil
}
