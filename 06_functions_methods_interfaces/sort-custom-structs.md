# sort custom structs

## Live interview task
Sort structs by multiple fields using sort.Slice.

## Concepts covered
- function literals
- sort.Slice
- closures

## Candidate solution

```go
package main

import (
    "fmt"
    "sort"
)

type User struct{ Name string; Age int }

func main() {
    users := []User{{"Bob", 30}, {"Ann", 30}, {"Cat", 20}}
    sort.Slice(users, func(i, j int) bool {
        if users[i].Age != users[j].Age { return users[i].Age < users[j].Age }
        return users[i].Name < users[j].Name
    })
    fmt.Println(users)
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
