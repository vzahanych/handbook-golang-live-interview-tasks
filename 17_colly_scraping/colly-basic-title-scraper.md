# colly basic title scraper

## Live interview task
Visit a page and extract its title with a Colly OnHTML callback.

## Concepts covered
- Colly
- OnHTML
- AllowedDomains

## Candidate solution

```go
package main

import (
    "fmt"
    "log"

    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.AllowedDomains("example.com"))
    c.OnHTML("title", func(e *colly.HTMLElement) { fmt.Println(e.Text) })
    c.OnRequest(func(r *colly.Request) { log.Println("visiting", r.URL) })
    if err := c.Visit("https://example.com"); err != nil { log.Fatal(err) }
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
