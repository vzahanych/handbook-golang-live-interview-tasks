# colly max depth crawler

## Live interview task
Crawl links but limit recursion depth.

## Concepts covered
- Colly
- MaxDepth
- AbsoluteURL

## Candidate solution

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.AllowedDomains("go.dev"), colly.MaxDepth(2))
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Request.AbsoluteURL(e.Attr("href"))
        if link != "" { e.Request.Visit(link) }
    })
    c.OnRequest(func(r *colly.Request) { fmt.Println("depth", r.Depth, r.URL.String()) })
    c.Visit("https://go.dev/")
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- `MaxDepth` counts from seed (depth 0) — off-by-one confusion in interviews.
- Without `AllowedDomains` + visited tracking, crawler may loop on calendar/archive pages.
- `AbsoluteURL` resolves relative hrefs — required before `Visit`.

## Q&A

**Q: BFS vs DFS?**  
A: Colly default is breadth-first via internal queue.

**Q: Complexity?**  
A: O(pages × links per page) bounded by depth and domain filter.

**Q: Duplicate URLs?**  
A: Colly deduplicates visited URLs per collector.

**Q: Edge cases?**  
A: Fragment-only links (`#section`), `mailto:`, query-string duplicates.

**Q: Production?**  
A: Sitemap seed, politeness limits, persistent frontier (DB queue).
