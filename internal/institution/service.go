package institution

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
)

type institutionRepository interface {
	GetInstitutionList(context.Context, domain.InstitutionFilter) ([]domain.Institution, error)
}

type Service struct {
	institution institutionRepository
}

func NewService(r institutionRepository) *Service {
	return &Service{
		institution: r,
	}
}

func (s *Service) GetInstitutionList(ctx context.Context, f domain.InstitutionFilter) ([]domain.Institution, error) {
	return s.institution.GetInstitutionList(ctx, f)
}
