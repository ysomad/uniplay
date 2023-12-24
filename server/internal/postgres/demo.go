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
	"github.com/ysomad/uniplay/server/internal/postgres/pgclient"
	pgmodel "github.com/ysomad/uniplay/server/internal/postgres/pgmodel"
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
		Columns("id, status, identity_id, uploaded_at").
		Values(d.ID, d.Status, d.IdentityID, d.UploadedAt).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := s.Pool.Exec(ctx, sql, args...); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("%w, demo_id = %s", ErrDemoAlreadyExists, d.ID)
		}

		return err
	}

	return nil
}

func (s *DemoStorage) GetOne(ctx context.Context, demoID, identityID string) (domain.Demo, error) {
	sql, args, err := s.Builder.
		Select("id, identity_id, status, reason, uploaded_at, processed_at").
		From(demoTable).
		Where(sq.And{
			sq.Eq{"id": demoID},
			sq.Eq{"identity_id": identityID},
		}).
		ToSql()
	if err != nil {
		return domain.Demo{}, err
	}

	rows, err := s.Pool.Query(ctx, sql, args...)
	if err != nil {
		return domain.Demo{}, fmt.Errorf("%w, demo_id = %s", err, demoID)
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

func (s *DemoStorage) GetAll(ctx context.Context, identityID string, status domain.DemoStatus) ([]domain.Demo, error) {
	b := s.Builder.
		Select("id, identity_id, status, reason, uploaded_at, processed_at").
		From(demoTable).
		Where(sq.Eq{"identity_id": identityID})

	if status.Valid() {
		b = b.Where(sq.Eq{"status": status})
	}

	sql, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	pgdemos, err := pgx.CollectRows[pgmodel.Demo](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	demos := make([]domain.Demo, len(pgdemos))

	for i, d := range pgdemos {
		demos[i] = domain.Demo{
			UploadedAt:  d.UploadedAt,
			ProcessedAt: time.Time(d.ProcessedAt),
			Status:      d.Status,
			Reason:      string(d.Reason),
			IdentityID:  d.IdentityID,
			ID:          d.ID,
		}
	}

	return demos, nil
}
