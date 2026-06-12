# assignability with underlying types

## Live interview task
Show assignability and conversion rules for named types with identical underlying types.

## Concepts covered
- assignability
- underlying types
- conversions

## Candidate solution

```go
package main

import "fmt"

type Meters int
type Seconds int

func main() {
    var m Meters = 10
    var s Seconds = 2
    // m = s // compile error: different named types
    m = Meters(s)
    fmt.Println(m)
}
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
