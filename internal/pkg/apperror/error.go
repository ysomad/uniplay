package apperror

var _ error = &Err{}

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, msg string) Err {
	return Err{
		Code:    code,
		Message: msg,
	}
}

func (e Err) Error() string { return e.Message }
