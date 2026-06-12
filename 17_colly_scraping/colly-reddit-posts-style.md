# colly reddit posts style

## Live interview task
Collect post titles and links from a listing page.

## Concepts covered
- Colly
- links
- AbsoluteURL

## Candidate solution

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.AllowedDomains("old.reddit.com"))
    c.OnHTML("a.title", func(e *colly.HTMLElement) {
        fmt.Println(e.Text, e.Request.AbsoluteURL(e.Attr("href")))
    })
    _ = c.Visit("https://old.reddit.com/r/golang/")
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
