package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
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

	// BUG: no lock yet; this map is read/written concurrently.
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
	r, ok := c.rates[pair]
	return r, c.updatedAt, ok
}

func (c *RateCache) ReplaceAll(newRates map[string]float64) {
	// BUG: concurrent read/write race; also replaces map without synchronization.
	c.rates = newRates
	c.updatedAt = time.Now()
}

func fetchRates(ctx context.Context, url string) (map[string]float64, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	// BUG: resp.Body not closed.

	var out map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return out, nil
}

func startRefreshLoop(ctx context.Context, cache *RateCache, srcURL string) {
	// BUG: time.Tick cannot be stopped; loop ignores ctx cancellation.
	for range time.Tick(500 * time.Millisecond) {
		newRates, err := fetchRates(ctx, srcURL)
		if err != nil {
			log.Printf("refresh error: %v", err)
			continue
		}
		cache.ReplaceAll(newRates)
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
	ctx := context.Background()
	go startRefreshLoop(ctx, cache, "http://localhost:8080/rates")

	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
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
		if age > 2*time.Second && !allowStale {
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
	})

	log.Printf("listening on :8080")
	_ = http.ListenAndServe(":8080", nil)
}
