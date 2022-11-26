package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type compendium struct {
	repo compendiumRepository
}

func NewCompendium(r compendiumRepository) *compendium {
	return &compendium{
		repo: r,
	}
}

func (c *compendium) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	// return c.repo.GetWeaponList(ctx)
	return nil, nil
}

func (c *compendium) GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error) {
	// return c.repo.GetWeaponClassList(ctx)
	return nil, nil
}
