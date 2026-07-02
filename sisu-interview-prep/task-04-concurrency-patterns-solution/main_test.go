package main

// Two tests + a benchmark, same package as main.go so we can call unexported
// functions.
//   go test ./sisu-interview-prep/task-04-concurrency-patterns-solution -v
//   go test ./sisu-interview-prep/task-04-concurrency-patterns-solution -bench=. -benchmem

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAggregate(t *testing.T) {
	txs := []Transaction{
		{"EUR", 1200}, {"USD", 999}, {"EUR", 350},
		{"GBP", 5000}, {"USD", 1}, {"EUR", 8000},
	}
	ctx := context.Background()
	got, err := aggregate(ctx, txs, 2, 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got["EUR"] != 1200+350+8000 {
		t.Fatalf("EUR = %d, want %d", got["EUR"], 1200+350+8000)
	}
	if got["USD"] != 999+1 {
		t.Fatalf("USD = %d, want %d", got["USD"], 999+1)
	}
	if got["GBP"] != 5000 {
		t.Fatalf("GBP = %d, want %d", got["GBP"], 5000)
	}
}

func TestAggregate_Cancelled(t *testing.T) {
	txs := make([]Transaction, 0, 1000)
	for i := 0; i < 1000; i++ {
		txs = append(txs, Transaction{Currency: "EUR", Amount: 1})
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := aggregate(ctx, txs, 4, 16)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func BenchmarkAggregate(b *testing.B) {
	txs := make([]Transaction, 0, 1<<16)
	for i := 0; i < cap(txs); i++ {
		txs = append(txs, Transaction{Currency: []string{"EUR", "USD", "GBP"}[i%3], Amount: int64(i % 1000)})
	}

	b.ReportAllocs()
	for b.Loop() {
		_, _ = aggregate(context.Background(), txs, 8, 1024)
	}
}

// Integration-ish test in the same style as other tasks: we don't have HTTP here,
// but keep one minimal handler test as an example of using httptest.
func TestMinimalHandlerExample(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodPost, "/x", strings.NewReader("x"))
	req = req.WithContext(context.Background())
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}

	_ = time.Second
}
