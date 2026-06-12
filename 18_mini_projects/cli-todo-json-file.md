# cli todo json file

## Live interview task
Implement a tiny JSON-backed todo storage layer for a CLI.

## Concepts covered
- JSON
- files
- CLI storage

## Candidate solution

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Todo struct { Text string `json:"text"`; Done bool `json:"done"` }

func load(path string) ([]Todo, error) { b, err := os.ReadFile(path); if os.IsNotExist(err) { return nil, nil }; if err != nil { return nil, err }; var xs []Todo; return xs, json.Unmarshal(b, &xs) }
func save(path string, xs []Todo) error { b, err := json.MarshalIndent(xs, "", "  "); if err != nil { return err }; return os.WriteFile(path, b, 0644) }

func main() { xs, _ := load("todo.json"); xs = append(xs, Todo{Text:"learn go"}); save("todo.json", xs); fmt.Println(xs) }
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
