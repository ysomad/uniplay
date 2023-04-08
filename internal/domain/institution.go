package domain

type Institution struct {
	ID        int32
	Name      string
	ShortName string
	LogoURL   string
}

type InstitutionFilter struct {
	ShortName string
}

type InstitutionPagination struct {
	PageSize int32
	Offset   int32
}

const (
	defaultInstitutionPageSize = 50
	maxInstitutionPageSize     = 250
)

func NewInstitutionPagination(psize, offset int32) InstitutionPagination {
	p := InstitutionPagination{
		PageSize: psize,
		Offset:   offset,
	}

	if psize > maxInstitutionPageSize || psize <= 0 {
		p.PageSize = defaultInstitutionPageSize
	}

	return p
}
