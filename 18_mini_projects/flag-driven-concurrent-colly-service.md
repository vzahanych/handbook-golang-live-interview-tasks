# Flag-driven concurrent Colly service

## Live interview task
Build an HTTP service that receives its initial scrape targets from repeated `-url` flags and scrapes them concurrently with Colly.

## Command-line contract

```text
scraper -addr=:8080 -workers=4 -timeout=8s \
  -url=https://example.com -url=https://go.dev
```

## API contract
- `GET /targets` returns configured targets.
- `POST /scrape` starts one bounded scrape and returns ordered JSON results.
- `GET /healthz` returns success without scraping.
- On shutdown, stop accepting requests, cancel active jobs, and wait for callbacks.

## Implementation requirements
- Implement repeated URLs with `flag.Value`.
- Validate all flags before `ListenAndServe`.
- Use a fresh async collector per scrape job and a `LimitRule` with `Parallelism: workers`.
- Associate callbacks with input indexes through `colly.Context`.
- Configure an `http.Server` with timeouts and graceful signal shutdown.
- Reject unsafe targets and limit response bytes.
- Return per-target errors; one failed site must not discard successful results.

## Evaluation rubric
- Correctness: every target produces exactly one ordered result.
- Concurrency: active visits never exceed `workers`; no callback data races.
- Cancellation: client disconnect and process shutdown stop useful work.
- Reliability: timeouts, body limits, status checks, and error propagation exist.
- Testability: handler and scraper are injected behind small interfaces and tested with `httptest`.

## Follow-up questions
- Add a job queue with `202 Accepted` and polling.
- Cache results with TTL and request coalescing.
- Add process-wide rate limits, metrics, tracing, and authentication.
