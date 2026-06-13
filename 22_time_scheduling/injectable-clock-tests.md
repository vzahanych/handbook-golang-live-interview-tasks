# Injectable clock for deterministic tests

## Live interview task
Refactor TTL logic so tests do not sleep.

## Candidate interface

```go
type Clock interface { Now() time.Time }
type realClock struct{}
func (realClock) Now() time.Time { return time.Now() }
```

## Interview notes / pitfalls
- A `Now` function is often enough; do not create a large abstraction unnecessarily.
- Timer-driven code may require a fake scheduler, not only a fake clock.
- Preserve monotonic time behavior when comparing durations inside one process.
