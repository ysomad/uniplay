package institution

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/trace"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

type Postgres struct {
	tracer trace.Tracer
	client *pgclient.Client
}

func NewPostgres(t trace.Tracer, c *pgclient.Client) *Postgres {
	return &Postgres{
		tracer: t,
		client: c,
	}
}

func (pg *Postgres) GetInstitutionList(ctx context.Context, f domain.InstitutionFilter, p paging.IntSeek[int32]) (paging.InfList[domain.Institution], error) {
	b := pg.client.Builder.
		Select("id, name, short_name, logo_url").
		From("institution").
		Where(sq.Gt{"id": p.LastID})

	if f.ShortName != "" {
		b = b.Where(sq.Eq{"short_name": f.ShortName})
	}

	sql, args, err := b.
		Limit(uint64(p.PageSize) + 1).
		OrderBy("id").
		ToSql()
	if err != nil {
		return paging.InfList[domain.Institution]{}, err
	}

	rows, err := pg.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return paging.InfList[domain.Institution]{}, err
	}

	institutions, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Institution])
	if err != nil {
		return paging.InfList[domain.Institution]{}, err
	}

	return paging.NewInfList(institutions, p.PageSize), nil
}
