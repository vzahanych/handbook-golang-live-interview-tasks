# Reset a timer correctly

## Live interview task
Reset a reusable timer without allowing a stale tick to trigger the next operation.

## Candidate pattern

```go
if !timer.Stop() {
	select { case <-timer.C: default: }
}
timer.Reset(delay)
```

## Interview notes / pitfalls
- Timer reset semantics depend on the Go version; explain the version targeted by the code.
- One owner goroutine is easier to reason about than concurrent `Stop` and `Reset` calls.
- Prefer a fresh timer when reuse is not performance-critical.
