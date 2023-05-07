package team

import (
	"context"

	"github.com/ysomad/uniplay/internal/domain"
	"github.com/ysomad/uniplay/internal/pkg/paging"
)

type repository interface {
	GetAll(context.Context, listParams) (paging.InfList[domain.TeamListItem], error)
	GetPlayers(ctx context.Context, teamID int32) ([]domain.TeamPlayer, error)
	Update(ctx context.Context, teamID int32, p updateParams) (domain.Team, error)
	SetCaptain(ctx context.Context, teamID int32, steamID domain.SteamID) error
}

type service struct {
	team repository
}

func NewService(r repository) *service {
	return &service{
		team: r,
	}
}

func (s *service) GetList(ctx context.Context, p listParams) (paging.InfList[domain.TeamListItem], error) {
	return s.team.GetAll(ctx, p)
}

func (s *service) GetPlayers(ctx context.Context, teamID int32) ([]domain.TeamPlayer, error) {
	return s.team.GetPlayers(ctx, teamID)
}

func (s *service) Update(ctx context.Context, teamID int32, p updateParams) (domain.Team, error) {
	return s.team.Update(ctx, teamID, p)
}

func (s *service) SetCaptain(ctx context.Context, teamID int32, steamID domain.SteamID) error {
	return s.team.SetCaptain(ctx, teamID, steamID)
}
