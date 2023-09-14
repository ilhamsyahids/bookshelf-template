package middleware

import (
	"errors"
)

var (
	ErrInvalidAPIKey = errors.New("ERR_INVALID_API_KEY")
	ErrMissingAPIKey = errors.New("ERR_MISSING_API_KEY")
)
