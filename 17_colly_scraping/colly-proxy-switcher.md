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
- `RoundRobinProxySwitcher` rotates per request — failed proxy still rotates (add health checks).
- HTTPS needs CONNECT support from proxy — HTTP proxy URL is typical.
- Respect target site ToS — proxies for evasion are unethical/illegal in many contexts.

## Q&A

**Q: Sticky sessions?**  
A: Custom `SetProxyFunc` mapping domain → fixed proxy.

**Q: Complexity?**  
A: O(1) proxy pick per request.

**Q: Edge cases?**  
A: Proxy auth in URL `http://user:pass@host:8080`, TLS MITM corporate proxies.

**Q: OnError handling?**  
A: Log and retry with different proxy or direct fallback.

**Q: Production?**  
A: Proxy pool health metrics, rate limit per egress IP.
