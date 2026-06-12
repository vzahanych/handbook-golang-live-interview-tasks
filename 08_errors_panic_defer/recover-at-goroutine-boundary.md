# recover at goroutine boundary

## Live interview task
Recover from a panic at a goroutine boundary and report it as an error.

## Concepts covered
- panic
- recover
- goroutines

## Candidate solution

```go
package main

import "fmt"

func safeGo(fn func()) <-chan error {
    done := make(chan error, 1)
    go func() {
        defer func() {
            if r := recover(); r != nil { done <- fmt.Errorf("panic: %v", r) } else { done <- nil }
        }()
        fn()
    }()
    return done
}

func main() { fmt.Println(<-safeGo(func(){ panic("boom") })) }
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
