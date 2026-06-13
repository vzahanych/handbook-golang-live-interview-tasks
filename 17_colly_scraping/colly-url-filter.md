# colly url filter

## Live interview task
Allow only selected URLs by matching path patterns before visiting.

## Concepts covered
- Colly
- URLFilters
- regexp

## Candidate solution

```go
package main

import (
    "fmt"
    "regexp"
    "github.com/gocolly/colly/v2"
)

func main() {
    allowed := regexp.MustCompile(`/docs/|/$`)
    c := colly.NewCollector(colly.AllowedDomains("go.dev"), colly.URLFilters(allowed))
    c.OnRequest(func(r *colly.Request) { fmt.Println("visit", r.URL) })
    c.Visit("https://go.dev/")
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- `URLFilters` regex matches full URL string — anchor patterns carefully (`$` end of path).
- Combine with `AllowedDomains` — filter alone won't block evil.com.
- `OnRequest` can `r.Abort()` for dynamic rules regex can't express.

## Q&A

**Q: Filter vs OnRequest abort?**  
A: URLFilters early reject; Abort for per-request logic (headers, depth).

**Q: Complexity?**  
A: O(regex) per candidate URL — compile once with `MustCompile`.

**Q: Edge cases?**  
A: Trailing slash, URL encoding `%20`, case sensitivity.

**Q: Production?**  
A: Allowlist sitemap URLs first, then expand with filters.

**Q: Interview trap?**  
A: Forgetting scheme/host in regex — match `r.URL.Path` in OnRequest instead.
