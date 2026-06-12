# sync once lazy init

## Live interview task
Initialize a value exactly once with sync.Once.

## Concepts covered
- sync.Once
- lazy initialization

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

var once sync.Once
var config map[string]string

func Config() map[string]string {
    once.Do(func(){ config = map[string]string{"env":"dev"} })
    return config
}

func main() { fmt.Println(Config(), Config()) }
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
