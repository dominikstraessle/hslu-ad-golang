package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	body   string
	url    string
	status bool
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func ConcurrentCrawl(url string, depth int, fetcher Fetcher) []*Result {
	ch := make(chan *Result)

	// start the crawling with a given super channel
	go concurrentCrawl(url, depth, fetcher, ch)

	// receive all values until channel is closed
	// append all values to string slice
	var urls []*Result
	for u := range ch {
		urls = append(urls, u)
	}
	return urls
}

func concurrentCrawl(url string, depth int, fetcher Fetcher, superChan chan *Result) {
	// no depth -> immediate return
	if depth <= 0 {
		close(superChan)
		return
	}

	// store the url in the cache -> only fetch every url once
	cache.Store(url, nil)

	// fetch the given URL
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		// send the current url to the super channel
		superChan <- &Result{
			body:   body,
			url:    url,
			status: false,
		}

		// close the current channel on error
		// indicates that this goroutine is finished with fetching
		close(superChan)
		return
	}

	// send the current url to the super channel
	superChan <- &Result{
		body:   body,
		url:    url,
		status: true,
	}

	// for each sub url -> create a new sub channel and fetch all sub url via a recursive concurrent method call
	var subChannels []chan *Result
	for _, u := range urls {
		if _, ok := cache.Load(u); ok {
			continue
		}
		subChan := make(chan *Result)
		go concurrentCrawl(u, depth-1, fetcher, subChan)
		subChannels = append(subChannels, subChan)
	}

	// loop over all sub channels and retrieve their values
	for _, subChan := range subChannels {
		for subUrl := range subChan {
			superChan <- subUrl
		}
	}

	// close the super channel to indicate that this goroutine is finished
	close(superChan)
}

var cache sync.Map

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	results := ConcurrentCrawl("https://golang.org/", 4, fetcher)
	for _, result := range results {
		fmt.Println(result)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Millisecond * 100)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
