package hw09structvalidator

import (
	"reflect"
	"strings"
)

const (
	validateTag = "validate"
	separatorOR = "|"
	separator   = ":"
)

type FieldData struct {
	Value reflect.Value
	Tag   string
	Name  string
}

type FieldDataSlice struct {
	Info  reflect.StructField
	Value reflect.Value
	Tag   string
}

type Rules struct {
	name  string
	value string
}

func (f *FieldData) GetRules() []Rules {
	rules := strings.Split(f.Tag, separatorOR)
	rulesInfo := make([]Rules, 0, len(rules))

	for _, item := range rules {
		rule := strings.Split(item, separator)
		rulesInfo = append(rulesInfo, Rules{rule[0], rule[1]})
	}
	return rulesInfo
}
