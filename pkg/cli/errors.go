package cli

import "errors"

var (
	ErrInvalidName         = errors.New("invalid name")
	ErrObjectAlreadyExists = errors.New("object already exists")
)
