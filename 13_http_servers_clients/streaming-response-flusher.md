# streaming response flusher

## Live interview task
Stream lines to the client with `http.Flusher`.

## Concepts covered
- streaming
- http.Flusher
- chunked transfer

## Candidate solution

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func stream(w http.ResponseWriter, r *http.Request) {
    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "streaming unsupported", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/plain")
    w.Header().Set("X-Content-Type-Options", "nosniff")

    for i := 0; i < 5; i++ {
        fmt.Fprintf(w, "data: %d\n", i)
        flusher.Flush()
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    http.HandleFunc("/stream", stream)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Run

```bash
go run .
curl -N localhost:8080/stream
```

## Interview notes / pitfalls
- `Flush` sends buffered data immediately — needed for SSE/long polling feel.
- Wrapping `ResponseWriter` in middleware may hide `Flusher` — use wrapper that implements Flusher.
- HTTP/1.1 chunked encoding used automatically when no Content-Length.
- Client needs `-N` (curl) to disable buffering.

## Q&A

**Q: SSE?**  
A: `Content-Type: text/event-stream`, `data: ...\n\n`, flush each event.

**Q: Context cancel?**  
A: Check `r.Context().Done()` in loop — stop streaming.

**Q: gzip middleware?**  
A: May buffer — flush less effective until buffer fills.

**Q: HTTP/2?**  
A: Flusher still works per response stream.

**Q: Use case?**  
A: Log tail, progress updates, LLM token stream.
