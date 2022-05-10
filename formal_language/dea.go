package formal_language

import "bytes"

func isWordLanguage(word string) bool {
	validZero, validOne := rune('0'), rune('1')

	if len(word) < 1 {
		return false
	}

	previous := validOne
	for _, c := range word {
		if c != validZero && c != validOne {
			return false
		}

		if previous == validZero && c == validZero {
			return false
		}

		previous = c
	}

	runes := bytes.Runes([]byte(word))
	if runes[0] != validZero {
		return false
	}
	if runes[len(runes)-1] != validZero {
		return false
	}

	return true
}
