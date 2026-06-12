# ordered min max constraint

## Live interview task
Define an Ordered constraint and implement generic Min and Max.

## Concepts covered
- type constraints
- underlying type terms
- ~T

## Candidate solution

```go
package main

import "fmt"

type Ordered interface { ~int | ~int64 | ~float64 | ~string }

func Min[T Ordered](a, b T) T { if a < b { return a }; return b }
func Max[T Ordered](a, b T) T { if a > b { return a }; return b }

type Age int

func main() { fmt.Println(Min(Age(10), Age(20)), Max("go", "java")) }
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
