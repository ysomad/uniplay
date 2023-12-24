package httpapi

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeStatus(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	s := http.StatusText(code)
	_ = json.NewEncoder(w).Encode(httpError{ //nolint:errcheck // not needed to check for error here
		Code:    s,
		Message: s,
	})
}

func writerError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(httpError{ //nolint:errcheck // not needed to check for error here
		Code:    http.StatusText(code),
		Message: err.Error(),
	})
}
