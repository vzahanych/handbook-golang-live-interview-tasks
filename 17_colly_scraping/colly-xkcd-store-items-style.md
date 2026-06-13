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
- E-commerce grids repeat structure — one OnHTML handler per `.product` card is idiomatic.
- Price strings need parsing (`$19.00` → cents) — don't store as float money.
- Out-of-stock items may omit price — handle empty ChildText.

## Q&A

**Q: Struct slice vs println?**  
A: Collect `[]Product` in OnHTML, emit in OnScraped — same as coursera pattern.

**Q: Complexity?**  
A: O(products) on page.

**Q: Edge cases?**  
A: Sale vs regular price, variant SKUs, lazy-loaded images (not in static HTML).

**Q: Pagination?**  
A: Follow rel=next or crawl collection sitemap.

**Q: Production?**  
A: Shopify JSON API often easier than HTML for Shopify stores.
