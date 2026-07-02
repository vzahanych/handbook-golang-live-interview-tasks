package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Many integration flows are "log/event shaped": huge streams of JSON objects
// where you want a summary (top partners, top error codes, top products) without
// loading everything into memory.
//
// This starter reads NDJSON (one JSON object per line) from stdin and prints the
// Top N "partner" values by count. It works on tiny input but fails at scale.
//
// 1. The current code uses io.ReadAll which loads the entire input into memory.
//    Refactor to stream line-by-line (bufio.Scanner or bufio.Reader).
//
// 2. bufio.Scanner has a default token limit (64K). Real NDJSON lines can be
//    larger (embedded payloads). Make it handle large lines safely, and explain
//    the trade-offs (Scanner buffer vs Reader).
//
// 3. JSON decoding errors are ignored. Return a report of:
//    - total lines,
//    - successfully decoded lines,
//    - a slice of the first K decode errors (line number + error),
//    while still producing the Top N from valid lines.
//
// 4. Make Top N deterministic: if counts tie, sort by partner name ascending.
//
// Sample input:
//   echo '{"partner":"BankA"}'$'\n''{"partner":"BankB"}'$'\n''{"partner":"BankA"}' | go run . -n 2
//
// Bigger input:
//   python - <<'PY'
//   import json, random
//   partners = ["BankA","BankB","Auto1","RetailX","AirlineZ"]
//   for i in range(100000):
//       print(json.dumps({"partner": random.choice(partners), "i": i}))
//   PY | go run . -n 3

type event struct {
	Partner string `json:"partner"`
}

type partnerCount struct {
	Partner string
	Count   int
}

func topNPartnerCounts(counts map[string]int, n int) []partnerCount {
	out := make([]partnerCount, 0, len(counts))
	for p, c := range counts {
		out = append(out, partnerCount{Partner: p, Count: c})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count != out[j].Count {
			return out[i].Count > out[j].Count
		}
		return out[i].Partner < out[j].Partner
	})
	if n > len(out) {
		n = len(out)
	}
	return out[:n]
}

func main() {
	n := 3
	if len(os.Args) >= 3 && os.Args[1] == "-n" {
		_, _ = fmt.Sscanf(os.Args[2], "%d", &n)
	}

	// BUG: reads everything into memory.
	all, _ := io.ReadAll(os.Stdin)

	counts := map[string]int{}
	lines := strings.Split(string(all), "\n")
	for _, ln := range lines {
		if strings.TrimSpace(ln) == "" {
			continue
		}
		var e event
		_ = json.Unmarshal([]byte(ln), &e) // BUG: ignores decode errors
		if e.Partner == "" {
			continue
		}
		counts[e.Partner]++
	}

	// Print Top N.
	for _, pc := range topNPartnerCounts(counts, n) {
		fmt.Printf("%s %d\n", pc.Partner, pc.Count)
	}

	// TODO: print the error report described in question 3.
	_ = bufio.ErrAdvanceTooFar
}
