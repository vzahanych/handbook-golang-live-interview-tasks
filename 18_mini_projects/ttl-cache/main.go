// Command ttl-cache is a generic in-memory cache with per-entry TTL. It shows
// the two ways expired entries leave the map:
//
//   - lazy expiry: Get treats an expired entry as absent (but leaves it in place)
//   - active cleanup: an optional background goroutine periodically deletes
//     expired entries so keys that are never read again don't leak memory
//
// The main function below demonstrates both, plus the Delete/Len helpers.
package main

import (
	"fmt"
	"sync"
	"time"
)

// item bundles a value with its absolute expiry time. Storing the deadline
// (rather than the TTL duration) means Get/cleanup just compare against now.
type item[V any] struct {
	value   V
	expires time.Time
}

// Cache is safe for concurrent use. K must be comparable (it's a map key); V is
// unconstrained. The RWMutex lets many Gets run in parallel while Set/Delete and
// the cleanup sweep take the exclusive write lock.
type Cache[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]item[V]
}

// New returns an empty cache ready to use.
func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{data: make(map[K]item[V])}
}

// Set stores v under k, expiring ttl from now. A repeated key overwrites both
// the value and its deadline.
func (c *Cache[K, V]) Set(k K, v V, ttl time.Duration) {
	c.mu.Lock()
	c.data[k] = item[V]{value: v, expires: time.Now().Add(ttl)}
	c.mu.Unlock()
}

// Get returns the value and true only if the key exists AND has not expired.
// This is the "lazy expiry" path: an expired entry is reported as absent here,
// but is not deleted until the cleanup sweep runs (or the key is overwritten).
func (c *Cache[K, V]) Get(k K) (V, bool) {
	c.mu.RLock()
	it, ok := c.data[k]
	c.mu.RUnlock()

	var zero V
	if !ok || time.Now().After(it.expires) {
		return zero, false
	}
	return it.value, true
}

// Delete removes a key immediately, whether or not it has expired.
func (c *Cache[K, V]) Delete(k K) {
	c.mu.Lock()
	delete(c.data, k)
	c.mu.Unlock()
}

// Len reports how many entries are currently stored. Note: this counts
// not-yet-swept expired entries too, since lazy expiry leaves them in the map.
func (c *Cache[K, V]) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.data)
}

// StartCleanup launches the background sweeper and returns a stop function.
// Calling stop() ends the goroutine; it's idempotent-safe to defer. Keeping the
// stop channel here (instead of a package global) makes the cache self-contained
// and the goroutine's lifetime explicit — no leaked tickers.
func (c *Cache[K, V]) StartCleanup(interval time.Duration) (stop func()) {
	done := make(chan struct{})
	go c.cleanupLoop(interval, done)

	var once sync.Once
	return func() { once.Do(func() { close(done) }) }
}

// cleanupLoop deletes expired entries on every tick until stop is closed.
func (c *Cache[K, V]) cleanupLoop(interval time.Duration, stop <-chan struct{}) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop() // releases the ticker's resources when we return
	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for k, it := range c.data {
				if now.After(it.expires) {
					delete(c.data, k) // deleting during range is allowed in Go
				}
			}
			c.mu.Unlock()
		}
	}
}

func main() {
	c := New[string, int]()

	// 1) Basic set/get within the TTL window.
	c.Set("a", 1, 50*time.Millisecond)
	v, ok := c.Get("a")
	fmt.Printf("get a immediately -> value=%d ok=%t\n", v, ok) // 1, true

	// 2) Lazy expiry: after the TTL passes, Get reports a miss even though the
	//    entry is still physically in the map (so Len is still 1).
	time.Sleep(80 * time.Millisecond)
	_, ok = c.Get("a")
	fmt.Printf("get a after expiry -> ok=%t (lazy: still in map, Len=%d)\n", ok, c.Len())

	// 3) Active cleanup: run the sweeper on a short interval and watch Len drop
	//    to 0 as the expired entry is reclaimed without anyone reading it.
	stop := c.StartCleanup(20 * time.Millisecond)
	defer stop()
	time.Sleep(40 * time.Millisecond)
	fmt.Printf("after cleanup tick -> Len=%d\n", c.Len()) // 0

	// 4) Delete is immediate, regardless of TTL.
	c.Set("b", 2, time.Hour)
	c.Delete("b")
	_, ok = c.Get("b")
	fmt.Printf("get b after delete -> ok=%t\n", ok) // false
}
