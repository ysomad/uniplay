package domain

type WeaponClassID uint8

const (
	WpClassPistols   WeaponClassID = 1
	WpClassSMG       WeaponClassID = 2
	WpClassHeavy     WeaponClassID = 3
	WpClassRifle     WeaponClassID = 4
	WpClassEquipment WeaponClassID = 5
	WpClassGrenade   WeaponClassID = 6
)

var (
	strToWeaponClassID = map[string]WeaponClassID{
		"pistol":    WpClassPistols,
		"smg":       WpClassSMG,
		"heavy":     WpClassHeavy,
		"rifle":     WpClassRifle,
		"equipment": WpClassEquipment,
		"grenade":   WpClassGrenade,
	}

	weaponClassIDToStr = map[WeaponClassID]string{
		WpClassPistols:   "pistol",
		WpClassSMG:       "smg",
		WpClassHeavy:     "heavy",
		WpClassRifle:     "rifle",
		WpClassEquipment: "equipment",
		WpClassGrenade:   "grenade",
	}
)

func NewWeaponClassID(class string) WeaponClassID {
	if c, ok := strToWeaponClassID[class]; ok {
		return c
	}
	return 0
}

func (c WeaponClassID) Valid() bool {
	switch c {
	case WpClassPistols, WpClassSMG, WpClassHeavy, WpClassRifle, WpClassEquipment, WpClassGrenade:
		return true
	}
	return false
}

func (c WeaponClassID) String() string {
	if s, ok := weaponClassIDToStr[c]; ok {
		return s
	}
	return "unknown"
}

type WeaponClass struct {
	ID   WeaponClassID
	Name string
}
