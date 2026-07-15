package main

import (
	"context" // graceful shutdown: signal-aware context + shutdown deadline
	"encoding/json"
	"fmt" // answer 3: format an invalid ?n= into a located 400 error
	"io"  // answer 1: the body is an io.Reader we stream into the JSON decoder
	"log"
	"net/http"
	"os"        // graceful shutdown: os.Interrupt
	"os/signal" // graceful shutdown: catch SIGINT/SIGTERM
	"sort"
	"strconv" // answer 2: parse the ?n= top-N switch
	"syscall" // graceful shutdown: syscall.SIGTERM
	"time"    // graceful shutdown: shutdown timeout
)

// We want to build an API endpoint to analyze web server logs.
//
// 1.  **Implement the `analyzeHandler` function.**
//     - It should decode the incoming JSON request body into the `LogRequest` struct.
//     - Handle potential JSON decoding errors gracefully.
//     Answer: the request body is an io.Reader (a stream). Three approaches:
//     (a) io.ReadAll + json.Unmarshal — buffers the whole upload, then parses it.
//     (b) json.NewDecoder(r.Body).Decode(&req) — reads from the stream, but still
//     materializes every log entry into req.Logs, so memory is O(number of logs).
//     (c) json.NewDecoder token walk — decode one LogEntry at a time from the
//     "logs" array and count immediately; only the counts map stays in memory
//     (O(unique paths)). We use (c). "Gracefully" means a decode failure returns
//     400 with the decoder's message (not a 500), so a malformed body is a client
//     error the caller can see and fix.
//
// 2.  **Process the logs to find the top 3 most visited pages.**
//     - Count the occurrences of each `path`.
//     - Determine the top 3 paths based on their counts.
//     Answer: counting is a single pass into a map[string]int keyed by path.
//     Picking the top 3 out of that map admits two algorithms: (a) sort every
//     distinct path by count and take the first N — O(u log u) in the number of
//     unique paths u, trivial to write; (b) keep a min-heap of size N as you scan
//     — O(u log N), which wins when u is huge and N is tiny. We take (a): u
//     (distinct URLs) is small in practice and a full sort is clearer under time
//     pressure; the heap is the obvious next optimization. Two design switches
//     make the "top 3" explicit rather than hard-coded: ?n= chooses how many
//     (default 3), and ?success_only=true counts only 2xx/3xx responses, since a
//     page that mostly 500s is not really a "visited" page.
//
// 3.  **Send the response.**
//     - The response should be a JSON object matching the `TopPagesResponse` struct.
//     - The `top_pages` slice should be sorted in descending order of `count`.
//     - Set the correct `Content-Type` header to `application/json`.
//     Answer: sorting by count descending is not a total order on its own — two
//     paths with the same count can come out in either order, because Go's map
//     iteration is randomized. That non-determinism is a real bug: the same input
//     yields different top-3 across runs and tests flake. So we sort by count
//     descending and break ties by path ascending, which is stable and testable.
//     Then truncate to N, set Content-Type: application/json, and encode.
//
// Request Example:
// {
//   "logs": [
//     {"path": "/home", "status_code": 200},
//     {"path": "/products", "status_code": 200},
//     {"path": "/home", "status_code": 200},
//     {"path": "/about", "status_code": 200},
//     {"path": "/products", "status_code": 200},
//     {"path": "/home", "status_code": 500}
//   ]
// }
//
// Response Example:
// {
//   "top_pages": [
//     {"path": "/home", "count": 3},
//     {"path": "/products", "count": 2},
//     {"path": "/about", "count": 1}
//   ]
// }

type LogRequest struct {
	Logs []LogEntry `json:"logs"`
}

type LogEntry struct {
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
}

type TopPagesResponse struct {
	TopPages []PageCount `json:"top_pages"`
}

type PageCount struct {
	Path  string `json:"path"`
	Count int    `json:"count"`
}

// analyze is the core, HTTP-free function (answers 1–3). It walks the JSON
// stream one log entry at a time, counts paths in one pass, then returns the
// top n pages sorted by count descending with a path-ascending tie-break.
//   - successOnly (answer 2): when true, only 2xx/3xx responses count as visits.
//   - n (answer 2): how many pages to return; the caller clamps it to >= 1.
func analyze(body io.Reader, page int, successOnly bool) (TopPagesResponse, error) {
	dec := json.NewDecoder(body)

	tok, err := dec.Token()
	if err != nil {
		return TopPagesResponse{}, err
	}
	if d, ok := tok.(json.Delim); !ok || d != '{' {
		return TopPagesResponse{}, fmt.Errorf("expected JSON object")
	}

	counts := make(map[string]int)
	for dec.More() {
		keyTok, err := dec.Token()
		if err != nil {
			return TopPagesResponse{}, err
		}
		key, ok := keyTok.(string)
		if !ok {
			return TopPagesResponse{}, fmt.Errorf("expected object key")
		}

		if key != "logs" {
			var skip json.RawMessage
			if err := dec.Decode(&skip); err != nil {
				return TopPagesResponse{}, err
			}
			continue
		}

		tok, err = dec.Token()
		if err != nil {
			return TopPagesResponse{}, err
		}
		if d, ok := tok.(json.Delim); !ok || d != '[' {
			return TopPagesResponse{}, fmt.Errorf("expected logs array")
		}

		// answer 1: decode one entry at a time — never hold the full logs slice.
		for dec.More() {
			var e LogEntry
			if err := dec.Decode(&e); err != nil {
				return TopPagesResponse{}, err
			}
			// answer 2: success_only switch — skip non-2xx/3xx so error pages don't count as visits.
			if successOnly && (e.StatusCode < 200 || e.StatusCode >= 400) {
				continue
			}
			counts[e.Path]++
		}

		if _, err := dec.Token(); err != nil { // closing ']'
			return TopPagesResponse{}, err
		}
	}

	if _, err := dec.Token(); err != nil { // closing '}'
		return TopPagesResponse{}, err
	}

	pages := make([]PageCount, 0, len(counts))
	for path, c := range counts {
		pages = append(pages, PageCount{Path: path, Count: c})
	}

	// answer 3: count desc, then path asc — a total order so the top-N is
	// deterministic instead of depending on randomized map iteration.
	sort.Slice(pages, func(i, j int) bool {
		if pages[i].Count != pages[j].Count {
			return pages[i].Count > pages[j].Count
		}
		return pages[i].Path < pages[j].Path
	})

	if page < len(pages) {
		pages = pages[:page] // top N
	}
	return TopPagesResponse{TopPages: pages}, nil
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	// answer 1: this endpoint mutates/reads a body, so restrict it to POST; a GET returns 405.
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// answer 2: switches — ?n= chooses how many pages (default 3), ?success_only=true
	// counts only 2xx/3xx responses.
	n := 3
	if s := r.URL.Query().Get("n"); s != "" {
		v, err := strconv.Atoi(s)
		if err != nil || v < 1 {
			http.Error(w, fmt.Sprintf("invalid n %q: want a positive integer", s), http.StatusBadRequest)
			return
		}
		n = v
	}
	successOnly := r.URL.Query().Get("success_only") == "true"

	// TODO: Implement your logic here.
	//
	// // A placeholder response to show it's working.
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]string{"status": "not implemented yet"})

	// answer 1: stream the body entry-by-entry; a decode failure is a 400, not a 500.
	resp, err := analyze(r.Body, n, successOnly)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid JSON body: %v", err), http.StatusBadRequest)
		return
	}

	// answer 3: correct Content-Type, then encode the sorted top-N.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// simplest graceful shutdown: run the server in a goroutine, wait for a
	// signal, then Shutdown() so in-flight requests finish and the port frees.

	http.HandleFunc("/analyze/top-pages", analyzeHandler) // registers on http.DefaultServeMux
	srv := &http.Server{Addr: ":8080"}                    // Handler nil ⇒ uses DefaultServeMux

	// ctx is cancelled on the first SIGINT (Ctrl+C) or SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// log.Println("Server starting on :8080...")
	// if err := http.ListenAndServe(":8080", nil); err != nil { // no server handle ⇒ can't Shutdown
	// 	log.Fatalf("Could not start server: %s\n", err)
	// }
	go func() {
		log.Println("Top-Pages analyzer listening on :8080")
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
// This is an HTTP server, not a CLI tool: there are no flags. Start it, then
// drive it with curl. The top-N count and the success filter are query params.
//
// 1. Start the server (from this folder):
//      go run .
//    It listens on :8080 and logs "Top-Pages analyzer listening on :8080".
//
// 2. Default — top 3 pages counting every request (any status code). Use
//    --data-binary so curl sends the file bytes exactly (preserves newlines):
//      curl -si --data-binary @sample.json http://localhost:8080/analyze/top-pages
//
// 3. success_only switch — count only 2xx/3xx responses, so pages that mostly
//    error drop in the ranking:
//      curl -si --data-binary @sample.json 'http://localhost:8080/analyze/top-pages?success_only=true'
//
// 4. n switch — return the top 5 instead of 3 (an invalid n returns 400):
//      curl -si --data-binary @sample.json 'http://localhost:8080/analyze/top-pages?n=5'
//
// 5. Inline happy path:
//      curl -si --data-binary '{"logs":[{"path":"/a","status_code":200},{"path":"/a","status_code":200},{"path":"/b","status_code":200}]}' http://localhost:8080/analyze/top-pages
//
// Note: only POST is allowed; a GET returns 405 method not allowed. A malformed
// JSON body returns 400 with the decoder's message.
