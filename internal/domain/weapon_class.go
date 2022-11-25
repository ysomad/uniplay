package domain

type WeaponClass uint8

const (
	WpClassPistols   WeaponClass = 1
	WpClassSMG       WeaponClass = 2
	WpClassHeavy     WeaponClass = 3
	WpClassRifle     WeaponClass = 4
	WpClassEquipment WeaponClass = 5
	WpClassGrenade   WeaponClass = 6
)

var (
	stringToWeaponClass = map[string]WeaponClass{
		"pistol":    WpClassPistols,
		"smg":       WpClassSMG,
		"heavy":     WpClassHeavy,
		"rifle":     WpClassRifle,
		"equipment": WpClassEquipment,
		"grenade":   WpClassGrenade,
	}

	weaponClassToString = map[WeaponClass]string{
		WpClassPistols:   "pistol",
		WpClassSMG:       "smg",
		WpClassHeavy:     "heavy",
		WpClassRifle:     "rifle",
		WpClassEquipment: "equipment",
		WpClassGrenade:   "grenade",
	}
)

func NewWeaponClass(class string) WeaponClass {
	if c, ok := stringToWeaponClass[class]; ok {
		return c
	}
	return 0
}

func (c WeaponClass) Valid() bool {
	switch c {
	case WpClassPistols, WpClassSMG, WpClassHeavy, WpClassRifle, WpClassEquipment, WpClassGrenade:
		return true
	}
	return false
}

func (c WeaponClass) String() string {
	if s, ok := weaponClassToString[c]; ok {
		return s
	}
	return "unknown"
}
