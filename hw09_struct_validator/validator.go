package hw09structvalidator

import (
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
	* type assertion к структуре
 	* прочитать с помощью reflect структуру


*/

func Validate(v interface{}) error {
	vType := reflect.TypeOf(v)
	fmt.Println("this vType", vType)

	var errors []


	ref := reflect.ValueOf(v)
	fmt.Println("this struct", ref.NumField())
	return nil
}
