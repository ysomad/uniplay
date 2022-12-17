package v1

import "errors"

var (
	errBodySizeLimitExceeded  = errors.New("body size is too big")
	errInvalidBodyContentType = errors.New("invalid body content type")
)
