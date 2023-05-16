package domain

const (
	MostPlayedMapsLimit  = 5
	MostSuccessMapsLimit = 5
)

type Map struct {
	Name    string
	IconURL string
}

type MostPlayedMap struct {
	Map
	PlayedTimes int8
}

type MostSuccessMap struct {
	Map
	WinRate float64
}
