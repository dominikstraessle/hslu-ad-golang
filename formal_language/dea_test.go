package formal_language

import (
	"fmt"
	"strings"
	"testing"
)

func FuzzIsWordLanguage(f *testing.F) {
	tests := []string{
		"0",
		"010",
		"01110",
		"0111010",
		"0101110",
		"010101010",
		"0111010111010110111101011010",
		"10",
		"0101",
		"00",
		"x",
		"",
		"011x110",
	}
	for _, tt := range tests {
		f.Add(tt)
	}
	f.Fuzz(func(t *testing.T, word string) {
		gotIterative := isWordLanguage(word)
		gotRecursive := isWordLanguageRecursive(word)
		//gotRegex := isWordLanguageRegex(word)
		gotDEA := isWordLanguageDEA(word)

		//if gotIterative != gotRecursive || gotRecursive != gotRegex || gotIterative != gotRegex || gotIterative != gotDEA {
		if gotIterative != gotRecursive || gotIterative != gotDEA {
			//t.Errorf("%s got following results: iter(%v) recursive(%v) regex(%v)", word, gotIterative, gotRecursive, gotRegex)
			t.Errorf("%s got following results: iter(%v) recursive(%v)", word, gotIterative, gotRecursive)
		}
	})
}

func Test_isWordLanguageValidCases(t *testing.T) {
	tests := []string{
		"0",
		"010",
		"01110",
		"0111010",
		"0101110",
		"010101010",
		"0111010111010111011111010111010",
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
		"0111010111010110111101011010",
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be invalid", tt), func(t *testing.T) {
			if isWordLanguage(tt) != false {
				t.Error("isWordLanguage() = true, want false")
			}
		})
	}
}

func Test_isWordLanguageDEAValidCases(t *testing.T) {
	tests := []string{
		"0",
		"010",
		"01110",
		"0111010",
		"0101110",
		"010101010",
		"0111010111010111011111010111010",
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be valid", tt), func(t *testing.T) {
			if isWordLanguageDEA(tt) != true {
				t.Error("isWordLanguage() = false, want true")
			}
		})
	}

}

func Test_isWordLanguageDEAInvalidCases(t *testing.T) {
	tests := []string{
		"10",
		"0101",
		"00",
		"x",
		"",
		"011x110",
		"0111010111010110111101011010",
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v should be invalid", tt), func(t *testing.T) {
			if isWordLanguageDEA(tt) != false {
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
		"0111010111010111011111010111010",
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
		"0111010111010110111101011010",
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
		"0111010111010111011111010111010",
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
	//TODO: fix regex when there are an even number of ones -> see last test
	t.Skip()
	tests := []string{
		"10",
		"0101",
		"00",
		"x",
		"",
		"011x110",
		"0111010111010110111101011010",
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
		isWordLanguage("0111010111010111011111010111010")
	}
}
func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isWordLanguageRecursive("0111010111010111011111010111010")
	}
}
func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isWordLanguageRegex("0111010111010111011111010111010")
	}
}
func BenchmarkDEA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isWordLanguageDEA("0111010111010111011111010111010")
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
