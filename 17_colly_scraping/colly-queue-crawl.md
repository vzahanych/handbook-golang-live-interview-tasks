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
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
