# colly url filter

## Live interview task
Allow only selected URLs by matching path patterns before visiting.

## Concepts covered
- Colly
- URLFilters
- regexp

## Candidate solution

```go
package main

import (
    "fmt"
    "regexp"
    "github.com/gocolly/colly/v2"
)

func main() {
    allowed := regexp.MustCompile(`/docs/|/$`)
    c := colly.NewCollector(colly.AllowedDomains("go.dev"), colly.URLFilters(allowed))
    c.OnRequest(func(r *colly.Request) { fmt.Println("visit", r.URL) })
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
