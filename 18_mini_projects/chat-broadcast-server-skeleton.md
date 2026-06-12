# chat broadcast server skeleton

## Live interview task
Implement the core of a TCP chat server broadcaster.

## Concepts covered
- TCP
- channels
- broadcast pattern

## Candidate solution

```go
package main

import (
    "bufio"
    "fmt"
    "net"
)

type client chan<- string

func broadcaster(join <-chan client, leave <-chan client, messages <-chan string) {
    clients := map[client]bool{}
    for {
        select {
        case c := <-join: clients[c] = true
        case c := <-leave: delete(clients, c)
        case msg := <-messages:
            for c := range clients { c <- msg }
        }
    }
}

func main() { fmt.Println("wire broadcaster to net.Listen; use bufio.Scanner(conn) per client", bufio.ScanLines, net.IPv4len) }
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
