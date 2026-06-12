# race detector shared map test

## Live interview task
Write a test that should be run with the race detector to catch shared map races.

## Concepts covered
- race detector
- maps
- tests

## Candidate solution

```go
package race

import (
    "sync"
    "testing"
)

func TestSharedMapRace(t *testing.T) {
    m := map[int]int{}
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        i := i
        wg.Add(1)
        go func(){ defer wg.Done(); m[i] = i }()
    }
    wg.Wait()
}
```

## Run

```bash
go test -race
```

## Interview notes / pitfalls
- None specific; discuss edge cases and complexity.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
