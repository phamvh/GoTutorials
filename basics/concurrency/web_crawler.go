package main

import (
	"fmt"
	"sync"
)

/**
This example demos the use of mutex.
This is an exercise from Go tutorial; https://go.dev/tour/concurrency/10
In this exercise you'll use Go's concurrency features to parallelize a web crawler.

Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.

Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!
*/

// Fetcher has method Fetch to get data given a url.
// This can be faked later, just assume we have this to use.
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	// Note thta you can even name the return when declaring a method: body, urls, err. Things look clearer.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, waitGroup *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	track := CrawlTrack{
		crawled: make(map[string]bool),
	}
	waitGroup.Add(1) // add a worker for each goroutine
	go crawlRecursive(url, depth, fetcher, &track, waitGroup)
}

type CrawlTrack struct {
	crawled map[string]bool
	mu      sync.Mutex // no need to initialize this one. It can be used right away.
}

func crawlRecursive(
	url string,
	depth int,
	fetcher Fetcher,
	track *CrawlTrack,
	waitGroup *sync.WaitGroup) {

	// The use of defer here is great. We want to call Done() whenever this function returns.
	// There are so many places in this function where "return" may happen, and scattering
	// Done() around can make code look messy.
	// So defer is a perfect solution here.
	defer waitGroup.Done()

	if depth < 0 {
		return
	}
	track.mu.Lock()
	_, ok := track.crawled[url]
	if ok {
		track.mu.Unlock()
		return
	}
	track.crawled[url] = true
	track.mu.Unlock()
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		track.mu.Lock()
		_, ok2 := track.crawled[u]
		track.mu.Unlock()
		if !ok2 {
			waitGroup.Add(1)
			go crawlRecursive(u, depth-1, fetcher, track, waitGroup)
		}
	}
}

func main() {
	waitGroup := sync.WaitGroup{} // or you can write: `var waitGroup sync.WaitGroup`
	Crawl("https://golang.org/", 4, fetcher, &waitGroup)
	waitGroup.Wait()
}

// Some fake fetchers. Do NOT pay attention to this. Just assume they exist for use.
// fakeFetcher is Fetcher that returns canned results.
type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

// add method Fetch to fakeFetcher to make it implement Fetcher interface
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	// try to get data from the map; if exits (ok), then return it, else return an error
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.

// fakeFetcher is a map, so initiate it with some data
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
