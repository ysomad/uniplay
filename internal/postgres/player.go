package postgres

import (
	"time"

	"go.uber.org/zap"

	"github.com/ssssargsian/uniplay/internal/pkg/pgclient"
)

type playerRepo struct {
	log    *zap.Logger
	client *pgclient.Client
}

func NewPlayerRepo(l *zap.Logger, c *pgclient.Client) *playerRepo {
	return &playerRepo{
		log:    l,
		client: c,
	}
}

type player struct {
	SteamID      uint64
	TeamName     *string
	TeamFlagCode *string
	CreateTime   time.Time
	UpdateTime   time.Time
}

// func (p player) toDomainModel() domain.Player {
// 	dp := domain.Player{
// 		SteamID:    p.SteamID,
// 		CreateTime: p.CreateTime,
// 		UpdateTime: p.UpdateTime,
// 	}
// 	if p.TeamName != nil {
// 		dp.TeamName = *p.TeamName
// 	}
// 	if p.TeamFlagCode != nil {
// 		dp.TeamFlagCode = *p.TeamFlagCode
// 	}
// 	return dp
// }

// func (r *playerRepo) FindBySteamID(ctx context.Context, steamID uint64) (domain.Player, error) {
// 	sql, args, err := r.builder.
// 		Select("p.steam_id, p.team_name, t.flag_code, p.create_time, p.update_time").
// 		From("player p").
// 		LeftJoin("team t ON p.team_name = t.name").
// 		Where(sq.Eq{"steam_id": steamID}).
// 		ToSql()
// 	if err != nil {
// 		return domain.Player{}, err
// 	}

// 	rows, err := r.pool.Query(ctx, sql, args...)
// 	if err != nil {
// 		return domain.Player{}, err
// 	}

// 	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[player])
// 	if err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			return domain.Player{}, domain.ErrPlayerNotFound
// 		}

// 		return domain.Player{}, err
// 	}

// 	return p.toDomainModel(), nil
// }
