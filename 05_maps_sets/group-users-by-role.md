# group users by role

## Live interview task
Group records by key using a map from key to slice.

## Concepts covered
- maps
- slices
- append to nil slice

## Candidate solution

```go
package main

import "fmt"

type User struct{ Name, Role string }

func groupByRole(users []User) map[string][]User {
    out := make(map[string][]User)
    for _, u := range users {
        out[u.Role] = append(out[u.Role], u)
    }
    return out
}

func main() {
    fmt.Println(groupByRole([]User{{"Ann","admin"},{"Bob","user"},{"Cat","admin"}}))
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
