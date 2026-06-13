# colly scraper server endpoint

## Live interview task
Expose a tiny HTTP endpoint that runs a Colly scrape on demand.

## Concepts covered
- Colly
- HTTP server
- JSON

## Candidate solution

```go
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gocolly/colly/v2"
)

func scrapeTitle(url string) (string, error) {
    c := colly.NewCollector()
    title := ""
    c.OnHTML("title", func(e *colly.HTMLElement) { title = e.Text })
    return title, c.Visit(url)
}

func main() {
    http.HandleFunc("/title", func(w http.ResponseWriter, r *http.Request) {
        title, err := scrapeTitle(r.URL.Query().Get("url"))
        if err != nil { http.Error(w, err.Error(), 500); return }
        json.NewEncoder(w).Encode(map[string]string{"title": title})
    })
    http.ListenAndServe(":8080", nil)
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- New collector per request is simple but heavy — pool collectors or set global timeout.
- **SSRF risk** — validate `url` query param (block localhost, file://, internal IPs).
- Scrape in request handler blocks HTTP worker — use job queue for slow targets.

## Q&A

**Q: SSRF mitigation?**  
A: Parse URL, allowlist schemes/https hosts, reject private IP ranges.

**Q: Complexity?**  
A: O(scrape time) per API call — unbounded if target slow.

**Q: Edge cases?**  
A: Missing url param, invalid URL, pages without title.

**Q: Concurrency?**  
A: Each request spawns collector — cap with semaphore or worker pool.

**Q: Production?**  
A: Auth on endpoint, cache titles, context timeout per scrape.
