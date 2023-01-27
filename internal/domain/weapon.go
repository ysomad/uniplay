package domain

type Weapon struct {
	WeaponID int32  `json:"weapon_id"`
	Weapon   string `json:"weapon"`
	ClassID  int32  `json:"class_id"`
	Class    string `json:"class"`
}
