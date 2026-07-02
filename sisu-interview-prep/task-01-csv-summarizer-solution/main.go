package main

// 1. does reading the entire request body into memory scale for large uploads?
//    No. io.ReadAll pulls the whole upload into a single []byte, so a 10 MB (or
//    1 GB) file becomes a 10 MB (or 1 GB) allocation, and every concurrent
//    request holds its own copy — memory grows with size × concurrency and can
//    OOM the server. But r.Body is an io.Reader, i.e. a stream: we can hand it
//    straight to csv.NewReader(r.Body) and process the file row by row, keeping
//    only one record in memory at a time regardless of how large the upload is.
// 2. how should we surface errors when the CSV contains non-numeric cells?
//    The current code hides them: strconv.ParseFloat returns (0, err) and we
//    discard the err with `f, _ :=`, so a bad cell like "5.O" silently counts as
//    0 and the caller gets a wrong sum with a 200 OK. Whichever way we go, each
//    error needs enough context to locate the offending cell — the row number,
//    the column name, and the raw value, e.g. "row 3, column \"jan\": invalid
//    number \"5.O\"". From there we have two ways:
//      (a) fail fast: stop at the first bad cell and return a 400 error response.
//          Simple, but the client learns about only one problem per round-trip
//          and gets no result at all.
//      (b) accumulate: keep going through every record, collecting the parse
//          errors as we stream, and at the end return both the computed sums and
//          the list of errors. The client sees all bad cells in one pass and
//          still gets the partial totals.
//    The second way is much better: one round-trip surfaces every problem and the
//    caller still gets useful output. (207-style "partial success": sums plus an
//    errors array in the JSON body.)
// 3. if a csv contains invalid values, save the partial and return an error
//    This changes the contract from all-or-nothing to partial + error. As we
//    stream rows we keep the running sums, and when a cell fails to parse we stop
//    (or skip, but the ask is to stop), persist what we have so far to durable
//    storage — the per-column sums plus how far we got, e.g. the last row index
//    or byte offset — keyed by an upload/job id, and then return the error. So
//    the response is not a plain 400 with nothing: it carries the partial result
//    and points at the failing cell, and the saved checkpoint is what question 4
//    resumes from. Saving before returning is the important ordering — the work
//    already done survives the failure instead of being thrown away.
// 4. implement a way to resume the calculation from where it was interrupted
//    Resume reads back the checkpoint that question 3 saved. Each upload gets a
//    stable job id (the client sends it, or we derive it from a hash of the
//    content), and against that id we store the running per-column sums plus the
//    position we reached — the row index or byte offset of the last good record.
//    On a retry for the same id we load that checkpoint, seek the CSV stream past
//    the rows already counted, and keep adding onto the saved sums instead of
//    starting from zero. If there is no checkpoint we start fresh at row 0; when
//    we reach EOF we mark the job done so a later retry returns the final result
//    rather than recomputing. For this to be correct the parse must be
//    deterministic and the checkpoint must be written atomically — sums and
//    position together — so we never resume from a half-updated state.
// 5. can we add a benchmark or integration test that parses a sample 10 MB CSV?
//    Yes, and they answer two different questions. A benchmark (func BenchmarkX(
//    b *testing.B)) measures the parser in isolation: generate ~10 MB of CSV once
//    before the loop, reset the timer, then parse it b.N times and run with
//    `go test -bench=. -benchmem` to get ns/op plus allocs/op and B/op. That last
//    number is what proves the streaming rewrite from question 1 — a constant,
//    small footprint instead of memory that tracks file size. An integration test
//    (func TestX(t *testing.T)) exercises the real HTTP path: spin up the handler
//    with httptest.NewServer, POST the 10 MB body, and assert on the status code
//    and the decoded sums (and, for the accumulate design, the errors array). To
//    keep the repo small we generate the sample in code rather than commit a 10 MB
//    fixture, and gate the slow one behind `if testing.Short() { t.Skip() }` so
//    `go test -short` stays fast. The benchmark guards performance/memory; the
//    integration test guards correctness end to end.

// city,jan,feb,mar
// Berlin,1.2,2.3,5
// Paris,3.4,4.5,6.7
// Madrid,5.O,6.1,7.2
// Milan,2,3,4

import (
	"context" // graceful shutdown: signal-aware context + shutdown deadline
	"encoding/csv"
	"encoding/json"
	"fmt" // answer 2: format the located cell error
	"io"
	"log"
	"net/http"
	"os"        // graceful shutdown: os.Interrupt
	"os/signal" // graceful shutdown: catch SIGINT/SIGTERM
	"strconv"
	"syscall" // graceful shutdown: syscall.SIGTERM
	"time"    // graceful shutdown: shutdown timeout
	// "strings" // no longer needed: we stream r.Body instead of strings.NewReader(string(data))
)

type summary struct {
	Column string  `json:"column"`
	Sum    float64 `json:"sum"`
}

// cellError locates a cell that failed to parse (answer to question 2): the row
// number, the column name, and the raw value. It implements error so the
// fail-fast path can return it directly.
type cellError struct {
	Row    int    `json:"row"`
	Column string `json:"column"`
	Value  string `json:"value"`
}

func (e cellError) Error() string {
	return fmt.Sprintf("row %d, column %q: invalid number %q", e.Row, e.Column, e.Value)
}

// func summarizeCSV(data []byte) ([]summary, error) {           // answer 1: buffered whole file
// 	reader := csv.NewReader(strings.NewReader(string(data)))    // answer 1: buffered whole file
// func summarizeCSV(body io.Reader) ([]summary, error) {        // answer 1: streamed, but still dropped bad cells
// answer 2: add failFast switch and return the list of cell errors alongside the sums.
func summarizeCSV(body io.Reader, failFast bool) ([]summary, []cellError, error) {
	reader := csv.NewReader(body) // stream: parse row by row, never buffer the whole file

	header, err := reader.Read()
	if err != nil {
		// return nil, err
		return nil, nil, err
	}

	sums := make([]float64, len(header))
	var cellErrors []cellError // answer 2: collect bad cells instead of silently dropping them
	row := 1                   // header was line 1; data rows start at 2

	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		row++
		for i, v := range rec {
			if i == 0 {
				continue // column 0 is the row label (e.g. city), not a number
			}
			// f, _ := strconv.ParseFloat(v, 64) // answer 2: was discarding the parse error
			// sums[i] += f
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				ce := cellError{Row: row, Column: header[i], Value: v}
				if failFast {
					return nil, nil, ce // (a) fail fast: stop at the first bad cell
				}
				cellErrors = append(cellErrors, ce) // (b) accumulate and keep going
				continue
			}
			sums[i] += f
		}
	}

	// out := make([]summary, len(header))
	// for i, col := range header {
	// 	out[i] = summary{Column: col, Sum: sums[i]}
	// }
	out := make([]summary, 0, len(header)-1)
	for i, col := range header {
		if i == 0 {
			continue // label column has no sum
		}
		out = append(out, summary{Column: col, Sum: sums[i]})
	}
	// return out, nil
	return out, cellErrors, nil
}

func summarizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
     
	// Answer to question 1:
	// data, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "cannot read body", http.StatusBadRequest)
	// 	return
	// }

	// list, err := summarizeCSV(data)
	// answer 2: switch — ?fail_fast=true stops at the first bad cell; default collects all.
	failFast := r.URL.Query().Get("fail_fast") == "true"

	// list, err := summarizeCSV(r.Body) // stream the body straight into the CSV parser
	list, cellErrors, err := summarizeCSV(r.Body, failFast) // stream the body straight into the CSV parser
	if err != nil {
		// http.Error(w, "csv error", http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest) // fail-fast path: report the located cell
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(list)
	// answer 2: 207-style partial success — return the sums plus the list of bad cells.
	if len(cellErrors) > 0 {
		w.WriteHeader(http.StatusMultiStatus)
	}
	json.NewEncoder(w).Encode(map[string]any{
		"summaries": list,
		"errors":    cellErrors,
	})
}

func main() {
	// simplest graceful shutdown: run the server in a goroutine, wait for a
	// signal, then Shutdown() so in-flight requests finish and the port frees.

	http.HandleFunc("/summarize", summarizeHandler) // registers on http.DefaultServeMux
	srv := &http.Server{Addr: ":8080"}              // Handler nil ⇒ uses DefaultServeMux

	// ctx is cancelled on the first SIGINT (Ctrl+C) or SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// log.Fatal(http.ListenAndServe(":8080", nil)) // no server handle ⇒ can't Shutdown
	go func() {
		log.Println("CSV Summarizer listening on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	<-ctx.Done() // block until a signal arrives
	log.Println("shutting down, waiting for in-flight requests...")

	// Give active requests up to 5s to finish, then force-close.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}
	log.Println("server stopped cleanly")
}


// How to run and test
// --------------------
// This is an HTTP server, not a CLI tool: there are no -fail_fast/-csv_file
// flags. Start it, then drive it with curl. fail_fast is a query parameter.
//
// 1. Start the server (from this folder):
//      go run .
//    It listens on :8080 and logs "CSV Summarizer listening on :8080".
//
// 2. Accumulate mode (default) — parses every row, returns 207 with the sums
//    plus the list of bad cells. Use --data-binary so curl keeps the newlines:
//      curl -si --data-binary @sample.csv http://localhost:8080/summarize
//
// 3. Fail-fast mode — stops at the first bad cell and returns 400 with the
//    located error (row, column, value):
//      curl -si --data-binary @sample.csv 'http://localhost:8080/summarize?fail_fast=true'
//
// 4. Happy path — a CSV with only numeric cells returns 200 and just the sums:
//      printf 'city,jan\nBerlin,1.2\nParis,3.4\n' | curl -si --data-binary @- http://localhost:8080/summarize
//
// Note: only POST is allowed; a GET returns 405 method not allowed.
