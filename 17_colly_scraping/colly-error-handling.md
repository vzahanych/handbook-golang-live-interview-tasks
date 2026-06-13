# colly error handling

## Live interview task
Handle request errors using OnError and log the failed URL and status code.

## Concepts covered
- Colly
- OnError

## Candidate solution

```go
package main

import (
    "log"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()
    c.OnError(func(r *colly.Response, err error) {
        status := 0
        if r != nil { status = r.StatusCode }
        log.Printf("failed url=%v status=%d err=%v", r.Request.URL, status, err)
    })
    _ = c.Visit("https://example.com/not-found")
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- `OnError` fires for network errors and HTTP errors depending on Colly version/settings — check `r != nil` before `StatusCode`.
- Distinguish retryable (5xx, timeout) vs permanent (404) in handler logic.
- Log `r.Request.URL` not bare `r` — Response may be partial on dial failure.

## Q&A

**Q: Retry failed requests?**  
A: `r.Request.Retry()` in OnError with max retry counter in context.

**Q: Complexity?**  
A: O(1) per failed request callback.

**Q: Edge cases?**  
A: TLS errors, DNS failure, context cancel mid-request.

**Q: vs OnResponse?**  
A: OnResponse only on success; OnError for failures.

**Q: Production?**  
A: Metrics per status class, circuit breaker on repeated 503.
