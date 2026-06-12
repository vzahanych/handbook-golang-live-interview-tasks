# colly xkcd store items style

## Live interview task
Extract item names and prices from a store grid.

## Concepts covered
- Colly
- store scraping pattern

## Candidate solution

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()
    c.OnHTML(".product", func(e *colly.HTMLElement) {
        fmt.Printf("%s %s\n", e.ChildText(".name"), e.ChildText(".price"))
    })
    _ = c.Visit("https://store.xkcd.com/")
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
