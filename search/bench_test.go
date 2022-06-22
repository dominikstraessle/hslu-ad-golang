package search

import (
	"math/rand"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var digits = []rune("012345")

var letterText = randSeq(100_000, letters)
var digitText = randSeq(100_000, digits)

func randSeq(n int, alphabet []rune) string {
	rand.Seed(1) // we want the same "random" sequence for each benchmark
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(alphabet))]
	}
	return string(b)
}

func BenchmarkQuickSearch(b *testing.B) {
	pattern := letterText[99_995:]
	result := -1
	for i := 0; i < b.N; i++ {
		result = QuickSearch(letterText, pattern)
	}
	b.Logf("pattern: %v\nresult: %v", pattern, result)
}

func BenchmarkKMPSearch(b *testing.B) {
	pattern := letterText[99_995:]
	result := -1
	for i := 0; i < b.N; i++ {
		result = KMPSearch(letterText, pattern)
	}
	b.Logf("pattern: %v\nresult: %v", pattern, result)
}
func BenchmarkQuickSearchDigits(b *testing.B) {
	pattern := digitText[99_990:]
	result := -1
	for i := 0; i < b.N; i++ {
		result = QuickSearch(digitText, pattern)
	}
	b.Logf("pattern: %v\nresult: %v", pattern, result)
}

func BenchmarkKMPSearchDigits(b *testing.B) {
	pattern := digitText[99_990:]
	result := -1
	for i := 0; i < b.N; i++ {
		result = KMPSearch(digitText, pattern)
	}
	b.Logf("pattern: %v\nresult: %v", pattern, result)
}
