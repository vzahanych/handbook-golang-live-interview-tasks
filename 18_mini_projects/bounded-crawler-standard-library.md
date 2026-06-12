# bounded crawler standard library

## Live interview task
Build a small bounded crawler using net/http and channels.

## Concepts covered
- HTTP client
- worker pool
- bounded concurrency

## Candidate solution

```go
package main

import (
    "fmt"
    "net/http"
    "sync"
)

func crawl(urls []string, workers int) {
    jobs := make(chan string)
    var wg sync.WaitGroup
    client := &http.Client{}
    for w:=0; w<workers; w++ { wg.Add(1); go func(){ defer wg.Done(); for u := range jobs { resp, err := client.Get(u); if err != nil { fmt.Println(u, err); continue }; fmt.Println(u, resp.StatusCode); resp.Body.Close() } }() }
    for _, u := range urls { jobs <- u }
    close(jobs); wg.Wait()
}

func main() { crawl([]string{"https://example.com", "https://go.dev"}, 2) }
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
