package main

// 1. `sequence` is a goroutine that receives ONE integer and returns ONE string
//    (the FizzBuzz sequence for 1..n). Modify the code so it accepts MULTIPLE
//    integers and returns MULTIPLE strings — one result string per input, in the
//    same order as the input.
//    Keep the goroutine-based API but call it once per input and collect results
//    into a `[]string` aligned with the input indexes. Preserving order is
//    simplest if each worker writes to `results[i]`.
// 2. Right now the calls happen one at a time. Modify the code to handle
//    MULTIPLE SIMULTANEOUS calls to the goroutine. Use a sync.WaitGroup to wait
//    for all of them, and a sync.Mutex (or channels / atomics) to avoid a data
//    race while collecting the results. Running `go run -race .` must be clean.
//    Concurrency is safe if we avoid shared append. Preallocate result slots and
//    write by index, or protect shared structures with a mutex.
// 3. Assume the list of integers can be very large (millions). Spawning one
//    goroutine per input exhausts memory and the scheduler. Bound the
//    concurrency with a worker pool of a fixed size, while still returning the
//    results in input order.
//    Use a bounded worker pool: a fixed number of goroutines reading jobs from a
//    channel. Each job includes the input index so we can place the output in
//    order without extra sorting.
// 4. How would you surface errors (e.g. n < 0, or n so large the string would
//    be huge) without aborting the whole batch?
//    Return `([]string, []error, error)` where the per-input errors slice is the
//    same length as the inputs, and the final error is only for systemic failure
//    (e.g. context cancelled). This keeps partial successes while reporting which
//    indexes failed.

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
)

// The interviewers hand you a small program built around a single goroutine.
// This is the exercise the briefing says they will *almost certainly* ask, so
// expect them to grow it live, one question at a time, while you narrate.
//
// 1. `sequence` is a goroutine that receives ONE integer and returns ONE string
//    (the FizzBuzz sequence for 1..n). Modify the code so it accepts MULTIPLE
//    integers and returns MULTIPLE strings — one result string per input, in the
//    same order as the input.
//
// 2. Right now the calls happen one at a time. Modify the code to handle
//    MULTIPLE SIMULTANEOUS calls to the goroutine. Use a sync.WaitGroup to wait
//    for all of them, and a sync.Mutex (or channels / atomics) to avoid a data
//    race while collecting the results. Running `go run -race .` must be clean.
//
// 3. Assume the list of integers can be very large (millions). Spawning one
//    goroutine per input exhausts memory and the scheduler. Bound the
//    concurrency with a worker pool of a fixed size, while still returning the
//    results in input order.
//
// 4. How would you surface errors (e.g. n < 0, or n so large the string would
//    be huge) without aborting the whole batch?

// sequence writes the FizzBuzz sequence for 1..n and sends the single result
// string on out.
func sequence(n int, out chan<- string) {
	var b strings.Builder
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			b.WriteString("FizzBuzz ")
		case i%3 == 0:
			b.WriteString("Fizz ")
		case i%5 == 0:
			b.WriteString("Buzz ")
		default:
			fmt.Fprintf(&b, "%d ", i)
		}
	}
	out <- strings.TrimSpace(b.String())
}

type seqJob struct {
	Idx int
	N   int
}

type seqResult struct {
	Idx int
	Str string
	Err error
}

func runSequences(ctx context.Context, input []int, workers int) ([]string, []error, error) {
	const maxN = 1_000_000 // answer 4: avoid accidental gigantic output in interview
	if workers <= 0 {
		workers = 1
	}

	results := make([]string, len(input))  // answer 1/2: order-preserving slots
	perErr := make([]error, len(input))    // answer 4: per-input error slot
	jobs := make(chan seqJob, workers*4)   // answer 3: bounded queue
	out := make(chan seqResult, workers*4) // result fan-in

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case job, ok := <-jobs:
					if !ok {
						return
					}
					if job.N < 0 {
						out <- seqResult{Idx: job.Idx, Err: fmt.Errorf("n must be >= 0, got %d", job.N)}
						continue
					}
					if job.N > maxN {
						out <- seqResult{Idx: job.Idx, Err: fmt.Errorf("n too large (%d), max %d", job.N, maxN)}
						continue
					}

					// answer 1: per-job goroutine primitive stays the same.
					ch := make(chan string, 1)
					sequence(job.N, ch)
					out <- seqResult{Idx: job.Idx, Str: <-ch}
				}
			}
		}()
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	for i, n := range input {
		select {
		case <-ctx.Done():
			close(jobs)
			return results, perErr, ctx.Err()
		case jobs <- seqJob{Idx: i, N: n}:
		}
	}
	close(jobs)

	for res := range out {
		if res.Idx < 0 || res.Idx >= len(results) {
			return nil, nil, errors.New("internal: result index out of range")
		}
		if res.Err != nil {
			perErr[res.Idx] = res.Err
			continue
		}
		results[res.Idx] = res.Str
	}

	return results, perErr, nil
}

func main() {
	// Sample data (paste from the meeting chat):
	input := []int{5, 15, 3, -1}

	// Single call, one sequence:
	// out := make(chan string)
	// go sequence(input[0], out)
	// fmt.Println(<-out)

	// TODO (questions 1–3): process ALL of `input` and print one line per int,
	// in input order, running the work concurrently with a bounded worker pool.
	results, errs, err := runSequences(context.Background(), input, 4)
	if err != nil {
		fmt.Printf("batch error: %v\n", err)
		return
	}
	for i := range input {
		if errs[i] != nil {
			fmt.Printf("%d => ERROR: %v\n", input[i], errs[i])
			continue
		}
		fmt.Printf("%d => %s\n", input[i], results[i])
	}
}

// How to run and test:
//   go run .
//   go run -race .
//
//   go test ./sisu-interview-prep/task-03-multi-sequence-goroutines-solution -v
//   go test ./sisu-interview-prep/task-03-multi-sequence-goroutines-solution -bench=. -benchmem
