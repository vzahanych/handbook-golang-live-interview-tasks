package main

// 1. Run `go run -race .` and hit /convert in parallel (see curl below). There is
//    a data race on the `rates` map. Fix it with sync.RWMutex (or another safe
//    structure), keeping reads cheap.
//    Reads dominate writes for a cache, so an RWMutex is a good fit: many
//    concurrent readers (RLock) and an occasional writer (Lock) on refresh.
// 2. The refresh goroutine uses time.Tick, which cannot be stopped, so it leaks
//    if you ever want to shut down or restart the component. Replace it with a
//    time.Ticker that you stop, and add cancellation via context.
//    Use time.NewTicker and `defer ticker.Stop()`. In the loop, select on
//    ctx.Done() to exit cleanly.
// 3. The refresh does not use a timeout. If the upstream rate source stalls, the
//    goroutine blocks forever and the cache stops updating. Add a per-refresh
//    deadline.
//    Wrap each fetch in context.WithTimeout so a slow source doesn't block the
//    refresh loop indefinitely.
// 4. Ensure the refresh does not stampede the upstream if /convert is called
//    frequently. Reads must never trigger a refresh; only the periodic loop does.
//    Keep refresh entirely in the background loop; handlers only read the cache.
// 5. Add a small "staleness" signal: if rates are older than X seconds, /convert
//    should return 503 (fail-fast) OR continue with stale data (degrade
//    gracefully) based on a query param like ?allow_stale=true.
//    Implement both behaviors behind the allow_stale switch.

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// A typical CIL back-end keeps small, frequently-used reference data in memory
// (e.g. currency rates, product mappings, translations). The refresh loop must be
// safe and must not overload the source system.
//
// This starter implements an in-memory FX rates cache, but it has classic
// concurrency and lifecycle problems.
//
// 1. Run `go run -race .` and hit /convert in parallel (see curl below). There is
//    a data race on the `rates` map. Fix it with sync.RWMutex (or another safe
//    structure), keeping reads cheap.
//
// 2. The refresh goroutine uses time.Tick, which cannot be stopped, so it leaks
//    if you ever want to shut down or restart the component. Replace it with a
//    time.Ticker that you stop, and add cancellation via context.
//
// 3. The refresh does not use a timeout. If the upstream rate source stalls, the
//    goroutine blocks forever and the cache stops updating. Add a per-refresh
//    deadline.
//
// 4. Ensure the refresh does not stampede the upstream if /convert is called
//    frequently. Reads must never trigger a refresh; only the periodic loop does.
//
// 5. Add a small "staleness" signal: if rates are older than X seconds, /convert
//    should return 503 (fail-fast) OR continue with stale data (degrade
//    gracefully) based on a query param like ?allow_stale=true.
//
// Run hints:
// - Start: `go run .`
// - Stress: `for i in {1..50}; do curl -sS "localhost:8080/convert?from=EUR&to=USD&amount=10" & done; wait`

type RateCache struct {
	rates     map[string]float64
	updatedAt time.Time

	mu sync.RWMutex // answer 1: protect rates+updatedAt
}

func NewRateCache() *RateCache {
	return &RateCache{
		rates: map[string]float64{
			"EURUSD": 1.10,
			"USDEUR": 0.91,
		},
		updatedAt: time.Now(),
	}
}

func (c *RateCache) Get(pair string) (float64, time.Time, bool) {
	// BUG: concurrent read/write race.
	c.mu.RLock()
	defer c.mu.RUnlock()
	r, ok := c.rates[pair]
	return r, c.updatedAt, ok
}

func (c *RateCache) ReplaceAll(newRates map[string]float64) {
	// BUG: concurrent read/write race; also replaces map without synchronization.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.rates = newRates
	c.updatedAt = time.Now()
}

func fetchRates(ctx context.Context, url string) (map[string]float64, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }() // answer 3: always close

	// var out map[string]float64
	// if err := json.NewDecoder(resp.Body).Decode(&out); err != nil { // old: materializes whole map at once
	// 	return nil, err
	// }
	// return out, nil
	dec := json.NewDecoder(resp.Body)
	tok, err := dec.Token()
	if err != nil {
		return nil, err
	}
	if d, ok := tok.(json.Delim); !ok || d != '{' {
		return nil, fmt.Errorf("expected JSON object")
	}
	rates := make(map[string]float64)
	for dec.More() {
		keyTok, err := dec.Token()
		if err != nil {
			return nil, err
		}
		key, ok := keyTok.(string)
		if !ok {
			return nil, fmt.Errorf("expected object key")
		}
		var rate float64
		if err := dec.Decode(&rate); err != nil {
			return nil, err
		}
		rates[key] = rate
	}
	if _, err := dec.Token(); err != nil { // closing '}'
		return nil, err
	}
	return rates, nil
}

func startRefreshLoop(ctx context.Context, cache *RateCache, srcURL string, interval time.Duration, perRefreshTimeout time.Duration) {
	// BUG: time.Tick cannot be stopped; loop ignores ctx cancellation.
	ticker := time.NewTicker(interval) // answer 2
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			refreshCtx, cancel := context.WithTimeout(ctx, perRefreshTimeout) // answer 3
			newRates, err := fetchRates(refreshCtx, srcURL)
			cancel()
			if err != nil {
				log.Printf("refresh error: %v", err)
				continue
			}
			cache.ReplaceAll(newRates)
		}
	}
}

func convertHandler(cache *RateCache, staleAfter time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		amountStr := r.URL.Query().Get("amount")
		allowStale := r.URL.Query().Get("allow_stale") == "true"

		var amount float64
		if _, err := fmt.Sscanf(amountStr, "%f", &amount); err != nil {
			http.Error(w, "bad amount", http.StatusBadRequest)
			return
		}
		pair := from + to
		rate, updatedAt, ok := cache.Get(pair)
		if !ok {
			http.Error(w, "unknown pair", http.StatusBadRequest)
			return
		}

		age := time.Since(updatedAt)
		if age > staleAfter && !allowStale {
			http.Error(w, "rates stale", http.StatusServiceUnavailable)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"from":       from,
			"to":         to,
			"amount":     amount,
			"rate":       rate,
			"converted":  amount * rate,
			"updated_at": updatedAt.Format(time.RFC3339Nano),
			"age_ms":     age.Milliseconds(),
		})
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Fake upstream source of rates (sometimes slow).
	http.HandleFunc("/rates", func(w http.ResponseWriter, r *http.Request) {
		// Occasionally slow to simulate upstream latency spikes.
		if rand.Intn(10) == 0 {
			time.Sleep(2 * time.Second)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]float64{
			"EURUSD": 1.05 + rand.Float64()*0.2,
			"USDEUR": 0.90 + rand.Float64()*0.1,
		})
	})

	cache := NewRateCache()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go startRefreshLoop(ctx, cache, "http://localhost:8080/rates", 500*time.Millisecond, 300*time.Millisecond)

	http.HandleFunc("/convert", convertHandler(cache, 2*time.Second))

	srv := &http.Server{Addr: ":8080"} // nil Handler => DefaultServeMux (prompt rule)
	log.Printf("listening on %s", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
//   # fresh (or stale allowed):
//   curl -sS "localhost:8080/convert?from=EUR&to=USD&amount=10" | jq
//   curl -sS "localhost:8080/convert?from=EUR&to=USD&amount=10&allow_stale=true" | jq
//
//   # stress reads:
//   for i in {1..50}; do curl -sS "localhost:8080/convert?from=EUR&to=USD&amount=10" >/dev/null & done; wait
//
//   go test ./sisu-interview-prep/task-08-rates-cache-refresh-solution -v
