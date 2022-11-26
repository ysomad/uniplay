package domain

type Weapon struct {
	ID        uint16
	Name      string
	ClassID   WeaponClassID
	ClassName string
}
