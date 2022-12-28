package apperror

import (
	"encoding/json"
	"net/http"
)

// Write writes error to response body with code equals to status.
func Write(w http.ResponseWriter, status int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	code := status
	if apperr, ok := err.(Err); ok {
		code = apperr.Code
	}

	return json.NewEncoder(w).Encode(Err{
		Code:    code,
		Message: err.Error(),
	})
}
