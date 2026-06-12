# type switch format any

## Live interview task
Use a type switch to format values of different dynamic types.

## Concepts covered
- interfaces
- type switches
- dynamic type

## Candidate solution

```go
package main

import "fmt"

func describe(v any) string {
    switch x := v.(type) {
    case nil:
        return "nil"
    case int:
        return fmt.Sprintf("int:%d", x)
    case string:
        return fmt.Sprintf("string:%q", x)
    case fmt.Stringer:
        return "stringer:" + x.String()
    default:
        return fmt.Sprintf("%T", x)
    }
}

func main() { fmt.Println(describe("go"), describe(42)) }
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
