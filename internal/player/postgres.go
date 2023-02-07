package player

import (
	"context"
	"errors"
	"time"

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
	Draws              int32         `db:"total_draws"`
	TimePlayed         time.Duration `db:"total_time_played"`
}

func (p *Postgres) GetBaseStats(ctx context.Context, steamID uint64, f domain.PlayerStatsFilter) (*domain.PlayerBaseStats, error) {
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
			"sum(m.team1_score) + sum(m.team2_score) as total_rounds_played",
			"count(m.id) as total_matches_played",
			"coalesce((case when pm.match_state = 1 then count(pm.*) end), 0) as total_wins",
			"coalesce((case when pm.match_state = -1 then count(pm.*) end), 0) as total_loses",
			"coalesce((case when pm.match_state = 0 then count(pm.*) end), 0) as total_draws",
			"sum(m.duration) as total_time_played").
		From("player_match_stat ps").
		InnerJoin("player_match pm ON ps.player_steam_id = pm.player_steam_id").
		InnerJoin("match m ON pm.match_id = m.id").
		Where(sq.Eq{"ps.player_steam_id": steamID})

	if (f.MatchID != uuid.UUID{}) {
		b = b.Where(sq.Eq{"ps.match_id": f.MatchID})
	}

	sql, args, err := b.GroupBy("pm.match_state").ToSql()
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
	WeaponID          int32  `db:"weapon_id"`
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
	HeadHits          int32  `db:"total_head_hits"`
	NeckHits          int32  `db:"total_neck_hits"`
	ChestHits         int32  `db:"total_chest_hits"`
	StomachHits       int32  `db:"total_stomach_hits"`
	LeftArmHits       int32  `db:"total_l_arm_hits"`
	RightArmHits      int32  `db:"total_r_arm_hits"`
	LeftLegHits       int32  `db:"total_l_leg_hits"`
	RightLegHits      int32  `db:"total_r_leg_hits"`
}

func (p *Postgres) GetWeaponBaseStats(ctx context.Context, steamID uint64, f domain.WeaponStatsFilter) ([]*domain.WeaponBaseStats, error) {
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
			"sum(ws.head_hits) as total_head_hits",
			"sum(ws.neck_hits) as total_neck_hits",
			"sum(ws.chest_hits) as total_chest_hits",
			"sum(ws.stomach_hits) as total_stomach_hits",
			"sum(ws.left_arm_hits) as total_l_arm_hits",
			"sum(ws.right_arm_hits) as total_r_arm_hits",
			"sum(ws.left_leg_hits) as total_l_leg_hits",
			"sum(ws.right_leg_hits) as total_r_leg_hits").
		From("player_match_weapon_stat ws").
		InnerJoin("weapon w ON ws.weapon_id = w.id").
		Where(sq.Eq{"ws.player_steam_id": steamID})

	switch {
	case f.WeaponID != nil:
		b = b.Where(sq.Eq{"ws.weapon_id": f.WeaponID})
	case f.ClassID != nil:
		b = b.Where(sq.Eq{"w.class_id": f.ClassID})
	case f.MatchID != uuid.UUID{}:
		b = b.Where(sq.Eq{"ws.match_id": f.MatchID})
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
			Shots:             s.Shots,
			HeadHits:          s.HeadHits,
			NeckHits:          s.NeckHits,
			ChestHits:         s.ChestHits,
			StomachHits:       s.StomachHits,
			LeftArmHits:       s.LeftArmHits,
			RightArmHits:      s.RightArmHits,
			LeftLegHits:       s.LeftLegHits,
			RightLegHits:      s.RightLegHits,
		}
	}

	return res, nil
}
