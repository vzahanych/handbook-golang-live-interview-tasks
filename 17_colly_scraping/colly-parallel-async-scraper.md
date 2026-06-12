# colly parallel async scraper

## Live interview task
Run asynchronous visits and wait for completion.

## Concepts covered
- Colly
- Async
- LimitRule
- Wait

## Candidate solution

```go
package main

import (
    "fmt"
    "log"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.Async(true))
    c.Limit(&colly.LimitRule{DomainGlob: "*httpbin.*", Parallelism: 2})
    c.OnResponse(func(r *colly.Response) { fmt.Println(r.Request.URL, r.StatusCode) })
    for i := 0; i < 5; i++ {
        if err := c.Visit(fmt.Sprintf("https://httpbin.org/delay/1?n=%d", i)); err != nil { log.Println(err) }
    }
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
