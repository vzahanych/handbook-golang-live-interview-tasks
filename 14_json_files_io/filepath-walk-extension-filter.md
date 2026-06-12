# filepath walk extension filter

## Live interview task
Walk a directory and collect files with a chosen extension.

## Concepts covered
- filepath.WalkDir
- interfaces

## Candidate solution

```go
package main

import (
    "fmt"
    "path/filepath"
)

func find(root, ext string) ([]string, error) {
    var out []string
    err := filepath.WalkDir(root, func(path string, d interface{ IsDir() bool }, err error) error {
        if err != nil { return err }
        if !d.IsDir() && filepath.Ext(path) == ext { out = append(out, path) }
        return nil
    })
    return out, err
}

func main() { xs, _ := find(".", ".go"); fmt.Println(xs) }
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- The exact signature uses fs.DirEntry; interface form shown highlights only IsDir.

## Follow-up questions
- What is the time and space complexity?
- What edge cases would you test?
- How would you make this production-ready?
