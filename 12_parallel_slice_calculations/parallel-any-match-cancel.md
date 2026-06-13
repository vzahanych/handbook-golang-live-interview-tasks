# Parallel any-match with cancellation

## Live interview task
Search a slice in parallel and return as soon as any worker finds a matching value.

## Candidate solution

```go
func ParallelAny[T any](ctx context.Context, values []T, workers int, match func(T) bool) bool {
	if workers <= 0 || len(values) == 0 {
		return false
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	jobs := make(chan T)
	found := make(chan struct{}, 1)
	var wg sync.WaitGroup

	for range min(workers, len(values)) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range jobs {
				if match(value) {
					select { case found <- struct{}{}: default: }
					cancel()
					return
				}
			}
		}()
	}
	go func() {
		defer close(jobs)
		for _, value := range values {
			select {
			case jobs <- value:
			case <-ctx.Done(): return
			}
		}
	}()
	wg.Wait()
	return len(found) != 0
}
```

Imports: `context`, `sync`.

## Interview notes / pitfalls
- The buffered result channel prevents losing the first match.
- The producer must observe cancellation or it can leak.
- Clarify whether caller cancellation and "not found" must be distinguishable.
