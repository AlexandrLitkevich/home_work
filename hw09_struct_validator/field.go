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
	value reflect.Value
	info  reflect.StructField
	tag   string
}

type Rules struct {
	name  string
	value string
}

func (f *FieldData) GetName() string {
	return f.info.Name
}

func (f *FieldData) GetRules() []Rules {
	rules := strings.Split(f.tag, separatorOR)
	var rulesInfo []Rules

	for _, item := range rules {
		rule := strings.Split(item, separator)
		rulesInfo = append(rulesInfo, Rules{rule[0], rule[1]})
	}
	return rulesInfo
}
