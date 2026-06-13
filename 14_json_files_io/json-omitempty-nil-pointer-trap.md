# json omitempty nil pointer trap

## Live interview task
Show how `omitempty` treats nil pointers vs empty strings in JSON output.

## Concepts covered
- json omitempty
- pointers in structs
- API design

## Candidate solution

```go
package main

import (
    "encoding/json"
    "fmt"
)

type Profile struct {
    Bio  *string `json:"bio,omitempty"`
    City string  `json:"city,omitempty"`
}

func main() {
    empty := ""
    withBio := Profile{Bio: &empty, City: ""}

    b1, _ := json.Marshal(Profile{})
    b2, _ := json.Marshal(withBio)

    fmt.Println(string(b1)) // {}
    fmt.Println(string(b2)) // {"bio":""} — pointer non-nil, city omitted
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `omitempty` omits **zero values** — `nil` pointer omitted, non-nil pointer to `""` **included** as `""`.
- Distinguish JSON `null`, missing key, and `""` — use pointers for optional fields.
- `*int` with value 0: non-nil pointer encodes `0` — not omitted.
- PATCH APIs: pointer = field present; omitempty alone loses "set to empty" semantics.

## Q&A

**Q: Emit `null`?**  
A: Non-nil pointer required; `json:"bio"` without omitempty on pointer field.

**Q: `sql.NullString` pattern?**  
A: Custom `MarshalJSON` for valid/invalid + value.

**Q: Unmarshal missing vs null?**  
A: Both often leave pointer nil — use `json.RawMessage` for strict diff.

**Q: Interview fix for optional bio?**  
A: `*string` — nil = absent, `&""` = explicitly empty.

**Q: Complexity?**  
A: O(fields) per marshal.
