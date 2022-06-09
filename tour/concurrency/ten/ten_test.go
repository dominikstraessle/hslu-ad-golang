package main

import "testing"

const baseUrl = "https://golang.org/"

func BenchmarkCrawl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Crawl(baseUrl, 4, fetcher)
	}
}

func BenchmarkConcurrentCrawl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentCrawl(baseUrl, 4, fetcher)
	}
}
