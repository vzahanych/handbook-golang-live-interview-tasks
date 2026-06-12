# blank identifier for compile time interface check

## Live interview task
Use the blank identifier for a compile-time interface implementation assertion.

## Concepts covered
- blank identifier
- interfaces
- method sets

## Candidate solution

```go
package main

import "fmt"

type Stringer interface{ String() string }

type User struct{ Name string }
func (u User) String() string { return u.Name }

var _ Stringer = User{} // fails at compile time if User stops implementing Stringer

func main() { fmt.Println(User{"Ada"}) }
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
