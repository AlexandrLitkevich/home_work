package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runes := []rune(str)

	if status := isValid(runes); !status {
		return "", ErrInvalidString
	}

	var store strings.Builder
	prevChar := ""
	isSymbol := false
	for _, char := range runes {
		if char == '\\' && !isSymbol {
			isSymbol = true
			continue
		}

		switch {
		case string(char) >= "0" && string(char) <= "9":
			repeat, err := strconv.Atoi(string(char))
			if err != nil {
				return "", ErrInvalidString
			}

			var repeatChar string
			if isSymbol {
				repeatChar = string('\\') + prevChar
				isSymbol = false
			} else {
				repeatChar = prevChar
			}

			newStr := strings.Repeat(repeatChar, repeat)
			store.WriteString(newStr)
			prevChar = ""
		default:
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

func isValid(runes []rune) bool {
	isNumber := false

	for i, char := range runes {
		if (i == 0) && (string(char) >= "0" && string(char) <= "9") {
			return false
		}
		if string(char) >= "0" && string(char) <= "9" {
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
