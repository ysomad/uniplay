package apperror

var _ error = &Err{}

type Err struct {
	Code int
	Msg  string
}

func New(code int, msg string) Err {
	return Err{
		Code: code,
		Msg:  msg,
	}
}

func (e Err) Error() string { return e.Msg }
