# csv reader to structs

## Live interview task
Parse CSV records into structs.

## Concepts covered
- encoding/csv
- strconv

## Candidate solution

```go
package main

import (
    "encoding/csv"
    "fmt"
    "strconv"
    "strings"
)

type Product struct { Name string; Price int }

func parseCSV(s string) ([]Product, error) {
    r := csv.NewReader(strings.NewReader(s))
    rows, err := r.ReadAll(); if err != nil { return nil, err }
    out := make([]Product, 0, len(rows))
    for _, row := range rows[1:] { price, _ := strconv.Atoi(row[1]); out = append(out, Product{row[0], price}) }
    return out, nil
}

func main() { xs, _ := parseCSV("name,price\nbook,10\npen,2\n"); fmt.Println(xs) }
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
