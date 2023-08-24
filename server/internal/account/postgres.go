package account

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/ysomad/uniplay/internal/domain"
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

func (p *postgres) Save(ctx context.Context, a *domain.Account) error {
	sql, args, err := p.client.Builder.
		Insert("account").
		Columns("id, email, password, is_verified, created_at").
		Values(a.ID, a.Email, a.Password, a.IsVerified, a.CreatedAt).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := p.client.Pool.Exec(ctx, sql, args...); err != nil {
		var e *pgconn.PgError

		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			return domain.ErrAccountEmailTaken
		}

		return err
	}

	return nil
}
