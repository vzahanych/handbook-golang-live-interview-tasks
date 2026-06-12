# in memory url shortener http

## Live interview task
Build a tiny in-memory URL shortener HTTP service.

## Concepts covered
- HTTP
- JSON
- mutex
- maps

## Candidate solution

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
)

type Store struct { mu sync.RWMutex; next int; data map[string]string }
func NewStore() *Store { return &Store{data: make(map[string]string)} }
func (s *Store) Put(url string) string { s.mu.Lock(); defer s.mu.Unlock(); s.next++; id := fmt.Sprintf("%x", s.next); s.data[id] = url; return id }
func (s *Store) Get(id string) (string, bool) { s.mu.RLock(); defer s.mu.RUnlock(); v, ok := s.data[id]; return v, ok }

func main() {
    st := NewStore()
    http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request){ var req struct{ URL string `json:"url"` }; json.NewDecoder(r.Body).Decode(&req); json.NewEncoder(w).Encode(map[string]string{"id": st.Put(req.URL)}) })
    http.HandleFunc("/r/", func(w http.ResponseWriter, r *http.Request){ id := r.URL.Path[len("/r/"):]; if url, ok := st.Get(id); ok { http.Redirect(w,r,url,302); return }; http.NotFound(w,r) })
    http.ListenAndServe(":8080", nil)
}
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
