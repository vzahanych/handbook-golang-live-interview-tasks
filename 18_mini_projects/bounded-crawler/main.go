// Command bounded-crawler is a bounded concurrent HTTP fetcher built with the
// standard library only (no Colly). A fixed pool of workers drains a jobs
// channel, which bounds concurrency the way a semaphore would.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// crawl fetches every URL in urls using exactly `workers` goroutines running
// concurrently. The number of in-flight HTTP requests never exceeds `workers`,
// no matter how many URLs are passed in — that is the "bounded" part.
func crawl(urls []string, workers int) {
	// jobs is the work queue. It is unbuffered, so a send (jobs <- u) blocks
	// until some worker is ready to receive. That natural back-pressure means
	// we never build up a huge in-memory backlog of pending URLs.
	jobs := make(chan string)

	// wg lets main wait until every worker goroutine has fully finished before
	// returning. Without it, crawl could return (and the program exit) while
	// requests are still in flight.
	var wg sync.WaitGroup

	// One shared client is reused by all workers. http.Client is safe for
	// concurrent use and pools/reuses TCP connections under the hood, so
	// sharing it is both correct and more efficient than one client per request.
	// Timeout caps the total time per request (dial + redirects + reading the
	// body); without it a slow/hung server could block a worker forever.
	client := &http.Client{Timeout: 10 * time.Second}

	// Start the worker pool. Each iteration launches one long-lived goroutine
	// that keeps pulling jobs until the channel is closed and drained.
	for w := 0; w < workers; w++ {
		wg.Add(1) // register this worker before it starts, so Wait() sees it
		go func() {
			defer wg.Done() // signal completion when this worker returns

			// range over a channel receives values until the channel is closed
			// AND empty, then the loop ends. This is how a worker knows there
			// is no more work to do.
			for u := range jobs {
				// Get performs a GET request. err is non-nil for transport-level
				// failures (DNS, connection refused, timeout, ...) — NOT for HTTP
				// status codes like 404 or 500, which are successful responses.
				resp, err := client.Get(u)
				if err != nil {
					fmt.Println(u, err)
					continue // skip to the next URL; nothing to close on error
				}

				// We must drain and close the body. Copying it to io.Discard
				// reads the response to completion so the underlying TCP
				// connection can be returned to the pool and reused (keep-alive).
				_, _ = io.Copy(io.Discard, resp.Body)
				// Closing the body releases the connection. Forgetting this is a
				// classic resource leak that eventually exhausts file descriptors.
				resp.Body.Close()

				fmt.Println(u, resp.StatusCode)
			}
		}()
	}

	// Feed the queue. Because jobs is unbuffered, this loop hands one URL at a
	// time to whichever worker is free, blocking when all workers are busy.
	for _, u := range urls {
		jobs <- u
	}

	// Closing the channel tells the workers "no more jobs are coming". Their
	// `for range jobs` loops finish once the channel is drained. Closing is the
	// producer's job and must happen exactly once, after the last send.
	close(jobs)

	// Block until all workers have returned (all bodies read, all prints done).
	wg.Wait()
}

func main() {
	// Take URLs from the command line (everything after the program name).
	urls := os.Args[1:]
	if len(urls) == 0 {
		// Sensible defaults so the example runs with no arguments.
		urls = []string{"https://example.com", "https://go.dev"}
	}
	// Two workers: with the two default URLs they are fetched in parallel.
	crawl(urls, 2)
}
