package team

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/filter"
	"github.com/ysomad/uniplay/internal/pkg/paging"
	"github.com/ysomad/uniplay/internal/pkg/pgclient"
)

type postgres struct {
	tracer trace.Tracer
	client *pgclient.Client
}

func NewPostgres(t trace.Tracer, c *pgclient.Client) *postgres {
	return &postgres{
		tracer: t,
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

func (p *postgres) GetAll(ctx context.Context, lp listParams) (paging.InfList[domain.TeamListItem], error) {
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
		return paging.InfList[domain.TeamListItem]{}, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return paging.InfList[domain.TeamListItem]{}, err
	}

	dbTeams, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbTeamListItem])
	if err != nil {
		return paging.InfList[domain.TeamListItem]{}, err
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

	return paging.NewInfList(teams, lp.paging.PageSize)
}
