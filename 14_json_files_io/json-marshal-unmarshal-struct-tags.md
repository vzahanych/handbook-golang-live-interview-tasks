# json marshal unmarshal struct tags

## Live interview task
Marshal and unmarshal JSON with struct tags.

## Concepts covered
- encoding/json
- struct tags
- pointer receiver for Unmarshal

## Candidate solution

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Role string `json:"role,omitempty"`
}

func main() {
    data, _ := json.Marshal(User{ID: 1, Name: "Ada"})
    fmt.Println(string(data)) // {"id":1,"name":"Ada"}

    var u User
    _ = json.Unmarshal(data, &u) // must pass pointer
    fmt.Printf("%+v\n", u)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `Unmarshal` needs **pointer** to struct — otherwise cannot modify.
- Unexported fields ignored — only exported names serialize.
- `omitempty` skips zero values in **output** — pointer `nil` omitted, empty string omitted.
- `json:"-"` excludes field; `json:",string"` encodes numbers as strings.

## Q&A

**Q: Unknown JSON fields?**  
A: Ignored by default — `Decoder.DisallowUnknownFields()` for strict parsing.

**Q: `Marshal` error handling?**  
A: Channels, funcs, complex types fail — handle errors in production.

**Q: `time.Time`?**  
A: RFC3339 by default — custom `MarshalJSON` for other formats.

**Q: `json.RawMessage`?**  
A: Delay parsing nested JSON — useful for polymorphic APIs.

**Q: Complexity?**  
A: O(size of JSON) reflection-based — use code gen (easyjson) if hot path.
