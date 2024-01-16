package hw09structvalidator

import (
	"fmt"
	"strconv"
)

const (
	rulesLen    = "len"
	rulesIn     = "in"
	rulesRegexp = "regexp"
)

func ValidateString(field FieldData, validateErrors *ValidationErrors) error {
	rules := field.GetRules()
	strFieldValue := field.value.String()

	for _, rule := range rules {
		if rule.name == rulesLen {
			fmt.Println("strFieldValue", strFieldValue)
			valRule, err := strconv.Atoi(rule.value)
			if err != nil {
				return InvalidRule
			}
			if len(strFieldValue) != valRule {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.GetName(), Err: InvalidLength})
			}
		}
	}
	return nil
}
