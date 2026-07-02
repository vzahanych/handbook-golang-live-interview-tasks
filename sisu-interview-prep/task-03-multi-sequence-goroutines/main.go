package main

import (
	"fmt"
	"strings"
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

func main() {
	// Sample data (paste from the meeting chat):
	input := []int{5, 15, 3}

	// Single call, one sequence:
	out := make(chan string)
	go sequence(input[0], out)
	fmt.Println(<-out)

	// TODO (questions 1–3): process ALL of `input` and print one line per int,
	// in input order, running the work concurrently with a bounded worker pool.
	_ = input
}
