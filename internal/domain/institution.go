package domain

type Institution struct {
	ID        int16
	Name      string
	ShortName string
	LogoURL   string
}

type InstitutionFilter struct {
	ShortName string
}
