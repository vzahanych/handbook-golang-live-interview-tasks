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
- `OnXML` not OnHTML — sitemaps are XML; namespace prefixes may affect XPath.
- Sitemap index files list other sitemaps — handle `sitemapindex` vs `urlset` recursively.
- Filter `/products/` before Visit — avoids crawling blog/policy URLs.

## Q&A

**Q: OnXML vs OnHTML?**  
A: OnXML for RSS/Atom/sitemap; OnHTML for HTML DOM.

**Q: Complexity?**  
A: O(urls in sitemap) parse; visiting each product is separate HTTP cost.

**Q: Edge cases?**  
A: Multi-lang hreflang URLs, gzip sitemap `.xml.gz`, 50k URL limit per file.

**Q: Visit from OnXML?**  
A: `c.Visit(u)` enqueue product pages — combine with queue + rate limit.

**Q: Production?**  
A: Shopify `/sitemap.xml` + JSON endpoints; respect Crawl-delay.
