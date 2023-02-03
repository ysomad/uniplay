package match

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
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

func (p *Postgres) DeleteByID(ctx context.Context, matchID uuid.UUID) error {
	txFunc := func(tx pgx.Tx) error {
		// player stats
		sql, args, err := p.client.Builder.
			Delete("player_match_stat").
			Where(sq.Eq{"match_id": matchID}).
			ToSql()
		if err != nil {
			return err
		}

		if _, err := tx.Exec(ctx, sql, args...); err != nil {
			return err
		}

		// weapon stats
		sql, args, err = p.client.Builder.
			Delete("player_match_weapon_stat").
			Where(sq.Eq{"match_id": matchID}).
			ToSql()
		if err != nil {
			return err
		}

		if _, err := tx.Exec(ctx, sql, args...); err != nil {
			return err
		}

		// player matches
		sql, args, err = p.client.Builder.
			Delete("player_match").
			Where(sq.Eq{"match_id": matchID}).
			ToSql()
		if err != nil {
			return err
		}

		if _, err := tx.Exec(ctx, sql, args...); err != nil {
			return err
		}

		// match
		sql, args, err = p.client.Builder.
			Delete("match").
			Where(sq.Eq{"id": matchID}).
			ToSql()
		if err != nil {
			return err
		}

		ct, err := tx.Exec(ctx, sql, args...)
		if err != nil {
			return err
		}

		if ct.RowsAffected() <= 0 {
			return domain.ErrMatchNotFound
		}

		return nil
	}

	if err := pgx.BeginTxFunc(ctx, p.client.Pool, pgx.TxOptions{}, txFunc); err != nil {
		return err
	}

	return nil
}
