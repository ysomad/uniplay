package domain

type Map struct {
	Name    string
	IconURL string
}

type MostPlayedMap struct {
	Map
	PlayedTimes int8
}
