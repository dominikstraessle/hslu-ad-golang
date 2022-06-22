package search

import (
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var digits = []rune("012345")

var letterText = randSeq(100_000, letters)
var digitText = randSeq(100_000, digits)
var patternLetters = letterText[99_980:]
var patternDigits = digitText[99_980:]

func randSeq(n int, alphabet []rune) string {
	rand.Seed(1) // we want the same "random" sequence for each benchmark
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(alphabet))]
	}
	return string(b)
}

func BenchmarkQuickSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSearch(letterText, patternLetters)
	}
}

func BenchmarkKMPSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KMPSearch(letterText, patternLetters)
	}
}

func BenchmarkSimpleSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleSearch(letterText, patternLetters)
	}
}

func BenchmarkQuickSearchDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSearch(digitText, patternDigits)
	}
}

func BenchmarkKMPSearchDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KMPSearch(digitText, patternDigits)
	}
}

func BenchmarkSimpleSearchDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleSearch(digitText, patternDigits)
	}
}
