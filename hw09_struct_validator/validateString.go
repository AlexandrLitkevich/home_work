package hw09structvalidator

import (
	"regexp"
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
			valRule, err := strconv.Atoi(rule.value)
			if err != nil {
				return InvalidRule
			}
			if len(strFieldValue) != valRule {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.GetName(), Err: InvalidLength})
			}
			continue
		}
		if rule.name == rulesIn {
			isMatch, err := regexp.MatchString(strFieldValue, rule.value)
			if err != nil {
				return err
			}

			if !isMatch {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.GetName(), Err: InvalidValue})
			}
		}
		if rule.name == rulesRegexp {
			re := regexp.MustCompile(rule.value)
			result := re.MatchString(strFieldValue)
			if !result {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.GetName(), Err: InvalidRegexp})
			}
		}
	}
	return nil
}
