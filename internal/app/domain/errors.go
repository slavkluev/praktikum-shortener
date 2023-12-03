package domain

import "errors"

// ErrNotFound will throw if the requested item is not exists
// ErrConflict will throw if the current action already exists
var (
	ErrNotFound = errors.New("your requested Record is not found")
	ErrConflict = errors.New("your Record already exist")
)
