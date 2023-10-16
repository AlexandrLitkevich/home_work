package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	notInt, err := regexp.MatchString("[0-9]+", str)
	if !notInt {
		return str, nil
	}
	if err != nil {
		return "", ErrInvalidString
	}
	runes := []rune(str)
	if unicode.IsNumber(runes[0]) {
		return "", ErrInvalidString
	}
	var res string
	flag := false
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\\' {
			res += string(runes[i])
			flag = true
			continue
		}
		if unicode.IsLetter(runes[i]) {
			res += string(runes[i])
		} else {
			repeatChar := ""
			for j := i; j < len(str); j++ {
				if unicode.IsNumber(runes[j]) && unicode.IsNumber(runes[j-1]) {
					return "", ErrInvalidString
				} else if unicode.IsNumber(runes[j]) {
					repeatChar += string(runes[j])
				} else {
					break
				}
			}
			if repeatChar == "0" {
				res = res[:len(res)-1]
				continue
			}
			count, _ := strconv.Atoi(repeatChar)
			if flag {
				res += strings.Repeat(string('\\')+string(str[i-1]), count-1)
			} else {
				res += strings.Repeat(string(str[i-1]), count-1)
			}
		}
	}
	return res, nil
}
