package hw09structvalidator

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	rulesLen    = "len"
	rulesIn     = "in"
	rulesRegexp = "regexp"
	rulesMin    = "min"
	rulesMax    = "max"
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
		return ErrNotStruct
	}

	itemsStructValue := reflect.ValueOf(v)
	validateErrors := ValidationErrors{}

	for i := 0; i < itemsStructValue.NumField(); i++ {
		itemStructField := vType.Field(i)

		tag, ok := itemStructField.Tag.Lookup(validateTag)
		if !ok {
			continue
		}

		if itemStructField.Type.Kind() == reflect.Slice {
			sliceLen := itemsStructValue.Field(i).Len()
			value := itemsStructValue.Field(i).Slice(0, sliceLen)

			err := validateSlice(FieldDataSlice{Value: value, Tag: tag, Info: itemStructField}, &validateErrors)
			if err != nil {
				return err
			}
			continue
		}

		field := FieldData{
			Tag:   tag,
			Name:  itemStructField.Name,
			Value: itemsStructValue.Field(i),
		}
		//nolint:exhaustive
		switch itemStructField.Type.Kind() {
		case reflect.String:
			err := ValidateString(field, &validateErrors)
			if err != nil {
				return err
			}
		case reflect.Int:
			err := validateInt(field, &validateErrors)
			if err != nil {
				return err
			}
		default:
			return nil
		}
	}

	if len(validateErrors) > 0 {
		return validateErrors
	}

	return nil
}
