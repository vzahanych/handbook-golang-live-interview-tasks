# Parallel sum with N goroutines

## Live interview task
Implement `ParallelSum(values []int64, workers int) (int64, error)`. Use at most `workers` goroutines, cover every element exactly once, and reject a non-positive worker count.

## Concepts covered
- slice partitioning
- bounded parallel reduction
- edge-case validation

## Candidate solution

```go
package main

import (
	"errors"
	"sync"
)

func ParallelSum(values []int64, workers int) (int64, error) {
	if workers <= 0 {
		return 0, errors.New("workers must be positive")
	}
	if len(values) == 0 {
		return 0, nil
	}
	workers = min(workers, len(values))
	partial := make([]int64, workers)

	var wg sync.WaitGroup
	for worker := range workers {
		start := worker * len(values) / workers
		end := (worker + 1) * len(values) / workers
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, value := range values[start:end] {
				partial[worker] += value
			}
		}()
	}
	wg.Wait()

	var sum int64
	for _, value := range partial {
		sum += value
	}
	return sum, nil
}
```

## Interview notes / pitfalls
- `worker*n/workers` boundaries handle uneven chunks without gaps.
- Cap workers to `len(values)` to avoid empty jobs.
- Each goroutine owns one result slot, so no mutex is required.
- Discuss integer overflow and benchmark against a sequential loop.

## Follow-up questions
- How would you cancel on overflow?
- When does the goroutine overhead exceed the benefit?
- How would you make the function generic over integer types?
