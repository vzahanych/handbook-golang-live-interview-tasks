# interface small storage

## Live interview task
Write code against a `Storage` interface and inject an in-memory implementation.

## Concepts covered
- interfaces
- implicit implementation
- dependency injection

## Candidate solution

```go
package main

import "fmt"

type Storage interface {
    Put(k, v string)
    Get(k string) (string, bool)
}

type MemStore map[string]string

func (m MemStore) Put(k, v string)           { m[k] = v }
func (m MemStore) Get(k string) (string, bool) { v, ok := m[k]; return v, ok }

func saveAndLoad(s Storage) string {
    s.Put("lang", "go")
    v, _ := s.Get("lang")
    return v
}

func main() {
    fmt.Println(saveAndLoad(make(MemStore))) // go
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Go interfaces satisfied **implicitly** — no `implements` keyword.
- Keep interfaces **small** (ISP) — 1–3 methods ideal for `Storage`, `Reader`, `Writer`.
- Accept interface in function params; return concrete type unless abstraction needed.
- `MemStore` is a map type with methods — nil map panics on Put; `make(MemStore)` required.

## Q&A

**Q: Why interface at function boundary?**  
A: Test with fake/mock; swap Redis/Postgres without changing `saveAndLoad`.

**Q: Pointer vs value receiver on MemStore?**  
A: Map is reference type — value receiver still mutates backing map; pointer also fine.

**Q: Nil interface param?**  
A: Calling method panics — guard or document non-nil requirement.

**Q: Interface size?**  
A: Two words (type, data) — passing interface copies those words.

**Q: Testing?**  
A: `type fakeStore struct{}` with map inside implementing `Storage`.
