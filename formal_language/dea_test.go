package formal_language

import (
	"fmt"
	"testing"
)

func Test_isWordLanguageValidCases(t *testing.T) {
	tests := []string{
		"0",
		"010",
		"01110",
		"0111010",
		"0101110",
		"010101010",
		"0111010111010110111101011010",
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be valid", tt), func(t *testing.T) {
			if isWordLanguage(tt) != true {
				t.Error("isWordLanguage() = false, want true")
			}
		})
	}

}

func Test_isWordLanguageInvalidCases(t *testing.T) {
	tests := []string{
		"10",
		"0101",
		"00",
		"x",
		"",
		"011x110",
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be invalid", tt), func(t *testing.T) {
			if isWordLanguage(tt) != false {
				t.Error("isWordLanguage() = true, want false")
			}
		})
	}
}

func Test_isWordLanguageValidCasesRecursive(t *testing.T) {
	tests := []string{
		"0",
		"010",
		"01110",
		"0111010",
		"0101110",
		"010101010",
		"0111010111010110111101011010",
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be valid", tt), func(t *testing.T) {
			if isWordLanguageRecursive(tt) != true {
				t.Error("isWordLanguage() = false, want true")
			}
		})
	}

}

func Test_isWordLanguageInvalidCasesRecursive(t *testing.T) {
	tests := []string{
		"10",
		"0101",
		"00",
		"x",
		"",
		"011x110",
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be invalid", tt), func(t *testing.T) {
			if isWordLanguageRecursive(tt) != false {
				t.Error("isWordLanguage() = true, want false")
			}
		})
	}
}

func Test_isWordLanguageValidCasesRegex(t *testing.T) {
	tests := []string{
		"0",
		"010",
		"01110",
		"0111010",
		"0101110",
		"010101010",
		"0111010111010110111101011010",
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be valid", tt), func(t *testing.T) {
			if isWordLanguageRegex(tt) != true {
				t.Error("isWordLanguage() = false, want true")
			}
		})
	}

}

func Test_isWordLanguageInvalidCasesRegex(t *testing.T) {
	tests := []string{
		"10",
		"0101",
		"00",
		"x",
		"",
		"011x110",
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be invalid", tt), func(t *testing.T) {
			if isWordLanguageRegex(tt) != false {
				t.Error("isWordLanguage() = true, want false")
			}
		})
	}
}
