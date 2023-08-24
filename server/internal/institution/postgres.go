package institution

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/filter"
	"github.com/ysomad/uniplay/internal/pkg/paging"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

type postgres struct {
	client *pgclient.Client
}

func NewPostgres(c *pgclient.Client) *postgres {
	return &postgres{
		client: c,
	}
}

func (pg *postgres) GetList(ctx context.Context, p listParams) (paging.List[domain.Institution], error) {
	b := pg.client.Builder.
		Select("id, type, name, short_name, city, logo_url").
		From("institution")

	filters := filter.New("id", filter.TypeGT, p.paging.LastID)

	if p.filter.City != "" {
		filters.Add("city", filter.TypeEQ, p.filter.City)
	}

	if p.filter.Type != 0 {
		filters.Add("type", filter.TypeEQ, p.filter.Type)
	}

	if p.searchQuery != "" {
		b = b.Where(sq.Expr("ts @@ phraseto_tsquery('russian', ?)", p.searchQuery))
	}

	sql, args, err := filters.
		Attach(b).
		OrderBy("id").
		OrderBy(fmt.Sprintf("ts_rank(ts, to_tsquery('russian', '%s')) DESC", p.searchQuery)).
		Limit(uint64(p.paging.PageSize) + 1).
		ToSql()
	if err != nil {
		return paging.List[domain.Institution]{}, err
	}

	rows, err := pg.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return paging.List[domain.Institution]{}, err
	}

	institutions, err := pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Institution])
	if err != nil {
		return paging.List[domain.Institution]{}, err
	}

	return paging.NewList(institutions, p.paging.PageSize)
}
