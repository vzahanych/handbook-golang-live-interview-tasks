# struct padding order fields

## Live interview task
Compare struct sizes when fields are ordered differently.

## Concepts covered
- struct padding
- unsafe.Sizeof
- alignment

## Candidate solution

```go
package main

import (
    "fmt"
    "unsafe"
)

type Bad struct {
    A bool  // 1 byte + 7 padding
    B int64 // 8
    C bool  // 1 + 7 padding
}

type Better struct {
    B int64 // 8
    A bool  // 1
    C bool  // 1 + 6 padding
}

func main() {
    fmt.Println(unsafe.Sizeof(Bad{}), unsafe.Sizeof(Better{})) // 24 16 (typical amd64)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Go inserts padding so each field aligns to its type's requirement (e.g. `int64` on 8-byte boundary).
- Order fields largest → smallest to minimize padding in hot structs.
- `unsafe.Sizeof` includes padding; `unsafe.Alignof` reports alignment requirement.
- For serialization wire format, use explicit layout or code generation — do not rely on padding for API.

## Q&A

**Q: Does reordering break ABI?**  
A: Within a package it's a source change; across shared memory / C interop, layout matters.

**Q: `struct{}` size?**  
A: 0 bytes — useful as map set values.

**Q: When does this matter in interviews?**  
A: Large slices of structs (cache lines), embedded systems, or explaining why `bool` between `int64`s wastes space.

**Q: Tooling?**  
A: `fieldalignment -fix` (golang.org/x/tools) suggests reordering.

**Q: Complexity?**  
A: O(1) per struct; savings are memory bandwidth at scale.
