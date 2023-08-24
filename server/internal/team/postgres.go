package team

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype/zeronull"

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

type dbTeamListItem struct {
	ID            int32         `db:"team_id"`
	ClanName      string        `db:"clan_name"`
	FlagCode      zeronull.Text `db:"flag_code"`
	InstID        zeronull.Int4 `db:"inst_id"`
	InstShortName zeronull.Text `db:"inst_short_name"`
	InstCity      zeronull.Text `db:"inst_city"`
	InstType      zeronull.Int2 `db:"inst_type"`
	InstLogoURL   zeronull.Text `db:"inst_logo_url"`
}

func (p *postgres) GetAll(ctx context.Context, lp listParams) (paging.List[domain.TeamListItem], error) {
	b := p.client.Builder.
		Select(
			"t.id as team_id",
			"t.clan_name as clan_name",
			"t.flag_code as flag_code",
			"i.id as inst_id",
			"i.short_name as inst_short_name",
			"i.city as inst_city",
			"i.type as inst_type",
			"i.logo_url as inst_logo_url",
		).
		From("team t").
		LeftJoin("institution i ON t.institution_id = i.id")

	filters := filter.New("t.id", filter.TypeGT, lp.paging.LastID)

	if lp.searchQuery != "" {
		b = b.Where(sq.Expr("t.ts @@ phraseto_tsquery('russian', ?)", lp.searchQuery))
	}

	sql, args, err := filters.
		Attach(b).
		OrderBy("t.id").
		OrderBy(fmt.Sprintf("ts_rank(t.ts, to_tsquery('russian', '%s')) DESC", lp.searchQuery)).
		Limit(uint64(lp.paging.PageSize) + 1).
		ToSql()
	if err != nil {
		return paging.List[domain.TeamListItem]{}, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return paging.List[domain.TeamListItem]{}, err
	}

	dbTeams, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbTeamListItem])
	if err != nil {
		return paging.List[domain.TeamListItem]{}, err
	}

	teams := make([]domain.TeamListItem, len(dbTeams))

	for i, t := range dbTeams {
		teams[i] = domain.TeamListItem{
			ID:            t.ID,
			ClanName:      t.ClanName,
			FlagCode:      string(t.FlagCode),
			InstID:        int32(t.InstID),
			InstShortName: string(t.InstShortName),
			InstCity:      string(t.InstCity),
			InstType:      domain.InstType(t.InstType),
			InstLogoURL:   string(t.InstLogoURL),
		}
	}

	return paging.NewList(teams, lp.paging.PageSize)
}

type dbTeamPlayer struct {
	SteamID     domain.SteamID `db:"steam_id"`
	DisplayName zeronull.Text  `db:"display_name"`
	FirstName   zeronull.Text  `db:"first_name"`
	LastName    zeronull.Text  `db:"last_name"`
	AvatarURL   zeronull.Text  `db:"avatar_url"`
	IsCaptain   bool           `db:"is_captain"`
}

func (p *postgres) GetPlayers(ctx context.Context, teamID int32) ([]domain.TeamPlayer, error) {
	sql, args, err := p.client.Builder.
		Select(
			"tp.player_steam_id as steam_id",
			"p.display_name as display_name",
			"p.first_name as first_name",
			"p.last_name as last_name",
			"p.avatar_url as avatar_url",
			"tp.is_captain as is_captain",
		).
		From("team_player tp").
		InnerJoin("player p ON tp.player_steam_id = p.steam_id").
		Where(sq.Eq{"tp.team_id": teamID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	dbPlayers, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbTeamPlayer])
	if err != nil {
		return nil, err
	}

	if len(dbPlayers) < 1 {
		return nil, domain.ErrTeamNotFound
	}

	players := make([]domain.TeamPlayer, len(dbPlayers))

	for i, pl := range dbPlayers {
		players[i] = domain.TeamPlayer{
			SteamID:     pl.SteamID,
			DisplayName: string(pl.DisplayName),
			FirstName:   string(pl.FirstName),
			LastName:    string(pl.LastName),
			AvatarURL:   string(pl.AvatarURL),
			IsCaptain:   pl.IsCaptain,
		}
	}

	return players, nil
}

func (p *postgres) Update(ctx context.Context, teamID int32, up updateParams) (domain.Team, error) {
	sql, args, err := p.client.Builder.
		Update("team").
		SetMap(map[string]any{
			"clan_name":      up.clanName,
			"flag_code":      up.flagCode,
			"institution_id": up.institutionID,
		}).
		Where(sq.Eq{"id": teamID}).
		Suffix("RETURNING id, institution_id, clan_name, flag_code").
		ToSql()
	if err != nil {
		return domain.Team{}, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return domain.Team{}, err
	}

	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Team])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Team{}, domain.ErrTeamNotFound
		}

		var e *pgconn.PgError
		if errors.As(err, &e) && e.Code == pgerrcode.UniqueViolation {
			return domain.Team{}, domain.ErrTeamClanNameTaken
		}

		return domain.Team{}, err
	}

	return t, nil
}

func (p *postgres) SetCaptain(ctx context.Context, teamID int32, steamID domain.SteamID) error {
	err := pgx.BeginTxFunc(ctx, p.client.Pool, pgx.TxOptions{}, func(tx pgx.Tx) error {
		// unset captain
		sql, args, err := p.client.Builder.
			Update("team_player").
			Set("is_captain", false).
			Where(sq.And{
				sq.Eq{"team_id": teamID},
				sq.Eq{"is_captain": true},
			}).
			ToSql()
		if err != nil {
			return err
		}

		if _, err = tx.Exec(ctx, sql, args...); err != nil {
			return err
		}

		// set new captain
		sql, args, err = p.client.Builder.
			Update("team_player").
			Set("is_captain", true).
			Where(sq.And{
				sq.Eq{"team_id": teamID},
				sq.Eq{"player_steam_id": steamID},
			}).
			ToSql()
		if err != nil {
			return err
		}

		if _, err = tx.Exec(ctx, sql, args...); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}