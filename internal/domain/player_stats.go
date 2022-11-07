package domain

// SteamID represents steam uint64 id.
type SteamID uint64

// stats is a map of events, key is event, value is amount of entries of the event.
type stats map[metric]uint16

// playerStats is a map of player event entries.
type playerStats struct {
	stats map[SteamID]stats
}

func NewPlayerStats() *playerStats {
	return &playerStats{
		stats: make(map[SteamID]stats),
	}
}

// Get returns stats of specific player with steamID.
func (s *playerStats) Get(steamID uint64) (stats, bool) {
	v, ok := s.stats[SteamID(steamID)]
	return v, ok
}

// Add n to amount of player metric entries in the stats map of specific player with steamID.
func (s *playerStats) Add(steamID uint64, m metric, n uint16) {
	_, ok := s.stats[SteamID(steamID)]
	if !ok {
		s.stats[SteamID(steamID)] = make(stats)
	}

	s.stats[SteamID(steamID)][m] += n
}
