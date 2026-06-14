# bounded crawler standard library

## Live interview task
Build bounded concurrent HTTP fetcher using stdlib only (no Colly).

## Concepts covered
- http.Client
- worker pool
- semaphore

## Candidate solution

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "sync"
    "time"
)

func crawl(urls []string, workers int) {
    jobs := make(chan string)
    var wg sync.WaitGroup
    client := &http.Client{Timeout: 10 * time.Second}

    for w := 0; w < workers; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for u := range jobs {
                resp, err := client.Get(u)
                if err != nil {
                    fmt.Println(u, err)
                    continue
                }
                _, _ = io.Copy(io.Discard, resp.Body)
                resp.Body.Close()
                fmt.Println(u, resp.StatusCode)
            }
        }()
    }

    for _, u := range urls {
        jobs <- u
    }
    close(jobs)
    wg.Wait()
}

func main() {
    crawl([]string{"https://example.com", "https://go.dev"}, 2)
}
```

## Run

Runnable version lives in [bounded-crawler/](bounded-crawler/main.go).

```bash
# default URLs (example.com, go.dev)
go run ./18_mini_projects/bounded-crawler

# or pass your own URLs
go run ./18_mini_projects/bounded-crawler https://example.com https://go.dev
```

## Interview notes / pitfalls
- Worker pool bounds concurrency — polite vs unbounded goroutines per URL.
- Always close response body — connection reuse.
- Respect `robots.txt` and rate limits in production.
- Extend: parse HTML links, visited set, max depth BFS.

## Q&A

**Q: vs Colly?**  
A: Colly handles cookies, callbacks, queue — stdlib minimal control.

**Q: Context cancel?**  
A: `NewRequestWithContext` per fetch.

**Q: Redirects?**  
A: Default client follows — limit with CheckRedirect.

**Q: Duplicate URLs?**  
A: `map[string]struct{}` visited set.

**Q: Complexity?**  
A: O(urls) tasks with W parallelism.
