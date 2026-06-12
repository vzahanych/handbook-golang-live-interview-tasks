# parallel word count

## Live interview task
Count words in many text chunks concurrently then merge maps.

## Concepts covered
- parallel maps
- map merge
- data race avoidance

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
        i, text := i, text
        wg.Add(1)
        go func(){ defer wg.Done(); m := make(map[string]int); for _, w := range strings.Fields(strings.ToLower(text)) { m[w]++ }; parts[i] = m }()
    }
    wg.Wait()
    out := make(map[string]int)
    for _, m := range parts { for k, v := range m { out[k] += v } }
    return out
}

func main() { fmt.Println(WordCount([]string{"go go", "maps go"})) }
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
