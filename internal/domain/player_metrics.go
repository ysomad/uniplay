package domain

// PlayerMetrics is a set of total statistics of a player.
type PlayerMetrics struct {
	Deaths            uint32
	Kills             uint32
	HSKills           uint32
	BlindKills        uint32
	WallbangKills     uint32
	NoScopeKills      uint32
	ThroughSmokeKills uint32
	Assists           uint32
	FlashbangAssists  uint32
	DamageTaken       uint32
	DamageDealt       uint32
	BombsPlanted      uint32
	BombsDefused      uint32
	MVPCount          uint32
	BlindedPlayers    uint32
	BlindedTimes      uint32
}
