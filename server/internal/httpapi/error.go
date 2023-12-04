package httpapi

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeMessage(w http.ResponseWriter, code int, msg string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(httpError{
		Code:    http.StatusText(code),
		Message: msg,
	})
}

func writeStatus(w http.ResponseWriter, code int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	s := http.StatusText(code)
	return json.NewEncoder(w).Encode(httpError{
		Code:    s,
		Message: s,
	})

}

func writerError(w http.ResponseWriter, code int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(httpError{
		Code:    http.StatusText(code),
		Message: err.Error(),
	})
}
