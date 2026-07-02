package main

// Two tests + a benchmark.
//   go test ./sisu-interview-prep/task-03-multi-sequence-goroutines-solution -v
//   go test ./sisu-interview-prep/task-03-multi-sequence-goroutines-solution -bench=. -benchmem

import (
	"context"
	"strings"
	"testing"
)

func TestRunSequences_OrderAndErrors(t *testing.T) {
	input := []int{5, 3, -1}
	got, perErr, err := runSequences(context.Background(), input, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != len(input) || len(perErr) != len(input) {
		t.Fatalf("lens mismatch")
	}
	if perErr[2] == nil {
		t.Fatalf("expected error for -1")
	}
	if got[0] != "1 2 Fizz 4 Buzz" {
		t.Fatalf("got[0]=%q", got[0])
	}
	if got[1] != "1 2 Fizz" {
		t.Fatalf("got[1]=%q", got[1])
	}
}

func TestSequence_Single(t *testing.T) {
	ch := make(chan string, 1)
	sequence(15, ch)
	if got := <-ch; !strings.Contains(got, "FizzBuzz") {
		t.Fatalf("expected FizzBuzz in %q", got)
	}
}

func BenchmarkRunSequences(b *testing.B) {
	input := make([]int, 0, 10_000)
	for i := 0; i < cap(input); i++ {
		input = append(input, 100+(i%50))
	}
	b.ReportAllocs()
	for b.Loop() {
		_, _, _ = runSequences(context.Background(), input, 8)
	}
}
