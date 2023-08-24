package domain

import (
	"strconv"
)

type SteamID uint64

func NewSteamID(s string) (SteamID, error) {
	steamID64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return SteamID(steamID64), nil
}

func (s SteamID) String() string {
	return strconv.FormatUint(uint64(s), 10)
}
