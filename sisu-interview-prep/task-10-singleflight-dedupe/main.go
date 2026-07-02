package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

// In the CIL, a single user action can cause multiple requests for the same
// reference data (partner config, product mapping, policy terms). If 1000
// requests arrive at once and the cache misses, you get a thundering herd: 1000
// identical downstream calls.
//
// This starter has an in-memory cache but still stampedes the backend because
// it fetches on every miss without deduping in-flight fetches.
//
// 1. Add request-scoped context handling (deadline + cancellation propagation)
//    from /get to the downstream call.
//
// 2. Add "singleflight" style dedupe so only ONE downstream fetch is in-flight
//    per key at a time; concurrent callers wait for the same result.
//    (Implement it yourself with a map+mutex+channels; no need for extra deps.)
//
// 3. Ensure errors do not get cached permanently. Decide whether to cache errors
//    briefly (negative caching) and explain the trade-off.
//
// 4. Add a max concurrency bound for distinct keys (so many different keys don't
//    spawn unlimited goroutines).
//
// Run hints:
// - Start: `go run .`
// - Single key stampede:
//   `for i in {1..50}; do curl -sS "localhost:8080/get?key=foo" & done; wait | wc -l`
// - Observe backend calls in logs and the /stats endpoint.

type cacheEntry struct {
	value     string
	expiresAt time.Time
}

type cache struct {
	ttl   time.Duration
	items map[string]cacheEntry
	// BUG: missing lock
}

func newCache(ttl time.Duration) *cache {
	return &cache{ttl: ttl, items: make(map[string]cacheEntry)}
}

func (c *cache) get(key string) (string, bool) {
	ent, ok := c.items[key]
	if !ok {
		return "", false
	}
	if time.Now().After(ent.expiresAt) {
		return "", false
	}
	return ent.value, true
}

func (c *cache) set(key, value string) {
	c.items[key] = cacheEntry{value: value, expiresAt: time.Now().Add(c.ttl)}
}

var backendCalls int64

func backendFetch(ctx context.Context, key string) (string, error) {
	atomic.AddInt64(&backendCalls, 1)
	// Random latency spike.
	time.Sleep(time.Duration(50+rand.Intn(200)) * time.Millisecond)
	// Random failure.
	if rand.Intn(20) == 0 {
		return "", fmt.Errorf("backend error for %q", key)
	}
	return "value-for-" + key, nil
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	c := newCache(2 * time.Second)

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key", http.StatusBadRequest)
			return
		}

		if v, ok := c.get(key); ok {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"key":    key,
				"value":  v,
				"cached": true,
			})
			return
		}

		// BUG: no context deadline; no dedupe; multiple goroutines will all fetch.
		v, err := backendFetch(context.Background(), key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		c.set(key, v)

		_ = json.NewEncoder(w).Encode(map[string]any{
			"key":    key,
			"value":  v,
			"cached": false,
		})
	})

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"backend_calls": atomic.LoadInt64(&backendCalls),
			"now_unix_ms":   strconv.FormatInt(time.Now().UnixMilli(), 10),
		})
	})

	log.Printf("listening on :8080")
	_ = http.ListenAndServe(":8080", nil)
}
