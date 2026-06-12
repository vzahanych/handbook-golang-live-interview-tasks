# colly hackernews comments style

## Live interview task
Extract comment-like rows from a page using CSS selectors.

## Concepts covered
- Colly
- CSS selectors
- real-life scraping pattern

## Candidate solution

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

type Comment struct { Author, Text string }

func main() {
    var comments []Comment
    c := colly.NewCollector()
    c.OnHTML(".comment", func(e *colly.HTMLElement) {
        comments = append(comments, Comment{Author: e.ChildText(".author"), Text: e.ChildText(".text")})
    })
    c.OnScraped(func(r *colly.Response) { fmt.Println("comments", len(comments)) })
    _ = c.Visit("https://news.ycombinator.com/item?id=1")
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
