package compendium

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
)

type compendiumRepository interface {
	GetWeaponList(context.Context) ([]domain.Weapon, error)
	GetWeaponClassList(context.Context) ([]domain.WeaponClass, error)
}
