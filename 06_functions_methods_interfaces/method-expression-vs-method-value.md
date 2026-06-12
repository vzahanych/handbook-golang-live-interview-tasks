# method expression vs method value

## Live interview task
Show the difference between method expressions and method values.

## Concepts covered
- method expressions
- method values

## Candidate solution

```go
package main

import "fmt"

type User struct{ Name string }
func (u User) Greet(prefix string) string { return prefix + " " + u.Name }

func main() {
    u := User{"Ada"}
    value := u.Greet           // receiver captured
    expr := User.Greet         // receiver is first argument
    fmt.Println(value("hi"))
    fmt.Println(expr(u, "hello"))
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
