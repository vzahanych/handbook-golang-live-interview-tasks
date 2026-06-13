# exported identifiers and package api

## Live interview task
Demonstrate exported versus unexported identifiers in a package API.

## Concepts covered
- exported identifiers
- packages
- methods

## Candidate solution

```go
// file: counter/counter.go
package counter

type Counter struct { n int } // Counter exported; n is package-private

func New() *Counter { return &Counter{} }
func (c *Counter) Inc() { c.n++ }
func (c *Counter) Value() int { return c.n }

// file: main.go
package main

import (
    "fmt"
    "example/counter"
)

func main() {
    c := counter.New()
    c.Inc()
    fmt.Println(c.Value()) // 1
    // c.n++  // compile error: n is unexported
}
```

## Run

```bash
go mod init example && go run .
```

## Interview notes / pitfalls
- Exported = name starts with **uppercase** letter (Unicode uppercase, not just ASCII — but stick to ASCII in practice).
- Unexported fields are inaccessible outside the package — encapsulation without `private` keyword.
- Constructor `New()` pattern exposes a controlled API while hiding representation.
- Cross-package embedding: embedded exported types promote exported methods only.

## Q&A

**Q: Can another package access `counter.Counter.n`?**  
A: No — lowercase `n` is unexported. Only methods in `counter` can touch it.

**Q: Why return `*Counter` from `New()`?**  
A: So `Inc` can mutate state with a pointer receiver; callers share one instance.

**Q: What is the "accept interfaces, return structs" rule?**  
A: Functions should accept small interfaces as parameters but return concrete types unless abstraction is required — keeps APIs flexible for callers.

**Q: How would you test unexported behavior?**  
A: Test in `counter_test` with `package counter` (white-box) or test only through exported API (black-box `package counter_test`).

**Q: Edge cases?**  
A: JSON unmarshaling into unexported fields fails unless you use custom `UnmarshalJSON` or exported fields with tags.
