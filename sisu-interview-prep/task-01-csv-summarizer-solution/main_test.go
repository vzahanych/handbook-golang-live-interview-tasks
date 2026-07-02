package main

// Two tests + a benchmark, same package as main.go so we can call the
// unexported functions.
//   go test ./... -v
//   go test -bench=. -benchmem
import (
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Unit test: the parser sums numeric columns, skips the label column, and
// collects a bad cell instead of silently dropping it.
func TestSummarizeCSV(t *testing.T) {
	in := "city,jan,feb\nBerlin,1.2,2.3\nParis,3.4,x\n"

	sums, cellErrs, err := summarizeCSV(strings.NewReader(in), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cellErrs) != 1 {
		t.Fatalf("errors = %d, want 1 (%v)", len(cellErrs), cellErrs)
	}
	// jan = 1.2 + 3.4 = 4.6 (label column "city" is excluded)
	if sums[0].Column != "jan" || math.Abs(sums[0].Sum-4.6) > 1e-9 {
		t.Errorf("got %+v, want jan sum 4.6", sums[0])
	}
}

// Integration test: a clean POST returns 200 through the real handler.
func TestSummarizeHandler(t *testing.T) {
	body := "city,jan\nBerlin,1.2\nParis,3.4\n"
	req := httptest.NewRequest(http.MethodPost, "/summarize", strings.NewReader(body))
	rec := httptest.NewRecorder()

	summarizeHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200\nbody: %s", rec.Code, rec.Body.String())
	}
}

// Benchmark: build ~1 MB of CSV once, then time the parser. -benchmem reports
// B/op and allocs/op — the numbers that show streaming keeps memory flat.
func BenchmarkSummarizeCSV(b *testing.B) {
	const (
		KB = 1 << 10 // 1024 bytes
		MB = 1 << 20 // 1024 * KB
	)
	const targetSize = 1 * MB // change this to resize the sample, e.g. 10 * MB

	var sb strings.Builder
	sb.WriteString("city,jan,feb,mar\n")
	for sb.Len() < targetSize {
		sb.WriteString("Berlin,1.2,2.3,5\n")
	}
	data := sb.String()

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	for b.Loop() {
		if _, _, err := summarizeCSV(strings.NewReader(data), false); err != nil {
			b.Fatal(err)
		}
	}
}
