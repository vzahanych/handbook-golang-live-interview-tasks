# nil interface vs interface holding nil pointer

## Live interview task
Explain why an interface value holding a nil pointer is not equal to `nil`.

## Concepts covered
- interface representation
- nil
- dynamic type
- (type, value) pair

## Candidate solution

```go
package main

import "fmt"

type Reader struct{}
func (*Reader) Read() {}

func main() {
    var p *Reader = nil
    var x any = p
    fmt.Println(p == nil) // true
    fmt.Println(x == nil) // false — interface holds (type=*Reader, value=nil)
    fmt.Printf("%T %#v\n", x, x) // *main.Reader <nil>
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Interface value = (dynamic type, dynamic value). `nil` interface means **both** are unset.
- Typed nil pointer in interface: type is set → `== nil` is **false**.
- Common bug: `return err` where `err` is `(*MyError)(nil)` implementing `error` — caller's `err != nil` is true.
- Fix: return bare `nil` for no error, or use `var err error = nil` without typed nil.

## Q&A

**Q: How to check interface is "really" nil?**  
A: `v == nil` only works for untyped nil assignment. Use reflection sparingly; fix API to return untyped nil.

**Q: JSON `null` unmarshaled into interface?**  
A: Often becomes typed nil for pointer fields — same gotcha in handlers.

**Q: Safe return pattern?**  
A: `if p == nil { return nil }; return p` when `p` is a pointer type assigned to interface.

**Q: Why does Go work this way?**  
A: Dynamic dispatch needs a type at runtime; nil pointer still has type `*Reader` for method calls (which would panic if called).

**Q: Interview one-liner?**  
A: "An interface is nil only when it has no type and no value."
