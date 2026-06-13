# colly robots txt etiquette

## Live interview task
Enable robots.txt checking and polite crawling defaults before scraping a site.

## Concepts covered
- Colly
- robots.txt
- User-Agent
- Crawl-delay

## Candidate solution

```go
package main

import (
    "log"
    "time"

    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("example.com"),
        colly.UserAgent("MyBot/1.0 (+https://example.com/bot)"),
    )

    // Colly fetches and respects robots.txt per domain by default.
    c.Limit(&colly.LimitRule{
        DomainGlob:  "*example.*",
        Parallelism: 1,
        Delay:       2 * time.Second,
    })

    c.OnRequest(func(r *colly.Request) {
        log.Println("allowed visit", r.URL)
    })

    if err := c.Visit("https://example.com"); err != nil {
        log.Fatal(err)
    }
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- Colly checks robots.txt automatically — disallowed paths return error/skip depending on version; verify behavior in docs.
- Identify your bot with honest `User-Agent` including contact URL — not a generic browser string.
- `Crawl-delay` in robots.txt should inform `LimitRule.Delay` — some sites specify 10s+.
- Legal/ethical: permission + ToS matter beyond robots.txt — robots is not a law, but industry norm.

## Q&A

**Q: What if robots.txt disallows `/`?**  
A: Do not scrape — use official API or get written permission.

**Q: Complexity?**  
A: One robots.txt fetch per host cached by Colly for session.

**Q: Edge cases?**  
A: robots.txt on CDN vs origin mismatch, wildcard `Disallow`, sitemap lines in robots.

**Q: Override for testing?**  
A: Use local fixtures or sites that allow crawling (e.g. httpbin) — never bypass production robots in prod.

**Q: Production checklist?**  
A: User-Agent, rate limit, robots respect, error metrics, documented data use, opt-out contact.
