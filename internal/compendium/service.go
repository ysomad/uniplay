package compendium

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
)

type repository interface {
	GetWeaponList(context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(context.Context) ([]domain.WeaponClass, error)
	GetMapList(context.Context) ([]domain.Map, error)
	GetCityList(ctx context.Context, searchQuery string) ([]domain.City, error)
}

type service struct {
	compendium repository
}

func NewService(r repository) *service {
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

func (s *service) GetMapList(ctx context.Context) ([]domain.Map, error) {
	return s.compendium.GetMapList(ctx)
}

func (s *service) GetCityList(ctx context.Context, searchQuerty string) ([]domain.City, error) {
	return s.compendium.GetCityList(ctx, searchQuerty)
}
