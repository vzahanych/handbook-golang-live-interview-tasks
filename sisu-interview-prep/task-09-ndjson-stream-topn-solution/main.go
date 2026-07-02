package main

// 1. The current code uses io.ReadAll which loads the entire input into memory.
//    Refactor to stream line-by-line (bufio.Scanner or bufio.Reader).
//    Streaming keeps memory O(max line size) instead of O(total input size),
//    which matters for multi-GB logs or many concurrent requests.
// 2. bufio.Scanner has a default token limit (64K). Real NDJSON lines can be
//    larger (embedded payloads). Make it handle large lines safely, and explain
//    the trade-offs (Scanner buffer vs Reader).
//    Scanner is convenient but needs Buffer() raised; very large lines can still
//    be risky because they allocate that much. A bufio.Reader avoids the token
//    limit and can read chunks/lines, but you must handle partial reads and
//    delimiters carefully. Here we use Scanner with an explicit max line size so
//    behavior is predictable.
// 3. JSON decoding errors are ignored. Return a report of:
//    - total lines,
//    - successfully decoded lines,
//    - a slice of the first K decode errors (line number + error),
//    while still producing the Top N from valid lines.
//    Keep counting valid lines and record first K errors without aborting.
// 4. Make Top N deterministic: if counts tie, sort by partner name ascending.
//    Already implemented in topNPartnerCounts.

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

type decodeErr struct {
	Line int    `json:"line"`
	Err  string `json:"err"`
}

type report struct {
	TotalLines   int         `json:"total_lines"`
	DecodedLines int         `json:"decoded_lines"`
	Errors       []decodeErr `json:"errors"`
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

func processNDJSON(r io.Reader, n int, maxErrs int, maxLineBytes int) ([]partnerCount, report, error) {
	if n <= 0 {
		n = 1
	}
	if maxErrs < 0 {
		maxErrs = 0
	}
	if maxLineBytes <= 0 {
		maxLineBytes = 1 << 20 // 1MB
	}

	counts := map[string]int{}
	rep := report{Errors: make([]decodeErr, 0, maxErrs)}

	sc := bufio.NewScanner(r)
	// answer 2: raise Scanner buffer limit explicitly.
	sc.Buffer(make([]byte, 64*1024), maxLineBytes)

	lineNo := 0
	for sc.Scan() {
		lineNo++
		rep.TotalLines++
		ln := strings.TrimSpace(sc.Text())
		if ln == "" {
			continue
		}

		var e event
		// if err := json.Unmarshal([]byte(ln), &e); err != nil { // old: copies line into []byte
		if err := json.NewDecoder(strings.NewReader(ln)).Decode(&e); err != nil {
			if len(rep.Errors) < maxErrs {
				rep.Errors = append(rep.Errors, decodeErr{Line: lineNo, Err: err.Error()})
			}
			continue
		}
		rep.DecodedLines++
		if e.Partner == "" {
			continue
		}
		counts[e.Partner]++
	}
	if err := sc.Err(); err != nil {
		return nil, rep, err
	}

	return topNPartnerCounts(counts, n), rep, nil
}

func main() {
	n := 3
	if len(os.Args) >= 3 && os.Args[1] == "-n" {
		_, _ = fmt.Sscanf(os.Args[2], "%d", &n)
	}

	// all, _ := io.ReadAll(os.Stdin) // old: buffers whole input in memory
	top, rep, err := processNDJSON(os.Stdin, n, 5, 4<<20) // 4MB max line
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	// Output as JSON so it's easy to diff and script.
	_ = json.NewEncoder(os.Stdout).Encode(map[string]any{
		"top":    top,
		"report": rep,
	})
}
