# stream json lines decoder

## Live interview task
Decode newline-delimited JSON from an io.Reader.

## Concepts covered
- json.Decoder
- streaming

## Candidate solution

```go
package main

   import (
       "encoding/json"
       "fmt"
       "strings"
   )

   type Event struct{ Type string `json:"type"` }

   func main() {
       dec := json.NewDecoder(strings.NewReader(`{"type":"a"}
{"type":"b"}
`))
       for dec.More() {
           var e Event
           if err := dec.Decode(&e); err != nil { panic(err) }
           fmt.Println(e.Type)
       }
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
