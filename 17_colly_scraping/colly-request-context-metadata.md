# colly request context metadata

## Live interview task
Attach metadata to requests and read it in callbacks.

## Concepts covered
- Colly
- request context

## Candidate solution

```go
package main

import (
    "fmt"
    "github.com/gocolly/colly/v2"
)

func main() {
    c := colly.NewCollector()
    c.OnResponse(func(r *colly.Response) { fmt.Println("source", r.Ctx.Get("source"), "url", r.Request.URL) })
    ctx := colly.NewContext()
    ctx.Put("source", "seed")
    c.Request("GET", "https://example.com", nil, ctx, nil)
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- `colly.Context` is string key-value per request — propagate via `e.Request.Visit` with `r.Ctx`.
- Prefer `Request("GET", ...)` when seeding with metadata; `Visit` creates fresh context.
- For typed data use JSON in context or attach via `r.Context` (stdlib context) in newer patterns.

## Q&A

**Q: Pass depth/parent URL?**  
A: `ctx.Put("parent", r.URL.String())` before child `Visit`.

**Q: Complexity?**  
A: O(1) map lookup per callback.

**Q: Edge cases?**  
A: Context not copied if you call `Visit` without inheriting — use `e.Request.Visit`.

**Q: vs closure variables?**  
A: Context survives async callbacks cleaner than loop capture bugs.

**Q: Production?**  
A: Trace ID in context for structured logs across redirect chain.
