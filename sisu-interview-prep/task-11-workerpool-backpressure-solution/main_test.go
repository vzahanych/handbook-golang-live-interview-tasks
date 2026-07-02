package main

// Two tests + a benchmark.
//   go test ./sisu-interview-prep/task-11-workerpool-backpressure-solution -v
//   go test ./sisu-interview-prep/task-11-workerpool-backpressure-solution -bench=. -benchmem

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIngestAndProcess_DropModeDrops(t *testing.T) {
	ctx := context.Background()
	// tiny queue + slow processing => drops
	got, err := ingestAndProcess(ctx, 10_000, 1, 1, "drop")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Dropped == 0 {
		t.Fatalf("expected some drops, got %+v", got)
	}
}

func TestHandler_FailFastReturns429(t *testing.T) {
	// Build a minimal handler using the same logic as main.
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 50*time.Millisecond)
		defer cancel()

		out, err := ingestAndProcess(ctx, 1_000_000, 1, 1, "fail_fast")
		_ = out
		if err != nil && err.Error() == "queue full" {
			http.Error(w, err.Error(), http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/ingest", strings.NewReader(""))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("status=%d want 429", rec.Code)
	}
}

func BenchmarkIngestAndProcess(b *testing.B) {
	ctx := context.Background()
	b.ReportAllocs()
	for b.Loop() {
		_, _ = ingestAndProcess(ctx, 5000, 8, 256, "block")
	}
}
