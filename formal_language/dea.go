package formal_language

import (
	"bytes"
	"fmt"
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

type state struct {
	current string
}

const (
	Z0 = "Z0"
	Z1 = "Z1"
	Z2 = "Z2"
	Z3 = "Z3"
	Z4 = "Z4"
)

func (s *state) Input(i rune) bool {
	switch s.current {
	case Z0:
		switch i {
		case ValidZero:
			s.current = Z1
			return true
		default:
			return false
		}
	case Z1:
		switch i {
		case ValidOne:
			s.current = Z2
			return true
		default:
			return false
		}
	case Z2:
		switch i {
		case ValidZero:
			s.current = Z4
			return true
		case ValidOne:
			s.current = Z3
			return true
		default:
			return false
		}
	case Z3:
		switch i {
		case ValidOne:
			s.current = Z2
			return true
		default:
			return false
		}
	case Z4:
		switch i {
		case ValidOne:
			s.current = Z2
			return true
		default:
			return false
		}
	default:
		return false
	}
}

func isWordLanguageDEA(word string) bool {
	s := &state{current: Z0}
	for i, r := range word {
		got := s.Input(r)
		if got == false {
			fmt.Println(i)
			return false
		}
	}

	return s.current == Z4 || s.current == Z1
}

func isWordLanguage(word string) bool {
	length := len(word) - 1
	if length < 0 {
		return false
	}
	previous := ValidOne
	numberOfOnes := 0
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

		if c == ValidOne {
			numberOfOnes++
		} else {
			if previous == ValidOne {
				if numberOfOnes%2 == 0 && numberOfOnes != 0 {
					return false
				}
				numberOfOnes = 0
			}
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

	return recursive(runes, ValidOne, 0, 0)
}

func recursive(runes []rune, previous rune, index int, numberOfOnes int) bool {
	c := runes[index]
	if c != ValidZero && c != ValidOne {
		return false
	}

	if previous == ValidZero && c == ValidZero {
		return false
	}

	if index < len(runes)-1 {
		if c == ValidOne {
			numberOfOnes++
		} else {
			if previous == ValidOne {
				if numberOfOnes%2 == 0 && numberOfOnes != 0 {
					return false
				}
				numberOfOnes = 0
			}
		}
		return recursive(runes, c, index+1, numberOfOnes)
	}
	return true
}
