package compendium

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
)

type Service struct {
	compendium compendiumRepository
}

func NewService(r compendiumRepository) *Service {
	return &Service{
		compendium: r,
	}
}

func (s *Service) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	return s.compendium.GetWeaponList(ctx)
}

func (s *Service) GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error) {
	return s.compendium.GetWeaponClassList(ctx)
}
