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
- `_ = courses[:0]` resets slice length keeping capacity — reuse buffer across pages in multi-page crawl.
- Example uses `example.com` — real Coursera is dynamic; pattern is what interview tests.
- Map DOM fields to struct early — cleaner JSON export downstream.

## Q&A

**Q: Pointer vs value in slice?**  
A: `[]Course` values fine for small structs; `[]*Course` if mutating after scrape.

**Q: Complexity?**  
A: O(cards) per listing page.

**Q: Edge cases?**  
A: Missing rating text, internationalized titles, sponsored cards duplicate selectors.

**Q: OnScraped timing?**  
A: Fires after all OnHTML for that response — safe to print `len(courses)`.

**Q: Production?**  
A: Schema validation, dedupe by course URL, incremental crawl by category sitemap.
