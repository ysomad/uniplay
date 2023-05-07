package team

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"
)

type repository interface {
	GetAll(context.Context, listParams) (paging.InfList[domain.TeamListItem], error)
}

type service struct {
	team repository
}

func NewService(r repository) *service {
	return &service{
		team: r,
	}
}

func (s *service) GetList(ctx context.Context, p listParams) (paging.InfList[domain.TeamListItem], error) {
	return s.team.GetAll(ctx, p)
}
