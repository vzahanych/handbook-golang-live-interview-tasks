# cli todo json file

## Live interview task
Implement JSON-backed todo persistence for a CLI tool.

## Concepts covered
- JSON file storage
- os.ReadFile/WriteFile
- atomic write pattern

## Candidate solution

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Todo struct {
    Text string `json:"text"`
    Done bool   `json:"done"`
}

func load(path string) ([]Todo, error) {
    b, err := os.ReadFile(path)
    if os.IsNotExist(err) {
        return nil, nil
    }
    if err != nil {
        return nil, err
    }
    var xs []Todo
    return xs, json.Unmarshal(b, &xs)
}

func save(path string, xs []Todo) error {
    b, err := json.MarshalIndent(xs, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(path, b, 0644)
}

func main() {
    xs, _ := load("todo.json")
    xs = append(xs, Todo{Text: "learn go"})
    _ = save("todo.json", xs)
    fmt.Println(xs)
}
```

## Run

Runnable version lives in [cli-todo/](cli-todo/main.go). It adds real
`add`/`list`/`done` subcommands and an atomic save (temp file + `os.Rename`).

```bash
go run ./18_mini_projects/cli-todo add "learn go"
go run ./18_mini_projects/cli-todo add "write tests"
go run ./18_mini_projects/cli-todo done 1
go run ./18_mini_projects/cli-todo list
cat todo.json   # written in the current directory (override with -file)
```

## Interview notes / pitfalls
- `IsNotExist` → empty list — first run friendly.
- Crash during `WriteFile` may corrupt — write temp + `Rename` for atomicity.
- Concurrent CLI processes — file lock or single writer.
- Extend with `cobra` subcommands: add, list, done.

## Q&A

**Q: Atomic save?**  
A: `WriteFile(path.tmp)` then `os.Rename(tmp, path)`.

**Q: Permissions `0644`?**  
A: User rw, group/other read — adjust for secrets.

**Q: Large todo list?**  
A: SQLite or line-delimited JSON stream.

**Q: Test?**  
A: `t.TempDir()` + load/save roundtrip.

**Q: Complexity?**  
A: O(n) file size per save — rewrite whole file.
