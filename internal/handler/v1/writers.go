package v1

import (
	"encoding/json"
	"net/http"

	v1 "github.com/ssssargsian/uniplay/internal/gen/oapi/v1"
	"github.com/ssssargsian/uniplay/internal/pkg/apperror"
)

// writeError writes error to response body with code equals to status.
func writeError(w http.ResponseWriter, status int, err error) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	var code uint16 = http.StatusInternalServerError
	if apperr, ok := err.(apperror.Err); ok {
		code = apperr.Code
	}

	if err := json.NewEncoder(w).Encode(v1.Error{
		Code:    code,
		Message: err.Error(),
	}); err != nil {
		return err
	}

	return nil
}
