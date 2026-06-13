# in memory url shortener http

## Live interview task
Build a tiny in-memory URL shortener HTTP service (POST shorten, GET redirect).

## Concepts covered
- HTTP handlers
- JSON API
- RWMutex
- maps

## Candidate solution

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"
)

type Store struct {
    mu   sync.RWMutex
    next int
    data map[string]string
}

func NewStore() *Store {
    return &Store{data: make(map[string]string)}
}

func (s *Store) Put(url string) string {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.next++
    id := fmt.Sprintf("%x", s.next)
    s.data[id] = url
    return id
}

func (s *Store) Get(id string) (string, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    v, ok := s.data[id]
    return v, ok
}

func main() {
    st := NewStore()
    mux := http.NewServeMux()

    mux.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
        var req struct {
            URL string `json:"url"`
        }
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        json.NewEncoder(w).Encode(map[string]string{"id": st.Put(req.URL)})
    })

    mux.HandleFunc("GET /r/{id}", func(w http.ResponseWriter, r *http.Request) {
        id := r.PathValue("id")
        url, ok := st.Get(id)
        if !ok {
            http.NotFound(w, r)
            return
        }
        http.Redirect(w, r, url, http.StatusFound)
    })

    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

## Run

```bash
go run .
curl -X POST localhost:8080/shorten -d '{"url":"https://go.dev"}' -H 'Content-Type: application/json'
curl -i localhost:8080/r/1
```

## Interview notes / pitfalls
- `RWMutex` — many reads (redirects), fewer writes (shorten).
- ID generation: counter hex demo — production uses base62 random + collision check.
- Validate URL scheme (`https://`) — reject javascript: URLs.
- Data lost on restart — persist to DB/Redis for production.

## Q&A

**Q: Complexity?**  
A: Get/Put O(1) average map ops.

**Q: Race on map?**  
A: Mutex prevents concurrent map panic.

**Q: 302 vs 301?**  
A: 302 temporary — 301 permanent redirect caching.

**Q: Extend?**  
A: Expiry TTL, click analytics, custom aliases.

**Q: Test?**  
A: httptest POST + GET redirect chain.
