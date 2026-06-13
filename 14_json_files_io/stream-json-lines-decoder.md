# stream json lines decoder

## Live interview task
Decode newline-delimited JSON (JSONL/NDJSON) from an `io.Reader` without loading entire file.

## Concepts covered
- json.Decoder
- streaming
- NDJSON

## Candidate solution

```go
package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

type Event struct {
    Type string `json:"type"`
}

func main() {
    input := `{"type":"a"}
{"type":"b"}
`
    dec := json.NewDecoder(strings.NewReader(input))
    for dec.More() {
        var e Event
        if err := dec.Decode(&e); err != nil {
            panic(err)
        }
        fmt.Println(e.Type)
    }
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `Decoder.Decode` reads one JSON value — newline between objects optional for stream.
- `dec.More()` peeks if another value in stream — use in loop until false.
- Memory O(1) per record vs `Unmarshal` entire file.
- Blank lines may cause errors — trim or skip empty decodes.

## Q&A

**Q: vs `Scanner` per line + `Unmarshal`?**  
A: Equivalent for JSONL; Decoder handles whitespace between values.

**Q: Large file?**  
A: `os.Open` + `json.NewDecoder(file)` — constant memory per row.

**Q: Encoder for JSONL?**  
A: `json.NewEncoder(w).Encode(v)` adds newline per value.

**Q: Invalid line mid-file?**  
A: Return error with line offset — wrap with line counter.

**Q: Complexity?**  
A: O(bytes) total, O(1) peak per record.
