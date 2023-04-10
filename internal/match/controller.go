package match

import (
	"errors"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	"github.com/ysomad/uniplay/internal/gen/swagger2/models"
	matchGen "github.com/ysomad/uniplay/internal/gen/swagger2/restapi/operations/match"

	"github.com/ysomad/uniplay/internal/domain"
)

type Controller struct {
	match *service
}

func NewController(s *service) *Controller {
	return &Controller{
		match: s,
	}
}

const msgReplayFileNotFound = "replay file not found in request"

func (c *Controller) CreateMatch(p matchGen.CreateMatchParams) matchGen.CreateMatchResponder {
	formFile, ok := p.Replay.(*runtime.File)
	if !ok {
		return matchGen.NewCreateMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: msgReplayFileNotFound,
		})
	}

	r, err := newReplay(formFile.Data, formFile.Header)
	if err != nil {
		return matchGen.NewCreateMatchBadRequest().WithPayload(&models.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	defer r.Close()

	matchID, err := c.match.CreateFromReplay(p.HTTPRequest.Context(), r)
	if err != nil {
		if errors.Is(err, domain.ErrMatchAlreadyExist) {
			return matchGen.NewCreateMatchConflict().WithPayload(&models.Error{
				Code:    domain.CodeMatchAlreadyExist,
				Message: err.Error(),
			})
		}

		return matchGen.NewCreateMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return matchGen.NewCreateMatchOK().WithPayload(&models.CreateMatchResponse{
		MatchID: strfmt.UUID(matchID.String()),
	})
}

func (c *Controller) DeleteMatch(p matchGen.DeleteMatchParams) matchGen.DeleteMatchResponder {
	matchID, err := uuid.Parse(p.MatchID.String())
	if err != nil {
		return matchGen.NewDeleteMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if err := c.match.DeleteByID(p.HTTPRequest.Context(), matchID); err != nil {
		if errors.Is(err, domain.ErrMatchNotFound) {
			return matchGen.NewDeleteMatchNotFound().WithPayload(&models.Error{
				Code:    domain.CodeMatchNotFound,
				Message: err.Error(),
			})
		}

		return matchGen.NewDeleteMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return matchGen.NewDeleteMatchNoContent()
}

func (c *Controller) GetMatch(p matchGen.GetMatchParams) matchGen.GetMatchResponder {
	matchID, err := uuid.Parse(p.MatchID.String())
	if err != nil {
		return matchGen.NewGetMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	match, err := c.match.GetByID(p.HTTPRequest.Context(), matchID)
	if err != nil {
		if errors.Is(err, domain.ErrMatchNotFound) {
			return matchGen.NewGetMatchNotFound().WithPayload(&models.Error{
				Code:    domain.CodeMatchNotFound,
				Message: err.Error(),
			})
		}

		return matchGen.NewGetMatchInternalServerError().WithPayload(&models.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	payload := models.Match{
		ID: strfmt.UUID(match.ID.String()),
		Map: models.Map{
			Name:    match.Map.Name,
			IconURL: match.Map.IconURL,
		},
		RoundsPlayed: match.RoundsPlayed,
		Team1: models.MatchTeam{
			ID:         match.Team1.ID,
			ClanName:   match.Team1.Name,
			FlagCode:   match.Team1.FlagCode,
			Score:      match.Team1.Score,
			Scoreboard: make([]models.MatchTeamScoreboard, len(match.Team1.ScoreBoard)),
		},
		Team2: models.MatchTeam{
			ID:         match.Team2.ID,
			ClanName:   match.Team2.Name,
			FlagCode:   match.Team2.FlagCode,
			Score:      match.Team2.Score,
			Scoreboard: make([]models.MatchTeamScoreboard, len(match.Team2.ScoreBoard)),
		},
		Duration:   int64(match.Duration),
		UploadedAt: strfmt.DateTime(match.UploadedAt),
	}

	for i, row := range match.Team1.ScoreBoard {
		payload.Team1.Scoreboard[i] = models.MatchTeamScoreboard{
			Assists:            row.Assists,
			DamagePerRound:     row.DamagePerRound,
			Deaths:             row.Deaths,
			HeadshotPercentage: row.HeadshotPercentage,
			KillDeathRatio:     row.KillDeathRatio,
			Kills:              row.Kills,
			Mvps:               row.MVPCount,
			PlayerName:         row.PlayerName,
			SteamID:            row.SteamID,
		}
	}

	for i, row := range match.Team2.ScoreBoard {
		payload.Team2.Scoreboard[i] = models.MatchTeamScoreboard{
			Assists:            row.Assists,
			DamagePerRound:     row.DamagePerRound,
			Deaths:             row.Deaths,
			HeadshotPercentage: row.HeadshotPercentage,
			KillDeathRatio:     row.KillDeathRatio,
			Kills:              row.Kills,
			Mvps:               row.MVPCount,
			PlayerName:         row.PlayerName,
			SteamID:            row.SteamID,
		}
	}

	return matchGen.NewGetMatchOK().WithPayload(&payload)
}
