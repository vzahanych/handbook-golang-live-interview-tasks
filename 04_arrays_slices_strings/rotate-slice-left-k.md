# rotate slice left k

## Live interview task
Rotate a slice left by k positions using the three-reversal algorithm.

## Concepts covered
- slice expressions
- in-place algorithms

## Candidate solution

```go
package main

import "fmt"

func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func rotateLeft(s []int, k int) {
    if len(s) == 0 {
        return
    }
    k %= len(s)
    if k < 0 {
        k += len(s)
    }
    reverse(s[:k])
    reverse(s[k:])
    reverse(s)
}

func main() {
    s := []int{1, 2, 3, 4, 5}
    rotateLeft(s, 2)
    fmt.Println(s) // [3 4 5 1 2]
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Normalize `k` with `k %= len(s)` — rotating by `len` is identity.
- Three-reversal trick: reverse `[0:k)`, `[k:]`, then whole slice — O(n) time, O(1) space.
- Alternative: GCD block swap (Juggling) — same complexity, harder to code live.
- Negative k: treat as rotate right after normalization.

## Q&A

**Q: Why three reversals?**  
A: Classic proof: `reverse(A)+reverse(B)` reversed = `B+A` when rotating partition at k.

**Q: `k > len(s)`?**  
A: `k %= len(s)` handles it — rotating 7 on length 5 equals rotating 2.

**Q: Edge cases?**  
A: `k==0`, `k==len`, empty slice, `k` negative.

**Q: Rotate right by k?**  
A: `rotateLeft(s, len(s)-k)` or reverse order of the three steps.
