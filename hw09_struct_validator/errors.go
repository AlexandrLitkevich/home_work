package hw09structvalidator

import "errors"

var (
	NotStruct     = errors.New("the data does not belong to the struct type")
	InvalidRule   = errors.New("invalid value of the validation rule")
	InvalidLength = errors.New("invalid string length")
)
