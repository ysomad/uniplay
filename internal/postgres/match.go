package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/ssssargsian/uniplay/internal/domain"
	"github.com/ysomad/pgxatomic"
)

type matchRepo struct {
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType
}

func NewMatchRepo(p pgxatomic.Pool, b sq.StatementBuilderType) *matchRepo {
	return &matchRepo{
		pool:    p,
		builder: b,
	}
}

func (r *matchRepo) Save(ctx context.Context, m *domain.Match) error {
	// insert teams
	sql, args, err := r.builder.
		Insert("team").
		Columns("name, flag_code, create_time, update_time").
		Values(m.Team1.Name, m.Team1.FlagCode, m.UploadTime, m.UploadTime).
		Values(m.Team2.Name, m.Team2.FlagCode, m.UploadTime, m.UploadTime).
		Suffix("ON CONFLICT DO NOTHING").
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	// insert match
	sql, args, err = r.builder.
		Insert("match").
		Columns("id, map_name, team1_name, team1_score, team2_name, team2_score, duration, upload_time").
		Values(m.ID, m.MapName, m.Team1.Name, m.Team1.Score, m.Team2.Name, m.Team2.Score, m.Duration, m.UploadTime).
		ToSql()
	if err != nil {
		return err
	}

	if _, err = r.pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}
