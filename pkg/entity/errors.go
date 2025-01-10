package entity

import "errors"

// ErrNotFound is returned when a requested entity is not found.
var ErrNotFound = errors.New("url not found")
