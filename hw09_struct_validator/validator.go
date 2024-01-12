package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	panic("implement me")

}

/*
STEP_BY_STEP
*
	* type assertion к структуре
 	* прочитать с помощью reflect структуру
	* получаем поле читаем struct tag
	* получаем  filed value
	* провалидировать string and int
*/

const (
	len    = "len:"
	regexp = "regexp:"
)

var (
	NotStruct = errors.New("the data does not belong to the struct type ")
)

func Validate(v interface{}) error {

	vValue := reflect.TypeOf(v)
	if vValue.Kind() != reflect.Struct {
		return NotStruct
	}

	//var errors []
	for i := 0; i < vValue.NumField(); i++ {
		field := vValue.Field(i)
		tag, ok := field.Tag.Lookup("validate")
		if !ok {
			continue
		}

		fmt.Println("this value tag validate", tag)

	}
	return nil
}
