# accept interfaces return concrete

## Live interview task
Demonstrate the Go proverb: accept interfaces, return concrete types (structs).

## Concepts covered
- interface design
- API boundaries
- concrete return types

## Candidate solution

```go
package main

import (
    "fmt"
    "io"
    "strings"
)

// Consumer accepts the minimal interface.
func digest(r io.Reader) (string, error) {
    b, err := io.ReadAll(r)
    if err != nil {
        return "", err
    }
    return string(b), nil
}

// Producer returns concrete *strings.Reader — caller can still pass as io.Reader.
func newInput() *strings.Reader {
    return strings.NewReader("payload")
}

func main() {
    s, _ := digest(newInput())
    fmt.Println(s)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Accept interfaces**: function params use `io.Reader`, `context.Context` — caller flexibility.
- **Return concrete**: return `*Client`, `*Server`, not `ClientInterface` — callers get full API, easier to extend.
- Exception: return interface when multiple implementations are intentional (`error`, `fmt.Stringer`).
- Define interfaces in **consumer** package, not producer — "interface segregation".

## Q&A

**Q: Who defines the interface?**  
A: The package that **uses** the dependency (e.g. `storage` interface in service, not in DB driver).

**Q: Returning interface from `New()`?**  
A: Sometimes for hiding implementation — trade-off: harder to access concrete methods.

**Q: Mocking?**  
A: Small interface in test package or generated mock — still accept interface at boundary.

**Q: Real stdlib example?**  
A: `json.Marshal` accepts `any`; returns `[]byte`.

**Q: Anti-pattern?**  
A: `type BigInterface` with 20 methods in shared package — forces fat mocks.
