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
- HN HTML structure changes — selectors break; prefer stable `data-*` or API when allowed.
- Append to slice in OnHTML from async collector — mutex if `Async(true)`.
- `OnScraped` runs once per page — good place to flush batch to DB.

## Q&A

**Q: ChildText vs Text?**  
A: `ChildText(".author")` scopes to subtree; `Text` is full element text.

**Q: Complexity?**  
A: O(comments on page) DOM nodes visited.

**Q: Edge cases?**  
A: Deleted comments, empty author, nested reply threads need recursive Visit.

**Q: Pagination?**  
A: Follow "More" link or use official Firebase API for HN.

**Q: Production?**  
A: Rate limit, store raw HTML snapshot, monitor selector drift alerts.
