package main

import (
	"fmt"
	"io"
	"net/http"
)

// The Central Integration Layer (CIL) calls many downstream back-ends to
// assemble one response. This starter fetches them one at a time, ignores
// errors, never closes bodies, and has no timeout — it covers the "http
// requests / memory resource issues / high throughput" thread in the briefing.
//
// 1. Make the fetches CONCURRENT (fan-out) and collect all results (fan-in),
//    preserving which URL produced which result. Use a sync.WaitGroup (or
//    errgroup).
//
// 2. A single slow back-end currently hangs the whole call forever. Add a
//    context deadline and apply it PER request (http.NewRequestWithContext) so
//    slow back-ends are cancelled and no goroutines leak.
//
// 3. Errors are dropped and response bodies are never closed — a
//    file-descriptor and memory leak under load, and a nil-pointer panic when a
//    request fails. Fix the resource handling (defer resp.Body.Close()) and
//    return partial results plus the per-URL errors.
//
// 4. With hundreds of back-ends, unbounded goroutines and connections overwhelm
//    both the host and the downstreams. Bound the concurrency (a semaphore or
//    worker pool) and tune the http.Transport (e.g. MaxIdleConnsPerHost).

type Result struct {
	URL   string
	Bytes int
	Err   error
}

func fetchAll(urls []string) []Result {
	results := make([]Result, 0, len(urls))
	for _, u := range urls {
		resp, _ := http.Get(u)           // err ignored; no timeout; nil resp on failure
		body, _ := io.ReadAll(resp.Body) // body never closed ⇒ leak
		results = append(results, Result{URL: u, Bytes: len(body)})
	}
	return results
}

func main() {
	// Sample data (paste from the meeting chat) — the last one is the slow one:
	urls := []string{
		"https://example.com",
		"https://www.google.com",
		"https://httpbin.org/delay/10",
	}

	for _, r := range fetchAll(urls) {
		if r.Err != nil {
			fmt.Printf("%s: ERROR %v\n", r.URL, r.Err)
			continue
		}
		fmt.Printf("%s: %d bytes\n", r.URL, r.Bytes)
	}
}
