package player

import (
	"context"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
	"go.opentelemetry.io/otel/trace"

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

type dbPlayer struct {
	SteamID     domain.SteamID `db:"steam_id"`
	TeamID      zeronull.Int2  `db:"team_id"`
	DisplayName zeronull.Text  `db:"display_name"`
	FirstName   zeronull.Text  `db:"first_name"`
	LastName    zeronull.Text  `db:"last_name"`
	AvatarURL   zeronull.Text  `db:"avatar_url"`
}

func (p *postgres) GetAll(ctx context.Context, lp listParams) (paging.List[domain.Player], error) {
	b := p.client.Builder.
		Select("steam_id, team_id, display_name, avatar_url, first_name, last_name").
		From("player")

	filters := filter.New("steam_id", filter.TypeGT, lp.paging.LastID)

	if lp.searchQuery != "" {
		b = b.Where(sq.Expr("ts @@ phraseto_tsquery('russian', ?)", lp.searchQuery))
	}

	sql, args, err := filters.
		Attach(b).
		OrderBy("steam_id").
		OrderBy(fmt.Sprintf("ts_rank(ts, to_tsquery('russian', '%s')) DESC", lp.searchQuery)).
		Limit(uint64(lp.paging.PageSize) + 1).
		ToSql()
	if err != nil {
		return paging.List[domain.Player]{}, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return paging.List[domain.Player]{}, err
	}

	dbPlayers, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbPlayer])
	if err != nil {
		return paging.List[domain.Player]{}, err
	}

	players := make([]domain.Player, len(dbPlayers))

	for i, player := range dbPlayers {
		players[i] = domain.Player{
			SteamID:     player.SteamID,
			TeamID:      int32(player.TeamID),
			DisplayName: string(player.DisplayName),
			FirstName:   string(player.FirstName),
			LastName:    string(player.LastName),
			AvatarURL:   string(player.AvatarURL),
		}
	}

	return paging.NewList(players, lp.paging.PageSize)
}

func (p *postgres) FindBySteamID(ctx context.Context, steamID domain.SteamID) (domain.Player, error) {
	sql, args, err := p.client.Builder.
		Select("steam_id, team_id, display_name, first_name, last_name, avatar_url").
		From("player").
		Where(sq.Eq{"steam_id": steamID}).
		ToSql()
	if err != nil {
		return domain.Player{}, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return domain.Player{}, err
	}

	player, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbPlayer])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Player{}, domain.ErrPlayerNotFound
		}

		return domain.Player{}, err
	}

	return domain.Player{
		SteamID:     player.SteamID,
		TeamID:      int32(player.TeamID),
		DisplayName: string(player.DisplayName),
		FirstName:   string(player.FirstName),
		LastName:    string(player.LastName),
		AvatarURL:   string(player.AvatarURL),
	}, nil
}

func (p *postgres) UpdateBySteamID(ctx context.Context, steamID domain.SteamID, up updateParams) (domain.Player, error) {
	sql, args, err := p.client.Builder.
		Update("player").
		SetMap(map[string]any{
			"first_name": up.firstName,
			"last_name":  up.lastName,
			"avatar_url": up.avatarURL,
		}).
		Where(sq.Eq{"steam_id": steamID}).
		Suffix("RETURNING steam_id, team_id, display_name, first_name, last_name, avatar_url").
		ToSql()
	if err != nil {
		return domain.Player{}, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return domain.Player{}, err
	}

	player, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dbPlayer])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Player{}, domain.ErrPlayerNotFound
		}

		return domain.Player{}, err
	}

	return domain.Player{
		SteamID:     player.SteamID,
		TeamID:      int32(player.TeamID),
		DisplayName: string(player.DisplayName),
		FirstName:   string(player.FirstName),
		LastName:    string(player.LastName),
		AvatarURL:   string(player.AvatarURL),
	}, nil
}

type playerBaseStats struct {
	Kills              int32         `db:"total_kills"`
	HeadshotKills      int32         `db:"total_hs_kills"`
	BlindKills         int32         `db:"total_blind_kills"`
	WallbangKills      int32         `db:"total_wb_kills"`
	NoScopeKills       int32         `db:"total_noscope_kills"`
	ThroughSmokeKills  int32         `db:"total_smoke_kills"`
	Deaths             int32         `db:"total_deaths"`
	Assists            int32         `db:"total_assists"`
	FlashbangAssists   int32         `db:"total_fb_assists"`
	MVPCount           int32         `db:"total_mvp_count"`
	DamageTaken        int32         `db:"total_dmg_taken"`
	DamageDealt        int32         `db:"total_dmg_dealt"`
	GrenadeDamageDealt int32         `db:"total_grenade_dmg_dealt"`
	BlindedPlayers     int32         `db:"total_blinded_players"`
	BlindedTimes       int32         `db:"total_blinded_times"`
	BombsPlanted       int32         `db:"total_bombs_planted"`
	BombsDefused       int32         `db:"total_bombs_defused"`
	RoundsPlayed       int32         `db:"total_rounds_played"`
	MatchesPlayed      int32         `db:"total_matches_played"`
	Wins               int32         `db:"total_wins"`
	Loses              int32         `db:"total_loses"`
	TimePlayed         time.Duration `db:"total_time_played"`
}

func (p *postgres) GetBaseStats(ctx context.Context, steamID uint64) (*domain.PlayerBaseStats, error) {
	ctx, span := p.tracer.Start(ctx, "player.Postgres.GetBaseStats")
	defer span.End()

	b := p.client.Builder.
		Select(
			"sum(ps.kills) as total_kills",
			"sum(ps.hs_kills) as total_hs_kills",
			"sum(ps.blind_kills) as total_blind_kills",
			"sum(ps.wallbang_kills) as total_wb_kills",
			"sum(ps.noscope_kills) as total_noscope_kills",
			"sum(ps.through_smoke_kills) as total_smoke_kills",
			"sum(ps.deaths) as total_deaths",
			"sum(ps.assists) as total_assists",
			"sum(ps.flashbang_assists) as total_fb_assists",
			"sum(ps.mvp_count) as total_mvp_count",
			"sum(ps.damage_taken) as total_dmg_taken",
			"sum(ps.damage_dealt) as total_dmg_dealt",
			"sum(ps.grenade_damage_dealt) as total_grenade_dmg_dealt",
			"sum(ps.blinded_players) as total_blinded_players",
			"sum(ps.blinded_times) as total_blinded_times",
			"sum(ps.bombs_planted) as total_bombs_planted",
			"sum(ps.bombs_defused) as total_bombs_defused",
			"m.rounds as total_rounds_played",
			"count(m.id) as total_matches_played",
			"coalesce((case when pm.match_state = 1 then count(pm.*) end), 0) as total_wins",
			"coalesce((case when pm.match_state = -1 then count(pm.*) end), 0) as total_loses",
			"sum(m.duration) as total_time_played").
		From("player_match_stat ps").
		InnerJoin("player_match pm ON ps.player_steam_id = pm.player_steam_id").
		InnerJoin("match m ON pm.match_id = m.id").
		Where(sq.Eq{"ps.player_steam_id": steamID})

	sql, args, err := b.GroupBy("pm.match_state, m.rounds").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	stats, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[playerBaseStats])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrPlayerNotFound
		}

		return nil, err
	}

	res := domain.PlayerBaseStats(stats)

	return &res, nil
}

type weaponBaseStats struct {
	WeaponID          int16  `db:"weapon_id"`
	Weapon            string `db:"weapon"`
	Kills             int32  `db:"total_kills"`
	HeadshotKills     int32  `db:"total_hs_kills"`
	BlindKills        int32  `db:"total_blind_kills"`
	WallbangKills     int32  `db:"total_wb_kills"`
	NoScopeKills      int32  `db:"total_noscope_kills"`
	ThroughSmokeKills int32  `db:"total_smoke_kills"`
	Deaths            int32  `db:"total_deaths"`
	Assists           int32  `db:"total_assists"`
	DamageTaken       int32  `db:"total_dmg_taken"`
	DamageDealt       int32  `db:"total_dmg_dealt"`
	Shots             int32  `db:"total_shots"`

	TotalHits   int32 `db:"total_hits"`
	HeadHits    int32 `db:"total_head_hits"`
	NeckHits    int32 `db:"total_neck_hits"`
	ChestHits   int32 `db:"total_chest_hits"`
	StomachHits int32 `db:"total_stomach_hits"`
	ArmHits     int32 `db:"total_arm_hits"`
	LegHits     int32 `db:"total_leg_hits"`
}

func (p *postgres) GetWeaponBaseStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]*domain.WeaponBaseStats, error) {
	ctx, span := p.tracer.Start(ctx, "player.Postgres.GetWeaponBaseStats")
	defer span.End()

	b := p.client.Builder.
		Select(
			"ws.weapon_id",
			"w.weapon",
			"sum(ws.kills) as total_kills",
			"sum(ws.hs_kills) as total_hs_kills",
			"sum(ws.blind_kills) as total_blind_kills",
			"sum(ws.wallbang_kills) as total_wb_kills",
			"sum(ws.noscope_kills) as total_noscope_kills",
			"sum(ws.through_smoke_kills) as total_smoke_kills",
			"sum(ws.deaths) as total_deaths",
			"sum(ws.assists) as total_assists",
			"sum(ws.damage_taken) as total_dmg_taken",
			"sum(ws.damage_dealt) as total_dmg_dealt",
			"sum(ws.shots) as total_shots",
			"sum(ws.total_hits) as total_hits",
			"sum(ws.head_hits) as total_head_hits",
			"sum(ws.neck_hits) as total_neck_hits",
			"sum(ws.chest_hits) as total_chest_hits",
			"sum(ws.stomach_hits) as total_stomach_hits",
			"sum(ws.arm_hits) as total_arm_hits",
			"sum(ws.leg_hits) as total_leg_hits").
		From("player_match_weapon_stat ws").
		InnerJoin("weapon w ON ws.weapon_id = w.id").
		Where(sq.Eq{"ws.player_steam_id": steamID})

	switch {
	case f.WeaponID != 0:
		b = b.Where(sq.Eq{"ws.weapon_id": f.WeaponID})
	case f.ClassID != 0:
		b = b.Where(sq.Eq{"w.class_id": f.ClassID})
	}

	sql, args, err := b.
		GroupBy("ws.weapon_id", "w.weapon").
		OrderBy("total_kills DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	weaponStats, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[weaponBaseStats])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrPlayerNotFound
		}

		return nil, err
	}

	res := make([]*domain.WeaponBaseStats, len(weaponStats))

	for i, s := range weaponStats {
		res[i] = &domain.WeaponBaseStats{
			WeaponID:          s.WeaponID,
			Weapon:            s.Weapon,
			Kills:             s.Kills,
			HeadshotKills:     s.HeadshotKills,
			BlindKills:        s.BlindKills,
			WallbangKills:     s.WallbangKills,
			NoScopeKills:      s.NoScopeKills,
			ThroughSmokeKills: s.ThroughSmokeKills,
			Deaths:            s.Deaths,
			Assists:           s.Assists,
			DamageTaken:       s.DamageTaken,
			DamageDealt:       s.DamageDealt,
			TotalHits:         s.TotalHits,
			Shots:             s.Shots,
			HeadHits:          s.HeadHits,
			NeckHits:          s.NeckHits,
			ChestHits:         s.ChestHits,
			StomachHits:       s.StomachHits,
			ArmHits:           s.ArmHits,
			LegHits:           s.LegHits,
		}
	}

	return res, nil
}

type dbPlayerMatch struct {
	ID                 uuid.UUID         `db:"match_id"`
	Score              domain.MatchScore `db:"match_score"`
	UploadedAt         time.Time         `db:"match_uploaded_at"`
	MapName            string            `db:"map_name"`
	MapIconURL         string            `db:"map_icon_url"`
	MatchState         domain.MatchState `db:"match_state"`
	MatchScore         domain.MatchScore `db:"match_score"`
	Kills              int32             `db:"kills"`
	Deaths             int32             `db:"deaths"`
	Assists            int32             `db:"assists"`
	HeadshotPercentage float64           `db:"headshot_percentage"`
	ADR                float64           `db:"adr"`
}

func (p *postgres) GetMatchList(ctx context.Context, lp matchListParams) (paging.TokenList[*domain.PlayerMatch], error) {
	b := p.client.Builder.
		Select(
			"m.id as match_id",
			"m.score as match_score",
			"m.uploaded_at as match_uploaded_at",
			"mp.name as map_name",
			"mp.icon_url as map_icon_url",
			"pm.match_state as match_state",
			"pms.kills as kills",
			"pms.deaths as deaths",
			"pms.assists as assists",
			"pms.headshot_percentage as headshot_percentage",
			"pms.adr as adr").
		From("player_match pm").
		InnerJoin("match m ON pm.match_id = m.id").
		InnerJoin("map mp ON mp.name = m.map").
		InnerJoin("player_match_stat pms ON m.id = pms.match_id AND pms.player_steam_id = pm.player_steam_id")

	if lp.PageToken != "" {
		lastID, lastUploadedAt, err := lp.PageToken.Decode()
		if err != nil {
			return paging.TokenList[*domain.PlayerMatch]{}, err
		}

		b = b.Where(sq.Expr("(m.uploaded_at, m.id) < (?, ?)", lastUploadedAt, lastID))
	}

	sql, args, err := b.
		Where(sq.Eq{"pm.player_steam_id": lp.steamID}).
		OrderBy("m.uploaded_at DESC", "m.id DESC").
		ToSql()
	if err != nil {
		return paging.TokenList[*domain.PlayerMatch]{}, err
	}

	rows, err := p.client.Pool.Query(ctx, sql, args...)
	if err != nil {
		return paging.TokenList[*domain.PlayerMatch]{}, err
	}

	dbMatches, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[dbPlayerMatch])
	if err != nil {
		return paging.TokenList[*domain.PlayerMatch]{}, err
	}

	matches := make([]*domain.PlayerMatch, len(dbMatches))

	for i, m := range dbMatches {
		matches[i] = &domain.PlayerMatch{
			ID: m.ID,
			Map: domain.Map{
				Name:    m.MapName,
				IconURL: m.MapIconURL,
			},
			Score: m.MatchScore,
			Stats: domain.PlayerMatchStats{
				Kills:              m.Kills,
				Deaths:             m.Deaths,
				Assists:            m.Assists,
				HeadshotPercentage: m.HeadshotPercentage,
				ADR:                m.ADR,
			},
			State:      m.MatchState,
			UploadedAt: m.UploadedAt,
		}
	}

	return paging.NewTokenList(matches, lp.PageSize)
}
