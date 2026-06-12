# http client timeout and status check

## Live interview task
Create an HTTP client with timeout and verify status codes.

## Concepts covered
- http.Client
- timeouts
- io.ReadAll

## Candidate solution

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

func fetch(url string) ([]byte, error) {
    client := &http.Client{Timeout: 3 * time.Second}
    resp, err := client.Get(url)
    if err != nil { return nil, err }
    defer resp.Body.Close()
    if resp.StatusCode < 200 || resp.StatusCode >= 300 { return nil, fmt.Errorf("bad status: %s", resp.Status) }
    return io.ReadAll(resp.Body)
}

func main() { b, err := fetch("https://example.com"); fmt.Println(len(b), err) }
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
