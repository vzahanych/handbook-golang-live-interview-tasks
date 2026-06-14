# rotate slice left k

## Live interview task
Rotate a slice left by k positions using the three-reversal algorithm — move the first k elements to the end (e.g. `[1,2,3,4,5]` left by `2` → `[3,4,5,1,2]`).

## Concepts covered
- slice expressions
- in-place algorithms

## Candidate solution

```go
package main

import "fmt"

// reverse swaps elements in place so s reads backwards.
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// rotateLeft moves the first k elements to the end: [A|B] → [B|A].
// Uses three reversals — O(n) time, O(1) extra space.
//
// Split at k: A = s[:k], B = s[k:].
// Property: reverse( reverse(A) + reverse(B) ) = B + A
//
// Example: [1,2,3,4,5], k=2  (A=[1,2] B=[3,4,5])
//   reverse A     → [2,1,3,4,5]
//   reverse B     → [2,1,5,4,3]
//   reverse whole → [3,4,5,1,2]
func rotateLeft(s []int, k int) {
    if len(s) == 0 {
        return
    }
    k %= len(s) // rotating by len(s) is a full circle — no-op
    if k < 0 {
        k += len(s) // negative k → equivalent left rotation
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
