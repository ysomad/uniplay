package service

import (
	"context"

	"github.com/ssssargsian/uniplay/internal/domain"
)

type Compendium struct {
	repo compendiumRepository
}

func NewCompendium(r compendiumRepository) *Compendium {
	return &Compendium{
		repo: r,
	}
}

func (c *Compendium) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	return c.repo.GetWeaponList(ctx)
}

func (c *Compendium) GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error) {
	return c.repo.GetWeaponClassList(ctx)
}
