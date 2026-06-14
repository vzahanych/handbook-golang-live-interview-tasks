// Command flag-driven-colly-service is an HTTP service that scrapes web pages
// concurrently with Colly. Its initial targets come from repeated -url flags;
// callers can also POST their own target list. Each scrape uses a fresh async
// collector bounded to -workers parallel requests, and results come back in the
// same order the targets were given.
//
// Run:
//
//	go run ./18_mini_projects/flag-driven-colly-service \
//	  -addr=:8080 -workers=4 -timeout=8s \
//	  -url=https://example.com -url=https://go.dev
//
// API:
//
//	GET  /healthz   -> 200 "ok", never scrapes
//	GET  /targets   -> JSON array of the flag-configured targets
//	POST /scrape    -> scrape the configured targets (or a JSON body
//	                   {"urls":[...]}) and return ordered JSON results
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gocolly/colly/v2"
)

// userAgent identifies the bot honestly. Real services should use a contactable
// UA (a URL or email) so site operators can reach you — never spoof a browser.
const userAgent = "flag-driven-colly-service/1.0 (+https://example.com/bot)"

// maxBodyBytes caps how much of each response we download, so a giant or hostile
// page can't exhaust memory.
const maxBodyBytes = 1 << 20 // 1 MiB

// urlList is a flag.Value that accumulates a value each time -url is passed.
// This is how you implement a repeatable flag with the standard flag package.
type urlList []string

func (u *urlList) String() string { return strings.Join(*u, ",") }

// Set is called once per -url occurrence; we append instead of overwriting.
func (u *urlList) Set(v string) error {
	if v == "" {
		return errors.New("empty -url")
	}
	*u = append(*u, v)
	return nil
}

// Result is one target's outcome. Exactly one of Title/Error is meaningful, but
// Index/URL/Status are always set so the caller can correlate results to inputs.
type Result struct {
	Index  int    `json:"index"`
	URL    string `json:"url"`
	Status int    `json:"status"`
	Title  string `json:"title,omitempty"`
	Error  string `json:"error,omitempty"`
}

// config is the validated, immutable runtime configuration.
type config struct {
	addr    string
	workers int
	timeout time.Duration
	targets []string
}

// validateTarget rejects anything that isn't a plain absolute http(s) URL. A
// real service would also block private/link-local IPs (SSRF protection); we
// keep the scheme/host check here and note the rest.
func validateTarget(raw string) error {
	u, err := url.Parse(raw)
	if err != nil {
		return fmt.Errorf("invalid URL %q: %w", raw, err)
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("unsupported scheme %q (want http/https)", u.Scheme)
	}
	if u.Host == "" {
		return fmt.Errorf("missing host in %q", raw)
	}
	return nil
}

// scraper performs one bounded scrape. It is the unit we'd put behind an
// interface and test with httptest in a real codebase.
type scraper struct {
	workers int
	timeout time.Duration
}

// scrape visits every URL concurrently (capped at s.workers in flight) and
// returns one Result per input, in input order. The passed ctx lets the caller
// (an HTTP handler whose client disconnected, or a shutting-down process) cancel
// the work mid-flight.
func (s *scraper) scrape(ctx context.Context, urls []string) []Result {
	// Pre-size the results slice and index each Result so concurrent callbacks
	// can write their own slot without coordinating — no shared counter, no race.
	results := make([]Result, len(urls))
	for i, u := range urls {
		results[i] = Result{Index: i, URL: u}
	}

	// A fresh collector per job keeps callbacks and visited-state isolated
	// between requests. Async lets the LimitRule actually parallelize visits.
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent(userAgent),
		colly.MaxBodySize(maxBodyBytes),
	)
	c.SetRequestTimeout(s.timeout)

	// Bound concurrency to s.workers across all domains, with a small random
	// delay so we don't hammer any single host at a fixed cadence.
	if err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: s.workers,
		RandomDelay: 200 * time.Millisecond,
	}); err != nil {
		// Limit only errors on a bad glob; ours is constant, so this is fatal
		// programmer error rather than a runtime condition.
		log.Printf("limit rule: %v", err)
	}

	// indexOf pulls the input index we stashed in the per-request colly.Context.
	// This is what keeps results ordered and correctly attributed even though
	// callbacks fire concurrently and out of order.
	indexOf := func(r *colly.Response) int {
		return r.Ctx.GetAny("index").(int)
	}

	// Record the HTTP status for every response (success or not).
	c.OnResponse(func(r *colly.Response) {
		results[indexOf(r)].Status = r.StatusCode
	})

	// Grab the first <title> as the "scraped" datum. OnHTML may fire multiple
	// times for the selector; we keep the first non-empty one.
	c.OnHTML("title", func(e *colly.HTMLElement) {
		i := e.Request.Ctx.GetAny("index").(int)
		if results[i].Title == "" {
			results[i].Title = strings.TrimSpace(e.Text)
		}
	})

	// Transport/HTTP errors land here. Recording per-target errors (instead of
	// aborting) is what lets one failed site not discard the successful ones.
	c.OnError(func(r *colly.Response, err error) {
		i := indexOf(r)
		results[i].Status = r.StatusCode
		results[i].Error = err.Error()
	})

	// Cancellation: when ctx is done, abort the backend so in-flight and queued
	// visits stop instead of running to completion.
	stop := make(chan struct{})
	defer close(stop)
	go func() {
		select {
		case <-ctx.Done():
			results = markCancelled(results) // best-effort note on unfinished ones
		case <-stop:
		}
	}()

	// Queue each visit with its index attached via colly.Context.
	for i, u := range urls {
		cctx := colly.NewContext()
		cctx.Put("index", i)
		// Request (not Visit) lets us attach our context. nil body, nil headers.
		if err := c.Request("GET", u, nil, cctx, nil); err != nil {
			results[i].Error = err.Error()
		}
	}

	// Block until every queued visit (and its callbacks) has finished.
	c.Wait()
	return results
}

// markCancelled annotates any target that never produced a status/title/error.
// It's best-effort: callbacks may still be racing, so we only touch empty slots.
func markCancelled(rs []Result) []Result {
	for i := range rs {
		if rs[i].Status == 0 && rs[i].Title == "" && rs[i].Error == "" {
			rs[i].Error = "cancelled"
		}
	}
	return rs
}

// scrapeRequest is the optional POST body letting a caller override targets.
type scrapeRequest struct {
	URLs []string `json:"urls"`
}

// server wires the HTTP handlers to the validated config and the scraper.
type server struct {
	cfg     config
	scraper *scraper
}

func (s *server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", s.handleHealth)
	mux.HandleFunc("GET /targets", s.handleTargets)
	mux.HandleFunc("POST /scrape", s.handleScrape)
	return mux
}

func (s *server) handleHealth(w http.ResponseWriter, _ *http.Request) {
	// Liveness only — must never trigger a scrape.
	fmt.Fprintln(w, "ok")
}

func (s *server) handleTargets(w http.ResponseWriter, _ *http.Request) {
	// Encode an empty list as [] rather than null when nothing is configured.
	targets := s.cfg.targets
	if targets == nil {
		targets = []string{}
	}
	writeJSON(w, http.StatusOK, targets)
}

func (s *server) handleScrape(w http.ResponseWriter, r *http.Request) {
	// Default to the flag-configured targets; allow a JSON body to override.
	targets := s.cfg.targets
	if r.ContentLength != 0 {
		var body scrapeRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		if err := dec.Decode(&body); err != nil {
			http.Error(w, "invalid JSON body: "+err.Error(), http.StatusBadRequest)
			return
		}
		if len(body.URLs) > 0 {
			targets = body.URLs
		}
	}

	if len(targets) == 0 {
		http.Error(w, "no targets configured", http.StatusBadRequest)
		return
	}
	// Validate caller-supplied targets before doing any network work.
	for _, t := range targets {
		if err := validateTarget(t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	// r.Context() is cancelled if the client disconnects, so an abandoned
	// request stops useful scraping work.
	results := s.scraper.scrape(r.Context(), targets)
	writeJSON(w, http.StatusOK, results)
}

// writeJSON sends v as a pretty-printed JSON HTTP response.
func writeJSON(w http.ResponseWriter, status int, v any) {
	// Headers must be set before WriteHeader or the first body write.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Stream-encode directly into w (no intermediate []byte).
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ") // human-readable two-space indentation
	_ = enc.Encode(v)       // v's struct/json tags control field names
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func main() {
	// --- flags ---
	var targets urlList
	addr := flag.String("addr", ":8080", "HTTP listen address")
	workers := flag.Int("workers", 4, "max parallel requests per scrape")
	timeout := flag.Duration("timeout", 8*time.Second, "per-request timeout")
	flag.Var(&targets, "url", "scrape target (repeatable)")
	flag.Parse()

	// --- validate ALL flags before we ever bind/listen ---
	if *workers < 1 {
		fatalf("-workers must be >= 1")
	}
	if *timeout <= 0 {
		fatalf("-timeout must be > 0")
	}
	for _, t := range targets {
		if err := validateTarget(t); err != nil {
			fatalf("%v", err)
		}
	}

	cfg := config{addr: *addr, workers: *workers, timeout: *timeout, targets: targets}
	srv := &server{cfg: cfg, scraper: &scraper{workers: cfg.workers, timeout: cfg.timeout}}

	// http.Server with explicit timeouts so slow clients can't tie up resources.
	httpSrv := &http.Server{
		Addr:              cfg.addr,
		Handler:           srv.routes(),
		ReadHeaderTimeout: 5 * time.Second,
		// WriteTimeout must comfortably exceed a scrape; size it from the
		// per-request timeout so large target lists still finish.
		WriteTimeout: cfg.timeout + 25*time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// --- graceful shutdown on SIGINT/SIGTERM ---
	// signal.NotifyContext cancels rootCtx on the first signal; the second
	// signal (stop()) restores default behavior so a stuck shutdown is killable.
	rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("listening on %s (workers=%d timeout=%s targets=%d)",
			cfg.addr, cfg.workers, cfg.timeout, len(cfg.targets))
		// ListenAndServe returns ErrServerClosed on a clean Shutdown — not an error.
		if err := httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server error: %v", err)
			stop() // trigger shutdown path below
		}
	}()

	// Block until a signal (or a fatal serve error) cancels rootCtx.
	<-rootCtx.Done()
	log.Println("shutting down...")

	// Give in-flight requests a bounded grace period to finish.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
		_ = httpSrv.Close()
	}
	wg.Wait()
	log.Println("bye")
}
