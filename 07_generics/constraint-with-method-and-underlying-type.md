# constraint with method and underlying type

## Live interview task
Use a constraint that requires both an underlying type and a method.

## Concepts covered
- general interfaces
- type sets
- methods in constraints

## Candidate solution

```go
package main

import "fmt"

type ID int
func (id ID) String() string { return fmt.Sprintf("ID-%d", id) }

type PrintableInt interface { ~int; String() string }

func PrintPlusOne[T PrintableInt](v T) string { return fmt.Sprintf("%s next=%d", v.String(), int(v)+1) }

func main() { fmt.Println(PrintPlusOne(ID(41))) }
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
