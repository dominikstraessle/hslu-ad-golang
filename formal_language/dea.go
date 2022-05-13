package formal_language

import (
	"bytes"
	"regexp"
)

const (
	ValidZero = '0'
	ValidOne  = '1'
)

var matcher = regexp.MustCompile("^0((1)+0)*$")

func isWordLanguageRegex(word string) bool {
	return matcher.MatchString(word)
}

func isWordLanguage(word string) bool {
	length := len(word) - 1
	if length < 0 {
		return false
	}
	previous := ValidOne
	for i, c := range word {
		if i == 0 && c != ValidZero {
			return false
		}
		if i >= length && c != ValidZero {
			return false
		}

		if c != ValidZero && c != ValidOne {
			return false
		}

		if previous == ValidZero && c == ValidZero {
			return false
		}

		previous = c
	}

	return true
}

func isWordLanguageRecursive(word string) bool {
	runes := bytes.Runes([]byte(word))
	if len(runes) < 1 {
		return false
	}
	if runes[0] != ValidZero {
		return false
	}
	if runes[len(runes)-1] != ValidZero {
		return false
	}

	return recursive(runes, ValidOne, 0)
}

func recursive(runes []rune, previous rune, index int) bool {
	c := runes[index]
	if c != ValidZero && c != ValidOne {
		return false
	}

	if previous == ValidZero && c == ValidZero {
		return false
	}

	if index < len(runes)-1 {
		return recursive(runes, c, index+1)
	}
	return true
}
