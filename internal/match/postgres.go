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
		if err := p.deletePlayerStats(ctx, tx, matchID); err != nil {
			return err
		}

		if err := p.deleteWeaponStats(ctx, tx, matchID); err != nil {
			return err
		}

		if err := p.deleteMatchFromHistory(ctx, tx, matchID); err != nil {
			return err
		}

		if err := p.deleteMatch(ctx, tx, matchID); err != nil {
			return err
		}

		return nil
	}

	if err := pgx.BeginTxFunc(ctx, p.client.Pool, pgx.TxOptions{}, txFunc); err != nil {
		return err
	}

	return nil
}

// deletePlayerStats deletes all players stats associated with match.
func (p *Postgres) deletePlayerStats(ctx context.Context, tx pgx.Tx, matchID uuid.UUID) error {
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

	return nil
}

// deleteWeaponStats deletes all players weapon stats associated with match.
func (p *Postgres) deleteWeaponStats(ctx context.Context, tx pgx.Tx, matchID uuid.UUID) error {
	sql, args, err := p.client.Builder.
		Delete("player_match_weapon_stat").
		Where(sq.Eq{"match_id": matchID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// deleteMatchFromHistory deletes match from player history of matches.
func (p *Postgres) deleteMatchFromHistory(ctx context.Context, tx pgx.Tx, matchID uuid.UUID) error {
	sql, args, err := p.client.Builder.
		Delete("player_match").
		Where(sq.Eq{"match_id": matchID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// deleteMatch deletes match by id.
func (p *Postgres) deleteMatch(ctx context.Context, tx pgx.Tx, matchID uuid.UUID) error {
	sql, args, err := p.client.Builder.
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
