package main

// Two tests + a benchmark.
//   go test ./sisu-interview-prep/task-08-rates-cache-refresh-solution -v
//   go test ./sisu-interview-prep/task-08-rates-cache-refresh-solution -bench=. -benchmem

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestConvertHandler_StaleFailFast(t *testing.T) {
	cache := &RateCache{
		rates:     map[string]float64{"EURUSD": 2},
		updatedAt: time.Now().Add(-10 * time.Second),
	}

	h := convertHandler(cache, 2*time.Second)
	req := httptest.NewRequest(http.MethodGet, "/convert?from=EUR&to=USD&amount=10", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status=%d, want 503, body=%s", rec.Code, rec.Body.String())
	}
}

func TestConvertHandler_AllowStale(t *testing.T) {
	cache := &RateCache{
		rates:     map[string]float64{"EURUSD": 2},
		updatedAt: time.Now().Add(-10 * time.Second),
	}

	h := convertHandler(cache, 2*time.Second)
	req := httptest.NewRequest(http.MethodGet, "/convert?from=EUR&to=USD&amount=10&allow_stale=true", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d, want 200, body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), `"converted"`) {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}

func BenchmarkConvertHandler(b *testing.B) {
	cache := &RateCache{
		rates:     map[string]float64{"EURUSD": 1.1},
		updatedAt: time.Now(),
	}
	h := convertHandler(cache, 2*time.Second)

	req := httptest.NewRequest(http.MethodGet, "/convert?from=EUR&to=USD&amount=10&allow_stale=true", nil)

	b.ReportAllocs()
	for b.Loop() {
		rec := httptest.NewRecorder()
		h(rec, req)
		if rec.Code != http.StatusOK {
			b.Fatal(rec.Body.String())
		}
	}
}
