package hw09structvalidator

import (
	"fmt"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	errBuffer := strings.Builder{}
	errBuffer.WriteString("Invalid fields: ")

	for _, err := range v {
		errBuffer.WriteString(fmt.Sprintf("%s: %s \n", err.Field, err.Err))
	}
	return errBuffer.String()
}

func Validate(v interface{}) error {
	vType := reflect.TypeOf(v)

	if vType.Kind() != reflect.Struct {
		return NotStruct
	}

	itemsStructValue := reflect.ValueOf(v)
	validateErrors := ValidationErrors{}

	for i := 0; i < itemsStructValue.NumField(); i++ {
		itemStructField := vType.Field(i)

		tag, ok := itemStructField.Tag.Lookup(validateTag)
		if !ok {
			continue
		}

		field := FieldData{
			value: itemsStructValue.Field(i),
			info:  itemStructField,
			tag:   tag,
		}

		switch itemStructField.Type.Kind() {
		case reflect.String:
			err := ValidateString(field, &validateErrors)
			if err != nil {
				return err
			}
		case reflect.Int:
			validateInt(field)
		default:
			return nil
		}
	}

	if len(validateErrors) > 0 {
		return validateErrors
	}

	return nil
}
