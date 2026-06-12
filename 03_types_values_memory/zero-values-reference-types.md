# zero values reference types

## Live interview task
Print useful behavior of zero values for slice, map, channel, function and pointer types.

## Concepts covered
- zero values
- nil
- reference types

## Candidate solution

```go
package main

import "fmt"

func main() {
    var s []int
    var m map[string]int
    var ch chan int
    var f func()
    var p *int
    fmt.Println(s == nil, len(s), cap(s))
    fmt.Println(m == nil, len(m))
    fmt.Println(ch == nil, f == nil, p == nil)
    s = append(s, 1) // ok on nil slice
    // m["x"] = 1    // panic: assignment to entry in nil map
    fmt.Println(s)
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
