package domain

type Player struct {
	SteamID     SteamID
	TeamID      int32
	DisplayName string
	FirstName   string
	LastName    string
	AvatarURL   string
}
