package eerr

import "errors"

var (
	ErrEmployeeNotFound      = errors.New("Employee not found")
	ErrEmployeeAlreadyExists = errors.New("Employee already exists")
)
