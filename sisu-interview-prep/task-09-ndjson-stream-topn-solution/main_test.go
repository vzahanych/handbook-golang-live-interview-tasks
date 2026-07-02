package main

// Two tests + a benchmark.
//   go test ./sisu-interview-prep/task-09-ndjson-stream-topn-solution -v
//   go test ./sisu-interview-prep/task-09-ndjson-stream-topn-solution -bench=. -benchmem

import (
	"strings"
	"testing"
)

func TestProcessNDJSON_TopAndErrors(t *testing.T) {
	in := "{\"partner\":\"B\"}\n{\"partner\":\"A\"}\n{\"partner\":\"A\"}\nnot json\n{\"partner\":}\n"
	top, rep, err := processNDJSON(strings.NewReader(in), 2, 2, 1<<20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if rep.TotalLines != 5 {
		t.Fatalf("total=%d", rep.TotalLines)
	}
	if rep.DecodedLines != 3 {
		t.Fatalf("decoded=%d", rep.DecodedLines)
	}
	if len(rep.Errors) != 2 {
		t.Fatalf("errors=%d", len(rep.Errors))
	}
	if len(top) != 2 {
		t.Fatalf("top=%v", top)
	}
	// A appears twice, B once.
	if top[0].Partner != "A" || top[0].Count != 2 {
		t.Fatalf("top[0]=%+v", top[0])
	}
	if top[1].Partner != "B" || top[1].Count != 1 {
		t.Fatalf("top[1]=%+v", top[1])
	}
}

func TestTopN_DeterministicTieBreak(t *testing.T) {
	counts := map[string]int{"B": 1, "A": 1}
	top := topNPartnerCounts(counts, 2)
	if top[0].Partner != "A" || top[1].Partner != "B" {
		t.Fatalf("got=%v", top)
	}
}

func BenchmarkProcessNDJSON(b *testing.B) {
	const (
		KB = 1 << 10
		MB = 1 << 20
	)
	const targetSize = 1 * MB

	var sb strings.Builder
	for sb.Len() < targetSize {
		sb.WriteString(`{"partner":"BankA","x":"`)
		sb.WriteString(strings.Repeat("y", 200))
		sb.WriteString(`"}` + "\n")
	}
	data := sb.String()

	b.SetBytes(int64(len(data)))
	b.ReportAllocs()
	for b.Loop() {
		if _, _, err := processNDJSON(strings.NewReader(data), 3, 0, 4<<20); err != nil {
			b.Fatal(err)
		}
	}
}
