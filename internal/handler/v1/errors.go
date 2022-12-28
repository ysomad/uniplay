package v1

import (
	"errors"
	"net/http"

	"github.com/ssssargsian/uniplay/internal/pkg/apperror"
)

var (
	errBodySizeLimitExceeded      = errors.New("body size is too big")
	errInvalidReplayFileExtension = apperror.New(http.StatusBadRequest, "replay must have .dem file extension")
)
