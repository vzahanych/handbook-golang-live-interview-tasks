# context values request id

## Live interview task
Pass request-scoped values with context.Value.

## Concepts covered
- context values
- custom key type

## Candidate solution

```go
package main

import (
    "context"
    "fmt"
)

type key string
const requestID key = "requestID"

func log(ctx context.Context, msg string) { fmt.Println(ctx.Value(requestID), msg) }

func main() { ctx := context.WithValue(context.Background(), requestID, "req-123"); log(ctx, "started") }
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
