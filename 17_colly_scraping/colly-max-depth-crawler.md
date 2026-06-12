# colly max depth crawler

## Live interview task
Crawl links but limit recursion depth.

## Concepts covered
- Colly
- MaxDepth
- AbsoluteURL

## Candidate solution

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.AllowedDomains("go.dev"), colly.MaxDepth(2))
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Request.AbsoluteURL(e.Attr("href"))
        if link != "" { e.Request.Visit(link) }
    })
    c.OnRequest(func(r *colly.Request) { fmt.Println("depth", r.Depth, r.URL.String()) })
    c.Visit("https://go.dev/")
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
