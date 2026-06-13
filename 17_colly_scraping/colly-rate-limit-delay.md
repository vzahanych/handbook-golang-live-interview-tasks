# colly rate limit delay

## Live interview task
Add polite delay and random delay between requests.

## Concepts covered
- Colly
- rate limiting
- RandomDelay

## Candidate solution

```go
package main

import (
    "time"
    "log"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.Async(true))
    c.Limit(&colly.LimitRule{DomainGlob: "*example.*", Parallelism: 1, Delay: time.Second, RandomDelay: 500 * time.Millisecond})
    c.OnRequest(func(r *colly.Request) { log.Println("visit", r.URL) })
    c.Visit("https://example.com")
    c.Wait()
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- `Async(true)` required for parallelism — must call `c.Wait()` or main exits early.
- `LimitRule` is per domain glob — multiple rules can apply; `Parallelism: 1` serializes.
- `RandomDelay` adds jitter on top of base `Delay` — reduces thundering herd patterns.

## Q&A

**Q: Why politeness matters?**  
A: Avoid IP bans; many sites rate-limit aggressive crawlers.

**Q: Complexity?**  
A: Crawl time grows with delay × request count.

**Q: Edge cases?**  
A: Subdomains — separate LimitRule per `*.api.example.com`.

**Q: vs global sleep?**  
A: LimitRule integrates with Colly scheduler per domain.

**Q: Production?**  
A: Read Crawl-delay from robots.txt; backoff on 429 Retry-After.
