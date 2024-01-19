package hw09structvalidator

import "errors"

var (
	ErrNotStruct         = errors.New("the data does not belong to the struct type")
	ErrInvalidRule       = errors.New("invalid value of the validation rule")
	ErrInvalidLength     = errors.New("invalid string length")
	ErrInvalidValue      = errors.New("invalid value")
	ErrInvalidRegexp     = errors.New("invalid regexp")
	ErrInvalidMinValue   = errors.New("the value is less than acceptable")
	ErrInvalidMaxValue   = errors.New("the value is greater than the allowed value")
	ErrInvalidSliceValue = errors.New("the elements are not supported")
)
