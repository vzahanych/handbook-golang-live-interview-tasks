package main

// 1. Add request-scoped context handling (deadline + cancellation propagation)
//    from /get to the downstream call.
//    Use r.Context() and wrap it with a timeout so slow backends cancel promptly.
// 2. Add "singleflight" style dedupe so only ONE downstream fetch is in-flight
//    per key at a time; concurrent callers wait for the same result.
//    Implement a map[key]*call protected by a mutex; callers either become the
//    leader (do the fetch) or wait on a done channel.
// 3. Ensure errors do not get cached permanently. Decide whether to cache errors
//    briefly (negative caching) and explain the trade-off.
//    Permanent error caching is dangerous (one transient outage poisons the key).
//    Optional short negative caching can reduce load during a backend outage, but
//    increases the chance of serving stale failures. We'll implement optional
//    negative caching with a short TTL behind a query param.
// 4. Add a max concurrency bound for distinct keys (so many different keys don't
//    spawn unlimited goroutines).
//    Use a semaphore that is acquired for the duration of the backend fetch.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
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
	mu    sync.RWMutex // answer: protect items
}

func newCache(ttl time.Duration) *cache {
	return &cache{ttl: ttl, items: make(map[string]cacheEntry)}
}

func (c *cache) get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
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
	c.mu.Lock()
	defer c.mu.Unlock()
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

type inFlight struct {
	done chan struct{}
	val  string
	err  error
}

type singleflightGroup struct {
	mu sync.Mutex
	m  map[string]*inFlight
}

func newSingleflightGroup() *singleflightGroup {
	return &singleflightGroup{m: make(map[string]*inFlight)}
}

func (g *singleflightGroup) do(ctx context.Context, key string, fn func(context.Context, string) (string, error)) (string, error) {
	g.mu.Lock()
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-c.done:
			return c.val, c.err
		}
	}
	c := &inFlight{done: make(chan struct{})}
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn(ctx, key)

	g.mu.Lock()
	delete(g.m, key)
	close(c.done)
	g.mu.Unlock()

	return c.val, c.err
}

type service struct {
	cache          *cache
	negCache       *cache
	sf             *singleflightGroup
	maxFetchSem    chan struct{}
	perRequestTTL  time.Duration
	negCacheEnable bool
	fetchFn        func(context.Context, string) (string, error)
}

func newService() *service {
	return &service{
		cache:         newCache(2 * time.Second),
		negCache:      newCache(250 * time.Millisecond),
		sf:            newSingleflightGroup(),
		maxFetchSem:   make(chan struct{}, 20),
		perRequestTTL: 300 * time.Millisecond,
		fetchFn:       backendFetch,
	}
}

func (s *service) getValue(ctx context.Context, key string) (string, bool, error) {
	if v, ok := s.cache.get(key); ok {
		return v, true, nil
	}
	if s.negCacheEnable {
		if msg, ok := s.negCache.get("err:" + key); ok {
			return "", false, errors.New(msg)
		}
	}

	v, err := s.sf.do(ctx, key, func(ctx context.Context, key string) (string, error) {
		// Concurrency bound for distinct keys (leaders).
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case s.maxFetchSem <- struct{}{}:
		}
		defer func() { <-s.maxFetchSem }()

		return s.fetchFn(ctx, key)
	})
	if err != nil {
		if s.negCacheEnable {
			s.negCache.set("err:"+key, err.Error())
		}
		return "", false, err
	}
	s.cache.set(key, v)
	return v, false, nil
}

func getHandler(s *service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing key", http.StatusBadRequest)
			return
		}

		s.negCacheEnable = r.URL.Query().Get("negative_cache") == "true"

		ctx := r.Context()
		ctx, cancel := context.WithTimeout(ctx, s.perRequestTTL) // answer 1: request-scoped timeout
		defer cancel()

		v, cached, err := s.getValue(ctx, key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		_ = json.NewEncoder(w).Encode(map[string]any{
			"key":    key,
			"value":  v,
			"cached": cached,
		})
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	s := newService()
	http.HandleFunc("/get", getHandler(s))

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"backend_calls": atomic.LoadInt64(&backendCalls),
			"now_unix_ms":   strconv.FormatInt(time.Now().UnixMilli(), 10),
		})
	})

	srv := &http.Server{Addr: ":8080"} // nil handler => DefaultServeMux
	log.Printf("listening on %s", srv.Addr)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
		}
	}()
	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownCtx)
}

// How to run and test:
//   go run .
//
//   # stampede on one key (watch /stats):
//   for i in {1..50}; do curl -sS "localhost:8080/get?key=foo" >/dev/null & done; wait
//   curl -sS "localhost:8080/stats" | jq
//
//   # optional negative caching:
//   curl -sS "localhost:8080/get?key=foo&negative_cache=true" | jq
//
//   go test ./sisu-interview-prep/task-10-singleflight-dedupe-solution -v
