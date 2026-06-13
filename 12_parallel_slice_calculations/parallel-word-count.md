# parallel word count

## Live interview task
Count words in many text chunks concurrently, then merge maps.

## Concepts covered
- parallel maps
- map merge
- race-free aggregation

## Candidate solution

```go
package main

import (
    "fmt"
    "strings"
    "sync"
)

func WordCount(chunks []string) map[string]int {
    parts := make([]map[string]int, len(chunks))
    var wg sync.WaitGroup

    for i, text := range chunks {
        wg.Add(1)
        go func(i int, text string) {
            defer wg.Done()
            m := make(map[string]int)
            for _, w := range strings.Fields(strings.ToLower(text)) {
                m[w]++
            }
            parts[i] = m
        }(i, text)
    }
    wg.Wait()

    out := make(map[string]int)
    for _, m := range parts {
        for k, v := range m {
            out[k] += v
        }
    }
    return out
}

func main() {
    fmt.Println(WordCount([]string{"go go", "maps go"}))
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- **Never** concurrent write same map — per-goroutine local map, merge single-threaded.
- Merge phase can be parallel by key shards if huge — usually sequential merge is fine.
- Same pattern as Hadoop map-reduce map + reduce locally.

## Q&A

**Q: Complexity?**  
A: O(total words) time; O(unique words) space.

**Q: Shard merge?**  
A: Hash key to worker bucket during merge — parallel reduce.

**Q: sync.Map for merge?**  
A: Possible but often slower than local maps + merge.

**Q: Empty chunks?**  
A: Empty map in `parts[i]` — merge skips.

**Q: Production?**  
A: Stream files per chunk; bound goroutine count with worker pool.
