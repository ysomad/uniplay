package institution

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"
)

type institutionRepository interface {
	GetInstitutionList(context.Context, domain.InstitutionFilter, paging.IntSeek[int32]) (paging.InfList[domain.Institution], error)
}

type Service struct {
	institution institutionRepository
}

func NewService(r institutionRepository) *Service {
	return &Service{
		institution: r,
	}
}

func (s *Service) GetInstitutionList(ctx context.Context, f domain.InstitutionFilter, p paging.IntSeek[int32]) (paging.InfList[domain.Institution], error) {
	return s.institution.GetInstitutionList(ctx, f, p)
}
