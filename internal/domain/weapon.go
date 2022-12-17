package domain

type Weapon struct {
	WeaponID uint16 `json:"weapon_id"`
	Weapon   string `json:"weapon"`
	ClassID  uint8  `json:"class_id"`
	Class    string `json:"class"`
}
