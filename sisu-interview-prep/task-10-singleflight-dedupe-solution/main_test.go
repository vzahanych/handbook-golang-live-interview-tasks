package main

// Two tests + a benchmark.
//   go test ./sisu-interview-prep/task-10-singleflight-dedupe-solution -v
//   go test ./sisu-interview-prep/task-10-singleflight-dedupe-solution -bench=. -benchmem

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSingleflight_DedupesBackendFetches(t *testing.T) {
	s := newService()
	var calls int64
	s.fetchFn = func(ctx context.Context, key string) (string, error) {
		atomic.AddInt64(&calls, 1)
		time.Sleep(10 * time.Millisecond)
		return "v", nil
	}

	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			if _, _, err := s.getValue(ctx, "k"); err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		}()
	}
	wg.Wait()

	if atomic.LoadInt64(&calls) != 1 {
		t.Fatalf("backend calls=%d, want 1", calls)
	}
}

func TestGetHandler_OK(t *testing.T) {
	s := newService()
	s.fetchFn = func(ctx context.Context, key string) (string, error) { return "value", nil }
	h := getHandler(s)

	req := httptest.NewRequest(http.MethodGet, "/get?key=foo", nil)
	rec := httptest.NewRecorder()
	h(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), `"value"`) {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}

func BenchmarkGetValue(b *testing.B) {
	s := newService()
	s.fetchFn = func(ctx context.Context, key string) (string, error) { return "value", nil }

	b.ReportAllocs()
	for b.Loop() {
		_, _, _ = s.getValue(context.Background(), "k")
		// clear cache to keep it measuring the path with singleflight+fetch
		s.cache = newCache(0)
	}
}
