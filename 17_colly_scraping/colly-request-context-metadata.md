# colly request context metadata

## Live interview task
Attach metadata to requests and read it in callbacks.

## Concepts covered
- Colly
- request context

## Candidate solution

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()
    c.OnResponse(func(r *colly.Response) { fmt.Println("source", r.Ctx.Get("source"), "url", r.Request.URL) })
    ctx := colly.NewContext()
    ctx.Put("source", "seed")
    c.Request("GET", "https://example.com", nil, ctx, nil)
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
