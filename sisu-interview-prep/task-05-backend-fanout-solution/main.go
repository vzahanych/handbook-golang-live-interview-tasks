package main

// 1. Make the fetches CONCURRENT (fan-out) and collect all results (fan-in),
//    preserving which URL produced which result. Use a sync.WaitGroup (or
//    errgroup).
//    Spawn workers (bounded) and have each worker write exactly one result slot
//    by index. A WaitGroup waits for completion.
// 2. A single slow back-end currently hangs the whole call forever. Add a
//    context deadline and apply it PER request (http.NewRequestWithContext) so
//    slow back-ends are cancelled and no goroutines leak.
//    Use context.WithTimeout and create each request with
//    http.NewRequestWithContext so cancellation interrupts the HTTP call.
// 3. Errors are dropped and response bodies are never closed — a
//    file-descriptor and memory leak under load, and a nil-pointer panic when a
//    request fails. Fix the resource handling (defer resp.Body.Close()) and
//    return partial results plus the per-URL errors.
//    Always check err before touching resp, and always close resp.Body. Return a
//    Result for every URL: success with Bytes or failure with Err.
// 4. With hundreds of back-ends, unbounded goroutines and connections overwhelm
//    both the host and the downstreams. Bound the concurrency (a semaphore or
//    worker pool) and tune the http.Transport (e.g. MaxIdleConnsPerHost).
//    Use a semaphore channel to cap in-flight requests and a Transport with
//    MaxIdleConnsPerHost so connections are reused safely under load.

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
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

// func fetchAll(urls []string) []Result {
func fetchAll(ctx context.Context, client *http.Client, urls []string, perRequestTimeout time.Duration, maxInFlight int, failFast bool) ([]Result, error) {
	if maxInFlight <= 0 {
		maxInFlight = 1
	}
	results := make([]Result, len(urls)) // answer 1/3: stable slots preserve URL->result by index

	sem := make(chan struct{}, maxInFlight) // answer 4: bound concurrency
	var wg sync.WaitGroup

	errCh := make(chan error, 1)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i, u := range urls {
		wg.Add(1)
		go func(i int, u string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			reqCtx, cancelReq := context.WithTimeout(ctx, perRequestTimeout) // answer 2
			defer cancelReq()

			req, _ := http.NewRequestWithContext(reqCtx, http.MethodGet, u, nil) // answer 2
			resp, err := client.Do(req)
			if err != nil {
				results[i] = Result{URL: u, Err: err}
				if failFast {
					select {
					case errCh <- err:
						cancel()
					default:
					}
				}
				return
			}
			defer func() { _ = resp.Body.Close() }() // answer 3

			// answer 3: stream instead of ReadAll to avoid big allocations; bytes counted.
			n, copyErr := io.Copy(io.Discard, resp.Body)
			if copyErr != nil {
				results[i] = Result{URL: u, Err: copyErr}
				if failFast {
					select {
					case errCh <- copyErr:
						cancel()
					default:
					}
				}
				return
			}

			if resp.StatusCode >= 400 {
				err := fmt.Errorf("status %d", resp.StatusCode)
				results[i] = Result{URL: u, Bytes: int(n), Err: err}
				if failFast {
					select {
					case errCh <- err:
						cancel()
					default:
					}
				}
				return
			}

			results[i] = Result{URL: u, Bytes: int(n)}
		}(i, u)
	}

	wg.Wait()
	close(errCh)
	if failFast {
		if err, ok := <-errCh; ok && err != nil {
			return results, err
		}
	}
	return results, nil
}

func main() {
	// Sample data (paste from the meeting chat) — the last one is the slow one:
	urls := []string{
		"https://example.com",
		"https://www.google.com",
		"https://httpbin.org/delay/10",
	}

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 50, // answer 4
			IdleConnTimeout:     30 * time.Second,
		},
	}

	results, err := fetchAll(context.Background(), client, urls, 2*time.Second, 10, false)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		fmt.Printf("fanout error: %v\n", err)
	}

	for _, r := range results {
		if r.Err != nil {
			fmt.Printf("%s: ERROR %v\n", r.URL, r.Err)
			continue
		}
		fmt.Printf("%s: %d bytes\n", r.URL, r.Bytes)
	}
}

// How to run and test:
//   go run .
//   go test ./sisu-interview-prep/task-05-backend-fanout-solution -v
//   go test ./sisu-interview-prep/task-05-backend-fanout-solution -bench=. -benchmem
