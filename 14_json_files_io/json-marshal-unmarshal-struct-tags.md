# json marshal unmarshal struct tags

## Live interview task
Marshal and unmarshal JSON with struct tags.

## Concepts covered
- encoding/json
- struct tags

## Candidate solution

```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct { ID int `json:"id"`; Name string `json:"name"` }

func main() {
    data, _ := json.Marshal(User{1, "Ada"})
    fmt.Println(string(data))
    var u User
    json.Unmarshal(data, &u)
    fmt.Printf("%+v\n", u)
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
