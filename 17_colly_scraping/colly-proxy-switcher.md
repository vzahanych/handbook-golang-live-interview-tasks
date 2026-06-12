# colly proxy switcher

## Live interview task
Configure rotating proxies for a collector.

## Concepts covered
- Colly
- proxy switcher

## Candidate solution

```go
package main

import (
    "log"
    "github.com/gocolly/colly/v2"
    "github.com/gocolly/colly/v2/proxy"
)

func main() {
    c := colly.NewCollector()
    rp, err := proxy.RoundRobinProxySwitcher(
        "http://proxy1.example:8080",
        "http://proxy2.example:8080",
    )
    if err != nil { log.Fatal(err) }
    c.SetProxyFunc(rp)
    c.OnError(func(r *colly.Response, err error) { log.Println("proxy/request failed", err) })
    _ = c.Visit("https://example.com")
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
