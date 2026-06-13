# labeled break out of nested loop

## Live interview task
Search a matrix and break out of two nested loops when a target is found.

## Concepts covered
- labeled statements
- break
- nested loops

## Candidate solution

```go
package main

import "fmt"

func find(matrix [][]int, target int) (row, col int, ok bool) {
outer:
    for i, r := range matrix {
        for j, v := range r {
            if v == target {
                row, col, ok = i, j, true
                break outer
            }
        }
    }
    return
}

func main() { fmt.Println(find([][]int{{1, 2}, {3, 4}}, 3)) }
```

## Run

```bash
go run .
```

## Expected output

```
1 0 true
```

## Alternatives (no labels)

```go
func findFunc(matrix [][]int, target int) (int, int, bool) {
    for i, r := range matrix {
        if j, ok := findInRow(r, target); ok {
            return i, j, true
        }
    }
    return 0, 0, false
}
```

Extracting inner search to a function with `return` is often clearer than labels.

## Interview notes / pitfalls
- `break` without label exits only the innermost loop.
- `break outer` exits the labeled `for` — not `switch` unless labeled.
- `continue outer` skips to the next iteration of the outer loop.
- Labels are rare in production Go — prefer helper functions or `goto` only in generated/parser code.

## Q&A

**Q: Can you label a `switch`?**  
A: Yes — `break label` can exit an outer switch or loop; rarely used.

**Q: Complexity?**  
A: O(rows × cols) worst case; O(1) extra space.

**Q: Edge cases?**  
A: Empty matrix, ragged rows, duplicate targets (first match wins), target not found → `ok == false`.

**Q: When would an interviewer reject labels?**  
A: If readability suffers — show you know the helper-function refactor.
