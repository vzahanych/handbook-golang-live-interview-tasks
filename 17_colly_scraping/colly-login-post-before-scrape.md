# colly login post before scrape

## Live interview task
Authenticate with POST before scraping authenticated pages.

## Concepts covered
- Colly
- Post
- cookies/session

## Candidate solution

```go
package main

import (
    "log"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()
    if err := c.Post("https://example.com/login", map[string]string{"username":"admin", "password":"admin"}); err != nil {
        log.Println("login failed:", err)
    }
    c.OnResponse(func(r *colly.Response) { log.Println("status", r.StatusCode) })
    _ = c.Visit("https://example.com/account")
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
