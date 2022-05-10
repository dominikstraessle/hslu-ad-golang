package formal_language

import (
	"fmt"
	"strings"
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

func BenchmarkIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isWordLanguage("0111010111010110111101011010")
	}
}
func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isWordLanguageRecursive("0111010111010110111101011010")
	}
}
func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isWordLanguageRegex("0111010111010110111101011010")
	}
}

var word = generate()

func BenchmarkIterative_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isWordLanguage(word)
	}
}
func BenchmarkRecursive_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isWordLanguageRecursive(word)
	}
}
func BenchmarkRegex_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isWordLanguageRegex(word)
	}
}

func generate() string {
	b := strings.Builder{}
	b.WriteString("0")
	for i := 0; i < 10000; i++ {
		b.WriteString("11101011101011011110101101")
	}
	b.WriteString("0")
	return b.String()
}
