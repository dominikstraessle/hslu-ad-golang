package formal_language

import (
	"fmt"
	"testing"
)

func Test_isWordLanguage(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "valid",
			word: "0",
			want: true,
		},

		{
			name: "valid",
			word: "0",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWordLanguage(tt.word); got != tt.want {
				t.Errorf("isWordLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be invalid", tt), func(t *testing.T) {
			if isWordLanguage(tt) != false {
				t.Error("isWordLanguage() = true, want false")
			}
		})
	}
}
