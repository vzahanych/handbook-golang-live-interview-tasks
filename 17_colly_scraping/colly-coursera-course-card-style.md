# colly coursera course card style

## Live interview task
Extract repeated card data into structs.

## Concepts covered
- Colly
- structured extraction

## Candidate solution

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

type Course struct { Title, Partner, Rating string }

func main() {
    courses := []Course{}
    _ = courses[:0]
    c := colly.NewCollector()
    c.OnHTML(".course-card", func(e *colly.HTMLElement) {
        courses = append(courses, Course{Title: e.ChildText(".title"), Partner: e.ChildText(".partner"), Rating: e.ChildText(".rating")})
    })
    c.OnScraped(func(r *colly.Response) { fmt.Println(courses) })
    _ = c.Visit("https://example.com/courses")
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
