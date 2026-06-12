# table driven unit test

## Live interview task
Write a table-driven unit test for a pure function.

## Concepts covered
- testing
- table-driven tests
- t.Run

## Candidate solution

```go
// file: add.go
package calc

func Add(a, b int) int { return a + b }

// file: add_test.go
package calc

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct{ name string; a, b, want int }{
        {"positive", 1, 2, 3},
        {"negative", -1, -2, -3},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.want { t.Fatalf("got %d want %d", got, tt.want) }
        })
    }
}
```

## Run

```bash
go test ./...
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
