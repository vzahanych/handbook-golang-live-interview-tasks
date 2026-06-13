# colly queue crawl

## Live interview task
Use Colly queue to manage URLs and consumer parallelism.

## Concepts covered
- Colly queue
- InMemoryQueueStorage

## Candidate solution

```go
package main

import (
    "log"
    "github.com/gocolly/colly/v2"
    "github.com/gocolly/colly/v2/queue"
)

func main() {
    c := colly.NewCollector()
    c.OnResponse(func(r *colly.Response) { log.Println(r.Request.URL, r.StatusCode) })
    q, _ := queue.New(2, &queue.InMemoryQueueStorage{MaxSize: 1000})
    q.AddURL("https://example.com")
    q.AddURL("https://example.com/about")
    q.Run(c)
}
```

## Run

```bash
go mod init scrape && go get github.com/gocolly/colly/v2 && go run .
```

## Interview notes / pitfalls
- First arg to `queue.New` is consumer parallelism — not the same as `LimitRule.Parallelism`.
- `InMemoryQueueStorage` loses queue on crash — use Redis/DB storage for resume.
- `q.Run(c)` blocks until queue drained — add URLs dynamically from OnHTML callbacks.

## Q&A

**Q: Queue vs direct Visit?**  
A: Queue serializes scheduling, supports pause/resume and persistent storage backends.

**Q: Complexity?**  
A: O(urls) with W parallel consumers.

**Q: MaxSize 1000?**  
A: Backpressure — producer must handle full queue.

**Q: Edge cases?**  
A: Duplicate AddURL — dedupe with visited set before enqueue.

**Q: Production?**  
A: Redis queue storage, dead-letter for permanent failures.
