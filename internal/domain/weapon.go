package domain

type Weapon struct {
	WeaponID int16  `json:"weapon_id"`
	Weapon   string `json:"weapon"`
	ClassID  int8   `json:"class_id"`
	Class    string `json:"class"`
}
