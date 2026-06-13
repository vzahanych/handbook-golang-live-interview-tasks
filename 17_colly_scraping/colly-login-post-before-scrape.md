# colly login post before scrape

## Live interview task
Authenticate with POST before scraping authenticated pages.

## Concepts covered
- Colly
- Post
- cookies/session

## Candidate solution

```go
package main

import (
    "log"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()
    if err := c.Post("https://example.com/login", map[string]string{"username":"admin", "password":"admin"}); err != nil {
        log.Println("login failed:", err)
    }
    c.OnResponse(func(r *colly.Response) { log.Println("status", r.StatusCode) })
    _ = c.Visit("https://example.com/account")
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- Colly jar persists cookies across requests on same collector — reuse collector after login.
- Real sites need CSRF tokens, hidden fields — scrape login form first with OnHTML.
- `Post` with map sends `application/x-www-form-urlencoded`, not JSON.

## Q&A

**Q: Session stickiness?**  
A: Same `Collector` instance shares cookie jar automatically.

**Q: OAuth / JWT?**  
A: Set `Authorization` header via `c.OnRequest` after token exchange.

**Q: Complexity?**  
A: Two requests minimum — login + protected page.

**Q: Edge cases?**  
A: 302 redirect after login, Set-Cookie Secure flag, expired session mid-crawl.

**Q: Production?**  
A: Never hardcode passwords — env vars, rotate sessions, detect login failure HTML.
