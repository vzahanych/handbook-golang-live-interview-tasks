# colly reddit posts style

## Live interview task
Collect post titles and links from a listing page.

## Concepts covered
- Colly
- links
- AbsoluteURL

## Candidate solution

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.AllowedDomains("old.reddit.com"))
    c.OnHTML("a.title", func(e *colly.HTMLElement) {
        fmt.Println(e.Text, e.Request.AbsoluteURL(e.Attr("href")))
    })
    _ = c.Visit("https://old.reddit.com/r/golang/")
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- Reddit rate-limits and blocks raw scrapers — prefer official API for production.
- `old.reddit.com` serves simpler HTML than new Reddit (JS-heavy).
- `AbsoluteURL` required — hrefs often relative `/r/golang/comments/...`.

## Q&A

**Q: AllowedDomains?**  
A: Stops following promoted external links off-domain.

**Q: Complexity?**  
A: O(posts on listing page) per visit.

**Q: Edge cases?**  
A: Stickied posts, ads with similar selectors, NSFW quarantine pages.

**Q: Next page?**  
A: `OnHTML` on `.next-button a` → Visit next listing URL.

**Q: Production?**  
A: OAuth API, respect robots.txt, identifiable User-Agent.
