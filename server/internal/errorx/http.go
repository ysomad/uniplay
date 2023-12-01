package errorx

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func WriteMessage(w http.ResponseWriter, code int, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(httpError{
		Code: http.StatusText(code),
		Msg:  msg,
	})
}

func WriteStatus(w http.ResponseWriter, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	s := http.StatusText(code)
	return json.NewEncoder(w).Encode(httpError{
		Code: s,
		Msg:  s,
	})
}

func WriteError(w http.ResponseWriter, code int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(httpError{
		Code: http.StatusText(code),
		Msg:  err.Error(),
	})
}
