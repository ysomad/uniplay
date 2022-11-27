package domain

type WeaponClassID uint8

const (
	ClassPistol       WeaponClassID = 1
	ClassSMG          WeaponClassID = 2
	ClassShotgun      WeaponClassID = 3
	ClassMachineGun   WeaponClassID = 4
	ClassAssaultRifle WeaponClassID = 5
	ClassSniperRifle  WeaponClassID = 6
	ClassEquipment    WeaponClassID = 7
	ClassOther        WeaponClassID = 8
	ClassGrenade      WeaponClassID = 9
)

type WeaponClass struct {
	ID    WeaponClassID `json:"id"`
	Class string        `json:"class"`
}
