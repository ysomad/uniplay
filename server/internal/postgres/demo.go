package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/ysomad/uniplay/server/internal/domain"
	pgmodel "github.com/ysomad/uniplay/server/internal/postgres/model"
	"github.com/ysomad/uniplay/server/internal/postgres/pgclient"
)

const demoTable = "demos"

var (
	ErrDemoAlreadyExists = errors.New("demo already exists")
	ErrDemoNotFound      = errors.New("demo not found")
)

type DemoStorage struct {
	*pgclient.Client
}

func NewDemoStorage(c *pgclient.Client) DemoStorage {
	return DemoStorage{c}
}

func (s *DemoStorage) Save(ctx context.Context, d domain.Demo) error {
	sql, args, err := s.Builder.
		Insert(demoTable).
		Columns("id, status, uploader, uploaded_at").
		Values(d.ID, d.Status, d.IdentityID, d.UploadedAt).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := s.Pool.Exec(ctx, sql, args...); err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("%w, demo_id = %s", ErrDemoAlreadyExists, d.ID)
		}

		return err
	}

	return nil
}

func (s *DemoStorage) GetOne(ctx context.Context, id string) (domain.Demo, error) {
	sql, args, err := s.Builder.
		Select("id, identity_id, status, reason, uploaded_at, processed_at").
		From(demoTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return domain.Demo{}, err
	}

	rows, err := s.Pool.Query(ctx, sql, args...)
	if err != nil {
		return domain.Demo{}, fmt.Errorf("%w, demo_id = %s", err, id)
	}

	res, err := pgx.CollectOneRow[pgmodel.Demo](rows, pgx.RowToStructByName)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Demo{}, ErrDemoNotFound
		}

		return domain.Demo{}, err
	}

	return domain.Demo{
		ID:          res.ID,
		IdentityID:  res.IdentityID,
		Status:      res.Status,
		Reason:      string(res.Reason),
		UploadedAt:  res.UploadedAt,
		ProcessedAt: time.Time(res.ProcessedAt),
	}, nil
}
