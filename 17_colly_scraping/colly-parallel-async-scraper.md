# colly parallel async scraper

## Live interview task
Run asynchronous visits and wait for completion.

## Concepts covered
- Colly
- Async
- LimitRule
- Wait

## Candidate solution

```go
package main

import (
    "fmt"
    "log"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector(colly.Async(true))
    c.Limit(&colly.LimitRule{DomainGlob: "*httpbin.*", Parallelism: 2})
    c.OnResponse(func(r *colly.Response) { fmt.Println(r.Request.URL, r.StatusCode) })
    for i := 0; i < 5; i++ {
        if err := c.Visit(fmt.Sprintf("https://httpbin.org/delay/1?n=%d", i)); err != nil { log.Println(err) }
    }
    c.Wait()
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- Missing `c.Wait()` — program exits before async visits complete (classic bug).
- `Parallelism: 2` caps concurrent requests per domain glob — not global goroutine count.
- Shared slice append in OnResponse needs mutex if multiple callbacks write concurrently.

## Q&A

**Q: Speedup with 5 URLs, parallelism 2?**  
A: ~3 waves if each takes 1s — not 5× faster.

**Q: Complexity?**  
A: O(n/p) wall time idealized; memory O(n) in-flight responses.

**Q: Edge cases?**  
A: One slow URL blocks a slot until timeout — set `Collector.Timeout`.

**Q: vs worker pool?**  
A: Colly LimitRule is domain-aware scheduler built-in.

**Q: Production?**  
A: Context cancel on shutdown, Wait with timeout wrapper.
