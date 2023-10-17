package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	if status := isValid(str); !status {
		return "", ErrInvalidString
	}

	var store strings.Builder
	prevChar := ""
	slash := false
	for _, char := range str {
		if char == '\\' && !slash {
			slash = true
			continue
		}

		switch {
		case unicode.IsNumber(char):
			repeat, err := strconv.Atoi(string(char))
			if err != nil {
				return "", ErrInvalidString
			}

			var repeatChar string
			if slash {
				repeatChar = string('\\') + prevChar
				slash = false
			} else {
				repeatChar = prevChar
			}

			newStr := strings.Repeat(repeatChar, repeat)
			store.WriteString(newStr)
			prevChar = ""
		case unicode.IsLetter(char):
			if prevChar != "" {
				store.WriteString(prevChar)
			}
			prevChar = string(char)
		}
	}
	if prevChar != "" {
		store.WriteString(prevChar)
	}

	return store.String(), nil
}

func isValid(str string) bool {
	isNumber := false
	for i, char := range str {
		if (i == 0) && (unicode.IsNumber(char)) {
			return false
		}
		if unicode.IsNumber(char) {
			if isNumber {
				return false
			}
			isNumber = true
		} else {
			isNumber = false
		}
	}
	return true
}
