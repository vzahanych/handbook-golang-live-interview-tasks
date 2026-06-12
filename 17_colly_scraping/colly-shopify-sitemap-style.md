# colly shopify sitemap style

## Live interview task
Parse sitemap URLs and visit product URLs.

## Concepts covered
- Colly
- XML callbacks
- sitemaps

## Candidate solution

```go
package main

import (
    "fmt"
    "strings"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()
    c.OnXML("//url/loc", func(e *colly.XMLElement) {
        u := strings.TrimSpace(e.Text)
        if strings.Contains(u, "/products/") { fmt.Println("product", u) }
    })
    _ = c.Visit("https://example.com/sitemap.xml")
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
