# Parallel sort and merge

## Live interview task
Sort a large integer slice by sorting independent chunks concurrently and merging them.

## Candidate solution

```go
func ParallelSort(values []int, workers int) []int {
	out := slices.Clone(values)
	if len(out) < 2 || workers <= 1 {
		slices.Sort(out)
		return out
	}
	workers = min(workers, len(out))
	parts := make([][]int, workers)
	var wg sync.WaitGroup
	for i := range workers {
		lo, hi := i*len(out)/workers, (i+1)*len(out)/workers
		parts[i] = out[lo:hi]
		wg.Add(1)
		go func(part []int) { defer wg.Done(); slices.Sort(part) }(parts[i])
	}
	wg.Wait()
	for len(parts) > 1 {
		next := make([][]int, 0, (len(parts)+1)/2)
		for i := 0; i < len(parts); i += 2 {
			if i+1 == len(parts) { next = append(next, parts[i]); continue }
			next = append(next, merge(parts[i], parts[i+1]))
		}
		parts = next
	}
	return parts[0]
}
```

Assume `merge` performs the standard linear merge. Imports: `slices`, `sync`.

## Interview notes / pitfalls
- Return a clone unless mutation is part of the contract.
- Sorting chunks is parallel, while this compact merge loop is sequential.
- A parallel merge tree improves span but increases allocations.
