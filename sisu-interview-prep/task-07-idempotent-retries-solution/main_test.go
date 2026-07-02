package main

// Two tests + a benchmark, same package as main.go so we can call unexported
// functions.
//   go test ./sisu-interview-prep/task-07-idempotent-retries-solution -v
//   go test ./sisu-interview-prep/task-07-idempotent-retries-solution -bench=. -benchmem

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestDoPaymentFanout_PartialResults(t *testing.T) {
	// Backend that fails once then succeeds (to exercise retry + idempotency header existence).
	var calls int
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Idempotency-Key") == "" {
			http.Error(w, "missing Idempotency-Key", http.StatusBadRequest)
			return
		}
		calls++
		if calls == 1 {
			http.Error(w, "temporary", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{"ok": true})
	}))
	defer backend.Close()

	client := backend.Client()
	ctx := context.Background()

	results, err := doPaymentFanout(ctx, client, 100, []string{backend.URL}, 0, 200*time.Millisecond, 1, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("len(results) = %d, want 1", len(results))
	}
	if results[0].Err != "" || results[0].Status != http.StatusOK {
		t.Fatalf("result = %+v, want success", results[0])
	}
}

func TestPayHandler_StatusOK(t *testing.T) {
	body := `{"amount":10,"backends":1,"fail_rate":0}`
	req := httptest.NewRequest(http.MethodPost, "/pay", strings.NewReader(body))
	rec := httptest.NewRecorder()

	payHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200\nbody: %s", rec.Code, rec.Body.String())
	}
}

func BenchmarkDoPaymentFanout(b *testing.B) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Idempotency-Key") == "" {
			http.Error(w, "missing Idempotency-Key", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer backend.Close()

	client := backend.Client()
	backends := []string{backend.URL, backend.URL, backend.URL, backend.URL, backend.URL}

	b.ReportAllocs()
	for b.Loop() {
		_, _ = doPaymentFanout(context.Background(), client, 100, backends, 0, 200*time.Millisecond, 5, false)
	}
}
