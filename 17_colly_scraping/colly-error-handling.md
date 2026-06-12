# colly error handling

## Live interview task
Handle request errors using OnError and log the failed URL and status code.

## Concepts covered
- Colly
- OnError

## Candidate solution

```go
package main

import (
    "log"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()
    c.OnError(func(r *colly.Response, err error) {
        status := 0
        if r != nil { status = r.StatusCode }
        log.Printf("failed url=%v status=%d err=%v", r.Request.URL, status, err)
    })
    _ = c.Visit("https://example.com/not-found")
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
