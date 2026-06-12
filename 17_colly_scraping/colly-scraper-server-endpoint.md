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
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
