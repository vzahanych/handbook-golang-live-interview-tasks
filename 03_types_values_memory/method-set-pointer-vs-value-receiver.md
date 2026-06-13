# method set pointer vs value receiver

## Live interview task
Show which methods belong to `T` and `*T` method sets.

## Concepts covered
- method sets
- pointer receiver
- addressability

## Candidate solution

```go
package main

import "fmt"

type Counter int

func (c Counter) Value() int { return int(c) }
func (c *Counter) Inc()     { *c++ }

type Valuer interface{ Value() int }
type Incer interface{ Inc() }

func main() {
    var c Counter
    var _ Valuer = c
    var _ Valuer = &c
    // var _ Incer = c  // compile error: Inc has pointer receiver
    var _ Incer = &c
    c.Inc() // ok: c is addressable, compiler rewrites to (&c).Inc()
    fmt.Println(c.Value())
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Method set of `T`: methods with receiver `T` (value).
- Method set of `*T`: methods with receiver `T` **or** `*T`.
- Value in interface: if only `*T` has the method, storing `T` value in interface fails.
- `c.Inc()` on value works only when `c` is **addressable** (variable, not map index or temp).
- Use pointer receiver when method mutates state or struct is large; be consistent within a type.

## Q&A

**Q: Why can't map elements call pointer-receiver mutators easily?**  
A: `m[k].Inc()` fails — map value is not addressable. Use `tmp := m[k]; tmp.Inc(); m[k] = tmp`.

**Q: Interface satisfaction with value receiver?**  
A: Both `T` and `*T` implement interface with only value-receiver methods.

**Q: `nil` pointer receiver?**  
A: Method runs but panics on dereference if it touches `*c` — valid for some nil-safe designs.

**Q: Compiler rewrite?**  
A: `c.Inc()` → `(&c).Inc()` only when addressable; otherwise compile error.

**Q: Rule of thumb?**  
A: All methods pointer receiver OR all value — mixing is OK but affects interface assignability.
