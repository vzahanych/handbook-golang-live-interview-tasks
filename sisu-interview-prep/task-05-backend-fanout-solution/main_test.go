package main

// Two tests + a benchmark.
//   go test ./sisu-interview-prep/task-05-backend-fanout-solution -v
//   go test ./sisu-interview-prep/task-05-backend-fanout-solution -bench=. -benchmem

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchAll_PartialResults(t *testing.T) {
	okSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello"))
	}))
	defer okSrv.Close()

	slowSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("slow"))
	}))
	defer slowSrv.Close()

	client := &http.Client{}
	urls := []string{okSrv.URL, slowSrv.URL}

	got, err := fetchAll(context.Background(), client, urls, 50*time.Millisecond, 2, false)
	if err != nil {
		t.Fatalf("unexpected error in partial mode: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("len=%d, want 2", len(got))
	}
	if got[0].Err != nil || got[0].Bytes != 5 {
		t.Fatalf("got[0]=%+v", got[0])
	}
	if got[1].Err == nil {
		t.Fatalf("expected timeout error for slow backend, got %+v", got[1])
	}
}

func TestFetchAll_FailFastReturnsError(t *testing.T) {
	badSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusServiceUnavailable)
	}))
	defer badSrv.Close()

	okSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))
	defer okSrv.Close()

	client := &http.Client{}
	urls := []string{badSrv.URL, okSrv.URL}

	_, err := fetchAll(context.Background(), client, urls, 500*time.Millisecond, 2, true)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func BenchmarkFetchAll(b *testing.B) {
	okSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(make([]byte, 32<<10)) // 32KB
	}))
	defer okSrv.Close()

	client := &http.Client{}
	urls := []string{okSrv.URL, okSrv.URL, okSrv.URL, okSrv.URL}

	b.ReportAllocs()
	for b.Loop() {
		_, _ = fetchAll(context.Background(), client, urls, 500*time.Millisecond, 4, false)
	}
}
