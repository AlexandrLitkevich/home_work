package hw09structvalidator

import (
	"fmt"
	"reflect"
)

func validateSlice(field FieldDataSlice, validateErrors *ValidationErrors) error {
	for i := 0; i < field.Value.Len(); i++ {
		item := FieldData{
			Value: field.Value.Index(i),
			Name:  fmt.Sprintf("%v index:%v", field.Info.Name, i),
			Tag:   field.Tag,
		}
		//nolint:exhaustive
		switch field.Value.Index(i).Kind() {
		case reflect.String:
			err := ValidateString(item, validateErrors)
			if err != nil {
				return err
			}
		case reflect.Int:
			err := validateInt(item, validateErrors)
			if err != nil {
				return err
			}
		default:
			return ErrInvalidSliceValue
		}
	}
	return nil
}
