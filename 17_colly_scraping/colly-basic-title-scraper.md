# colly basic title scraper

## Live interview task
Visit a page and extract its title with a Colly OnHTML callback.

## Concepts covered
- Colly
- OnHTML
- AllowedDomains

## Candidate solution

```go
package main

import (
    "fmt"
    "log"

    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.AllowedDomains("example.com"))
    c.OnHTML("title", func(e *colly.HTMLElement) { fmt.Println(e.Text) })
    c.OnRequest(func(r *colly.Request) { log.Println("visiting", r.URL) })
    if err := c.Visit("https://example.com"); err != nil { log.Fatal(err) }
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Expected output

```
visiting https://example.com
Example Domain
```

## Interview notes / pitfalls
- `AllowedDomains` blocks off-domain redirects — without it Colly may follow external links.
- `OnHTML` runs per matching element — multiple `<title>` tags mean multiple callbacks.
- Colly uses goquery under the hood — CSS selectors, not full XPath by default.

## Q&A

**Q: Collector vs NewCollector options?**  
A: Options set defaults (domain, depth, async); child collectors inherit with `c.Clone()`.

**Q: Complexity?**  
A: O(page size) DOM walk; one HTTP request here.

**Q: Edge cases?**  
A: Empty title, JS-rendered title (needs headless browser), charset decoding.

**Q: Production?**  
A: Timeouts, retries, User-Agent, respect robots.txt, rate limits.

**Q: Test without network?**  
A: `colly-local-html-file` pattern or httptest server serving HTML.
