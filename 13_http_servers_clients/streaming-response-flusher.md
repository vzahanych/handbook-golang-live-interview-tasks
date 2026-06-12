# streaming response flusher

## Live interview task
Stream lines to the client with http.Flusher.

## Concepts covered
- streaming
- type assertion
- http.Flusher

## Candidate solution

```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func stream(w http.ResponseWriter, r *http.Request) {
    flusher, ok := w.(http.Flusher)
    if !ok { http.Error(w, "streaming unsupported", 500); return }
    for i := 0; i < 5; i++ {
        fmt.Fprintf(w, "data: %d\n", i)
        flusher.Flush()
        time.Sleep(500 * time.Millisecond)
    }
}

func main() { http.HandleFunc("/stream", stream); http.ListenAndServe(":8080", nil) }
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
