package player

import (
	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"
)

type updateParams struct {
	lastName  string
	firstName string
	avatarURL string
}

type listParams struct {
	searchQuery string
	paging      paging.IntSeek[domain.SteamID]
}

func newListParams(search, steamID *string, psize *int32) (lp listParams, err error) {
	if search != nil {
		lp.searchQuery = *search
	}

	var (
		steamID64 domain.SteamID
		pageSize  int32
	)

	if steamID != nil {
		steamID64, err = domain.NewSteamID(*steamID)
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

type matchListParams struct {
	steamID domain.SteamID
	paging.Seek
}

func newMatchListParams(steamID string, token *string, psize *int32) (lp matchListParams, err error) {
	lp.steamID, err = domain.NewSteamID(steamID)
	if err != nil {
		return matchListParams{}, err
	}

	if token != nil {
		lp.PageToken = paging.Token(*token)
	}

	if psize != nil {
		lp.PageSize = *psize
	}

	lp.Seek = paging.NewSeek(lp.PageToken, lp.PageSize)

	return lp, nil
}
