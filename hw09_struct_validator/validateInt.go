package hw09structvalidator

import (
	"regexp"
	"strconv"
)

func validateInt(field FieldData, validateErrors *ValidationErrors) error {
	rules := field.GetRules()
	value := field.Value.Int()

	for _, rule := range rules {
		if rule.name == rulesMin {
			ruleValue, err := strconv.Atoi(rule.value)
			if err != nil {
				return err
			}

			if value < int64(ruleValue) {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.Name, Err: ErrInvalidMinValue})
			}
		}

		if rule.name == rulesMax {
			intValue, err := strconv.Atoi(rule.value)
			if err != nil {
				return err
			}

			if value > int64(intValue) {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.Name, Err: ErrInvalidMaxValue})
			}
		}
		if rule.name == rulesIn {
			strFieldValue := strconv.Itoa(int(value))

			isMatch, err := regexp.MatchString(strFieldValue, rule.value)
			if err != nil {
				return err
			}

			if !isMatch {
				*validateErrors = append(*validateErrors, ValidationError{Field: field.Name, Err: ErrInvalidValue})
			}
		}
	}
	return nil
}
