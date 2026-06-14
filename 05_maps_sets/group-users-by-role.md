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

// groupByRole builds map[role] → users with that role, preserving input order within each group.
//
// Example: Ann(admin), Bob(user), Cat(admin) →
//   "admin": [Ann, Cat]
//   "user":  [Bob]
func groupByRole(users []User) map[string][]User {
    out := make(map[string][]User)
    for _, u := range users {
        // out[u.Role] is nil on first sight of a role — append still works
        out[u.Role] = append(out[u.Role], u)
    }
    return out
}

func main() {
    fmt.Println(groupByRole([]User{
        {"Ann", "admin"},
        {"Bob", "user"},
        {"Cat", "admin"},
    }))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `append` to `out[u.Role]` when key missing — `nil` slice append works (allocates on first append).
- Stores **copies** of `User` structs in each group slice.
- For pointers: `[]*User` avoids copy but shares mutable state.
- Generic version: `groupBy[T any, K comparable](items []T, key func(T) K) map[K][]T`.

## Q&A

**Q: Complexity?**  
A: O(n) time, O(n) space for stored users.

**Q: Preserve input order within group?**  
A: Yes — append preserves encounter order.

**Q: Empty role string?**  
A: Valid key — users with `Role: ""` group together.

**Q: Preallocate slice per key?**  
A: Hard without knowing sizes — optional second pass count then fill.

**Q: sql/database pattern?**  
A: Same as `GROUP BY` in SQL — very common in service layers.
