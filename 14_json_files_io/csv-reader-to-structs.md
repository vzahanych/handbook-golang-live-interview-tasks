# csv reader to structs

## Live interview task
Parse CSV records into structs with header row.

## Concepts covered
- encoding/csv
- strconv
- error handling

## Candidate solution

```go
package main

import (
    "encoding/csv"
    "fmt"
    "strconv"
    "strings"
)

type Product struct {
    Name  string
    Price int
}

func parseCSV(s string) ([]Product, error) {
    r := csv.NewReader(strings.NewReader(s))
    rows, err := r.ReadAll()
    if err != nil {
        return nil, err
    }
    if len(rows) < 2 {
        return nil, nil
    }

    out := make([]Product, 0, len(rows)-1)
    for _, row := range rows[1:] {
        if len(row) < 2 {
            continue
        }
        price, err := strconv.Atoi(row[1])
        if err != nil {
            return nil, err
        }
        out = append(out, Product{Name: row[0], Price: price})
    }
    return out, nil
}

func main() {
    xs, _ := parseCSV("name,price\nbook,10\npen,2\n")
    fmt.Println(xs)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- `ReadAll` loads entire CSV — use `Read()` in loop for large files.
- `r.FieldsPerRecord = -1` allows variable columns — or enforce count.
- Handle quoted commas — csv.Reader handles RFC 4180 quoting.
- Skip header with `rows[1:]` or read first row as column names map.

## Q&A

**Q: Lazy CSV?**  
A: `for { row, err := r.Read(); ... }`.

**Q: Different delimiter?**  
A: `r.Comma = ';'`.

**Q: Empty fields?**  
A: `""` in row — `Atoi` fails if price empty — validate.

**Q: Struct tags?**  
A: Libraries map by header name — manual loop fine for interviews.

**Q: Complexity?**  
A: O(rows * cols).
