package models

import "errors"

var (
	ErrIncorrectJSON  = errors.New("INCORRECT_JSON_ERROR")
	ErrInternalServer = errors.New("INTERNAL_SERVER_ERROR")
)
