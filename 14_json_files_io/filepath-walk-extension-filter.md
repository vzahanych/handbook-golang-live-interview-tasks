# filepath walk extension filter

## Live interview task
Walk a directory and collect files with a chosen extension.

## Concepts covered
- filepath.WalkDir
- io/fs.DirEntry
- pruning walks

## Candidate solution

```go
package main

import (
    "fmt"
    "io/fs"
    "path/filepath"
)

func find(root, ext string) ([]string, error) {
    var out []string
    err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            return nil
        }
        if filepath.Ext(path) == ext {
            out = append(out, path)
        }
        return nil
    })
    return out, err
}

func main() {
    xs, _ := find(".", ".go")
    fmt.Println(len(xs), "go files")
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `WalkDir` preferred over deprecated `Walk` — fewer syscalls.
- Return `filepath.SkipDir` to skip subdirectories (e.g. `.git`, `vendor`).
- `filepath.Ext` includes dot — compare `".go"` not `"go"`.
- Symlinks: `WalkDir` does not follow by default into symlinked dirs.

## Q&A

**Q: Skip `node_modules`?**  
A: `if d.IsDir() && d.Name() == "node_modules" { return filepath.SkipDir }`.

**Q: `fs.WalkDir`?**  
A: Same callback style on `fs.FS` — embed for testable walks.

**Q: Complexity?**  
A: O(files + dirs) visited.

**Q: Parallel walk?**  
A: Collect dirs sequentially or use worker pool per subtree — advanced.

**Q: Edge cases?**  
A: Permission error on subdir — return err or skip with log.
