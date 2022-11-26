package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ysomad/pgxatomic"
	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ssssargsian/uniplay/internal/inmemory"
)

type compendiumRepo struct {
	log     *zap.Logger
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType

	weaponCache      *inmemory.WeaponCache
	weaponClassCache *inmemory.WeaponClassCache
}

func NewCompendiumRepo(l *zap.Logger, p pgxatomic.Pool, b sq.StatementBuilderType, wc *inmemory.WeaponCache, wcs *inmemory.WeaponClassCache) *compendiumRepo {
	return &compendiumRepo{
		log:              l,
		pool:             p,
		builder:          b,
		weaponCache:      wc,
		weaponClassCache: wcs,
	}
}

func (r *compendiumRepo) GetWeaponList(ctx context.Context) ([]domain.Weapon, error) {
	weapons := r.weaponCache.Get()
	if len(weapons) != 0 {
		return weapons, nil
	}

	sql, args, err := r.builder.
		Select("DISTINCT weapon_name, weapon_class").
		From("weapon_metric").
		OrderBy("weapon_name").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var w domain.Weapon

		if err := rows.Scan(&w.Name, &w.ClassID); err != nil {
			return nil, err
		}

		w.ClassName = w.ClassID.String()
		weapons = append(weapons, w)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	r.weaponCache.Save(weapons)
	return weapons, nil
}

func (r *compendiumRepo) GetWeaponClassList(ctx context.Context) ([]domain.WeaponClass, error) {
	classes := r.weaponClassCache.Get()
	if len(classes) != 0 {
		return classes, nil
	}

	sql, args, err := r.builder.
		Select("weapon_class").
		From("weapon_metric").
		GroupBy("weapon_class").
		OrderBy("weapon_class").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c domain.WeaponClass

		if err := rows.Scan(&c.ID); err != nil {
			return nil, err
		}

		c.Name = c.ID.String()
		classes = append(classes, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	r.weaponClassCache.Save(classes)
	return classes, nil
}
