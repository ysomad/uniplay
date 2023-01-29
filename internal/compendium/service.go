package compendium

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
)

type service struct {
	compendium compendiumRepository
}

func NewService(r compendiumRepository) *service {
	return &service{
		compendium: r,
	}
}

func (s *service) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	return s.compendium.GetWeaponList(ctx)
}

func (s *service) GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error) {
	return s.compendium.GetWeaponClassList(ctx)
}
