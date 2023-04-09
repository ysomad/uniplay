package institution

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"
)

type repository interface {
	GetList(context.Context, getListParams) (paging.InfList[domain.Institution], error)
}

type service struct {
	institution repository
}

func NewService(r repository) *service {
	return &service{
		institution: r,
	}
}

func (s *service) GetList(ctx context.Context, p getListParams) (paging.InfList[domain.Institution], error) {
	return s.institution.GetList(ctx, p)
}
