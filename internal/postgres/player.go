package postgres

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/ysomad/pgxatomic"
)

type playerRepo struct {
	pool    pgxatomic.Pool
	builder sq.StatementBuilderType
}

func NewPlayerRepo(p pgxatomic.Pool, b sq.StatementBuilderType) *playerRepo {
	return &playerRepo{
		pool:    p,
		builder: b,
	}
}
