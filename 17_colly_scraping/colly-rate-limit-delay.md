# colly rate limit delay

## Live interview task
Add polite delay and random delay between requests.

## Concepts covered
- Colly
- rate limiting
- RandomDelay

## Candidate solution

```go
package main

import (
    "time"
    "log"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.Async(true))
    c.Limit(&colly.LimitRule{DomainGlob: "*example.*", Parallelism: 1, Delay: time.Second, RandomDelay: 500 * time.Millisecond})
    c.OnRequest(func(r *colly.Request) { log.Println("visit", r.URL) })
    c.Visit("https://example.com")
    c.Wait()
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
