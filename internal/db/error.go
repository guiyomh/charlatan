package db

import (
	"github.com/pkg/errors"
)

// ErrCouldNotConvertToTime is returns when a string is not a reconizable time format
var ErrCouldNotConvertToTime = errors.New("Could not convert string to time")
