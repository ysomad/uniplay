package domain

type InstType int8

func (t InstType) Int() int8 { return int8(t) }

const (
	InstTypeUniversity InstType = 1
	InstTypeCollege    InstType = 2
)

type Institution struct {
	ID        int32
	Type      InstType
	Name      string
	ShortName string
	City      string
	LogoURL   string
}

type InstitutionFilter struct {
	Type InstType
	City string
}

func NewInstitutionFilter(city *string, itype *int8) InstitutionFilter {
	f := InstitutionFilter{}

	if itype != nil {
		f.Type = InstType(*itype)
	}

	if city != nil {
		f.City = *city
	}

	return f
}
