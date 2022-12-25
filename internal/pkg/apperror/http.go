package apperror

import (
	"encoding/json"
	"net/http"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
)

// Write writes error to response body with code equals to status.
func Write(w http.ResponseWriter, status int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	code := status
	if apperr, ok := err.(Err); ok {
		code = apperr.Code
	}

	return json.NewEncoder(w).Encode(v1.Error{
		Code:    code,
		Message: err.Error(),
	})
}
