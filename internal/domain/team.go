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

type TeamPlayer struct {
	SteamID     SteamID
	DisplayName string
	FirstName   string
	LastName    string
	AvatarURL   string
	IsCaptain   bool
}
