package main

// Two tests + a benchmark, same package as main.go so we can call the
// unexported functions.
//   go test ./... -v
//   go test -bench=. -benchmem
import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Unit test: analyze counts paths, returns the top N sorted by count descending,
// breaks ties by path ascending, and — with success_only — drops error pages.
func TestAnalyze(t *testing.T) {
	// /checkout has 3 hits but 2 are 500s; /about has 2 clean hits.
	in := `{"logs":[
		{"path":"/home","status_code":200},
		{"path":"/home","status_code":200},
		{"path":"/checkout","status_code":500},
		{"path":"/checkout","status_code":500},
		{"path":"/checkout","status_code":200},
		{"path":"/about","status_code":200},
		{"path":"/about","status_code":200}
	]}`

	// Default: count every request. Ties break by path ascending, so at count 2
	// "/about" precedes "/checkout"... but /checkout has 3, /home 2, /about 2.
	resp, err := analyze(strings.NewReader(in), 3, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := resp.TopPages
	if len(got) != 3 {
		t.Fatalf("top pages = %d, want 3 (%v)", len(got), got)
	}
	// /checkout=3, then /home=2 and /about=2 tie → path asc puts /about first.
	if got[0] != (PageCount{"/checkout", 3}) ||
		got[1] != (PageCount{"/about", 2}) ||
		got[2] != (PageCount{"/home", 2}) {
		t.Errorf("got %+v, want [/checkout:3 /about:2 /home:2]", got)
	}

	// success_only: /checkout drops to 1 (only one 200), so /home and /about lead.
	resp, err = analyze(strings.NewReader(in), 3, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got = resp.TopPages
	if got[0] != (PageCount{"/about", 2}) || got[1] != (PageCount{"/home", 2}) || got[2] != (PageCount{"/checkout", 1}) {
		t.Errorf("success_only got %+v, want [/about:2 /home:2 /checkout:1]", got)
	}
}

// Integration test: a clean POST returns 200 through the real handler.
func TestAnalyzeHandler(t *testing.T) {
	body := `{"logs":[{"path":"/home","status_code":200},{"path":"/home","status_code":200},{"path":"/x","status_code":200}]}`
	req := httptest.NewRequest(http.MethodPost, "/analyze/top-pages", strings.NewReader(body))
	rec := httptest.NewRecorder()

	analyzeHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200\nbody: %s", rec.Code, rec.Body.String())
	}
}

// Benchmark: build ~1 MB of JSON logs once, then time analyze (decode + count +
// sort). -benchmem reports B/op and allocs/op — the numbers that show streaming
// the body through the decoder keeps memory tied to the parsed data, not size².
func BenchmarkAnalyze(b *testing.B) {
	const (
		KB = 1 << 10 // 1024 bytes
		MB = 1 << 20 // 1024 * KB
	)
	const targetSize = 1 * MB // change this to resize the sample, e.g. 10 * MB

	paths := []string{"/home", "/products", "/about", "/contact", "/blog", "/search"}
	var sb strings.Builder
	sb.WriteString(`{"logs":[`)
	for i := 0; sb.Len() < targetSize; i++ {
		if i > 0 {
			sb.WriteByte(',')
		}
		fmt.Fprintf(&sb, `{"path":%q,"status_code":200}`, paths[i%len(paths)])
	}
	sb.WriteString(`]}`)
	data := sb.String()

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	for b.Loop() {
		if _, err := analyze(strings.NewReader(data), 3, false); err != nil {
			b.Fatal(err)
		}
	}
}
