package rest

import (
	"errors"
)

var (
	ErrInvalidPage  = errors.New("ERR_INVALID_PAGE")
	ErrInvalidLimit = errors.New("ERR_INVALID_LIMIT")
)
