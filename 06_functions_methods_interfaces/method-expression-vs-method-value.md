# method expression vs method value

## Live interview task
Show the difference between method expressions and method values.

## Concepts covered
- method expressions
- method values
- bound receivers

## Candidate solution

```go
package main

import "fmt"

type User struct{ Name string }

func (u User) Greet(prefix string) string { return prefix + " " + u.Name }

func main() {
    u := User{"Ada"}
    value := u.Greet    // method value — receiver bound to u
    expr := User.Greet  // method expression — receiver is first arg

    fmt.Println(value("hi"))       // hi Ada
    fmt.Println(expr(u, "hello"))  // hello Ada
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Method value** `u.Greet`: `func(string) string` — `u` captured (copied for value receiver).
- **Method expression** `User.Greet`: `func(User, string) string` — pass receiver explicitly.
- Pointer receiver: `(&u).Greet` vs `(*User).Greet` — same distinction.
- Method values keep a copy of value receiver — mutations inside method on copy don't affect original `u` unless pointer receiver.

## Q&A

**Q: When use method expression?**  
A: Pass as callback where first arg should be supplied later — rare; often use closure instead.

**Q: `http.HandlerFunc` relation?**  
A: Adapter pattern — function type implementing interface.

**Q: Nil pointer method value?**  
A: `var u *User; f := u.Greet` — call panics if method dereferences receiver.

**Q: Complexity?**  
A: O(1) — syntactic sugar for calls.

**Q: Interview trick?**  
A: Ask type of `User.Greet` — `func(User, string) string`.
