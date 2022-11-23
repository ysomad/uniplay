package apperror

var _ error = &Err{}

type Err struct {
	Code uint16
	Msg  string
}

func New(code uint16, msg string) Err {
	return Err{
		Code: code,
		Msg:  msg,
	}
}

func (e Err) Error() string { return e.Msg }
