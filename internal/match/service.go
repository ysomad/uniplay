package match

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/google/uuid"
	"github.com/ysomad/uniplay/internal/domain"
)

type matchRepository interface {
	CreateWithStats(context.Context, *replayMatch, []*playerStat, []*weaponStat) error
	Exists(ctx context.Context, matchID uuid.UUID) (found bool, err error)
	DeleteByID(ctx context.Context, matchID uuid.UUID) error
	GetScoreBoardRowsByID(ctx context.Context, matchID uuid.UUID) ([]*matchScoreBoardRow, error)
}

type Service struct {
	tracer trace.Tracer
	match  matchRepository
}

func NewService(t trace.Tracer, m matchRepository) *Service {
	return &Service{
		tracer: t,
		match:  m,
	}
}

func (s *Service) CreateFromReplay(ctx context.Context, r replay) (uuid.UUID, error) {
	ctx, span := s.tracer.Start(ctx, "match.Service.CreateFromReplay")
	defer span.End()

	p := newParser(r)
	defer p.close()

	if err := p.parseReplayHeader(); err != nil {
		return uuid.Nil, err
	}

	matchExists, err := s.match.Exists(ctx, p.match.id)
	if err != nil {
		return uuid.Nil, err
	}

	if matchExists {
		return uuid.Nil, domain.ErrMatchAlreadyExist
	}

	match, playerStats, weaponStats, err := p.collectStats(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	if err := s.match.CreateWithStats(ctx, match, playerStats, weaponStats); err != nil {
		return uuid.Nil, err
	}

	return p.match.id, nil
}

// DeleteByID deletes match and all stats associated with it, including player match history.
func (s *Service) DeleteByID(ctx context.Context, matchID uuid.UUID) error {
	return s.match.DeleteByID(ctx, matchID)
}

func (s *Service) GetByID(ctx context.Context, matchID uuid.UUID) (domain.Match, error) {
	ctx, span := s.tracer.Start(ctx, "match.Service.GetByID")
	defer span.End()

	rows, err := s.match.GetScoreBoardRowsByID(ctx, matchID)
	if err != nil {
		return domain.Match{}, err
	}

	if len(rows) < 1 {
		return domain.Match{}, domain.ErrMatchNotFound
	}

	match := domain.Match{
		ID: rows[0].MatchID,
		Map: domain.Map{
			ID:           rows[0].MapID,
			Name:         rows[0].MapName,
			InternalName: rows[0].MapInternalName,
			IconURL:      rows[0].MapIconURL,
		},
		RoundsPlayed: rows[0].RoundsPlayed,
		Duration:     rows[0].MatchDuration,
		UploadedAt:   rows[0].MatchUploadedAt,
	}

	for _, row := range rows {
		r := domain.NewMatchScoreBoardRow(
			row.SteamID,
			row.PlayerName,
			row.Kills,
			row.HeadshotKills,
			row.Deaths,
			row.Assists,
			row.MVPCount,
			row.DamageDealt,
			row.RoundsPlayed,
		)

		t := domain.NewMatchTeam(
			row.TeamID,
			row.TeamScore,
			row.TeamMatchState,
			row.TeamName,
			row.TeamFlagCode,
		)

		// инициализировать команд матча
		if match.Team1 == nil {
			match.Team1 = t
			match.Team1.ScoreBoard = append(match.Team1.ScoreBoard, &r)

			continue
		}

		if match.Team2 == nil {
			match.Team2 = t
			match.Team2.ScoreBoard = append(match.Team2.ScoreBoard, &r)

			continue
		}

		// Добавить строку в таблицу соответствующей команды
		switch row.TeamID {
		case match.Team1.ID:
			match.Team1.ScoreBoard = append(match.Team1.ScoreBoard, &r)
		case match.Team2.ID:
			match.Team2.ScoreBoard = append(match.Team2.ScoreBoard, &r)
		}
	}

	return match, nil
}
