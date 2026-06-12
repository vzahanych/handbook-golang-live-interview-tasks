# interface small storage

## Live interview task
Write code against a Storage interface and inject an in-memory implementation.

## Concepts covered
- interfaces
- method sets
- dependency injection

## Candidate solution

```go
package main

import "fmt"

type Storage interface { Put(string, string); Get(string) (string, bool) }

type MemStore map[string]string
func (m MemStore) Put(k, v string) { m[k] = v }
func (m MemStore) Get(k string) (string, bool) { v, ok := m[k]; return v, ok }

func saveAndLoad(s Storage) string {
    s.Put("lang", "go")
    v, _ := s.Get("lang")
    return v
}

func main() { fmt.Println(saveAndLoad(make(MemStore))) }
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
