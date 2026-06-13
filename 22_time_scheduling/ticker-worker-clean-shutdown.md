# Ticker worker with clean shutdown

## Live interview task
Run a job periodically until context cancellation, with no overlapping executions.

## Candidate solution

```go
ticker := time.NewTicker(interval)
defer ticker.Stop()
for {
	select {
	case <-ctx.Done(): return ctx.Err()
	case <-ticker.C:
		if err := job(ctx); err != nil { return err }
	}
}
```

## Interview notes / pitfalls
- Tickers may drop ticks when the receiver is slow.
- Decide whether the first job runs immediately or after one interval.
- For overlapping jobs, add an explicit concurrency policy.
