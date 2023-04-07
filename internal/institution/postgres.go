package institution

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/trace"

	"github.com/ysomad/uniplay/internal/domain"
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

func (p *Postgres) GetInstitutionList(ctx context.Context, f domain.InstitutionFilter) ([]domain.Institution, error) {
	b := p.client.Builder.
		Select("id, name, short_name, logo_url").
		From("institution")

	if f.ShortName != "" {
		b = b.Where(sq.Eq{"short_name": f.ShortName})
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	institutions, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Institution])
	if err != nil {
		return nil, err
	}

	return institutions, nil
}
