package domain

type TeamListItem struct {
	ID       int32
	ClanName string
	FlagCode string

	InstID        int32
	InstShortName string
	InstCity      string
	InstType      InstType
	InstLogoURL   string
}
