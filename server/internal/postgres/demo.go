package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/postgres/pgclient"
)

var ErrDemoAlreadyExists = errors.New("demo already exists")

type DemoStorage struct {
	*pgclient.Client
}

func (s *DemoStorage) Save(ctx context.Context, d domain.Demo) error {
	sql, args, err := s.Builder.
		Insert("demos").
		Columns("id, status, uploader, uploaded_at").
		Values(d.ID, d.Status, d.Uploader, d.UploadedAt).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := s.Pool.Exec(ctx, sql, args...); err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && pgErr.Code == pgerrcode.UniqueViolation {
			return ErrDemoAlreadyExists
		}

		return err
	}

	return nil
}
