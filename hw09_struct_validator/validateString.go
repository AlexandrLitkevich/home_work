package hw09structvalidator

import (
	"regexp"
	"strconv"
)

func ValidateString(field FieldData, validateErrors *ValidationErrors) error {
	rules := field.GetRules()
	value := field.Value.String()

	for _, rule := range rules {
		if rule.name == rulesLen {
			valRule, err := strconv.Atoi(rule.value)
			if err != nil {
				return ErrInvalidRule
			}
			if len(value) != valRule {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.Name, Err: ErrInvalidLength})
			}
			continue
		}
		if rule.name == rulesIn {
			isMatch, err := regexp.MatchString(value, rule.value)
			if err != nil {
				return err
			}

			if !isMatch {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.Name, Err: ErrInvalidValue})
			}
		}
		if rule.name == rulesRegexp {
			re := regexp.MustCompile(rule.value)
			result := re.MatchString(value)
			if !result {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.Name, Err: ErrInvalidRegexp})
			}
		}
	}
	return nil
}
