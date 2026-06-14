# chat broadcast server skeleton

## Live interview task
Implement TCP chat broadcaster — central goroutine fans out messages to clients.

## Concepts covered
- TCP
- channels
- fan-out broadcast
- select loop

## Candidate solution

```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
)

type client chan<- string

func broadcaster(incoming chan client, leaving chan client, messages chan string) {
    clients := map[client]bool{}
    for {
        select {
        case c := <-incoming:
            clients[c] = true
        case c := <-leaving:
            delete(clients, c)
            close(c)
        case msg := <-messages:
            for c := range clients {
                select {
                case c <- msg:
                default:
                    // slow client — skip or disconnect
                }
            }
        }
    }
}

func handleConn(conn net.Conn, incoming, leaving chan client, messages chan string) {
    ch := make(chan string, 8)
    incoming <- ch
    defer func() { leaving <- ch }()

    go func() {
        scanner := bufio.NewScanner(conn)
        for scanner.Scan() {
            messages <- scanner.Text()
        }
    }()

    for msg := range ch {
        fmt.Fprintln(conn, msg)
    }
}

func main() {
    incoming := make(chan client)
    leaving := make(chan client)
    messages := make(chan string)
    go broadcaster(incoming, leaving, messages)

    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }
    for {
        conn, err := ln.Accept()
        if err != nil {
            continue
        }
        go handleConn(conn, incoming, leaving, messages)
    }
}
```

## Run

Runnable version: server in [server/](server/main.go), client in [client/](client/main.go).

```bash
# terminal 1 — start the server
go run ./18_mini_projects/chat-broadcast-server/server

# terminals 2+ — one per chat participant (built-in Go client)
go run ./18_mini_projects/chat-broadcast-server/client
# ...or just use netcat
nc localhost 8080
```

## Interview notes / pitfalls
- **One** broadcaster goroutine owns `clients` map — no mutex needed.
- Client send channel `chan<- string` — type enforces send-only from broadcaster.
- Slow client: non-blocking send or dedicated write goroutine per conn.
- Based on Go blog chat server pattern — classic interview architecture question.

## Q&A

**Q: Why channels not mutex for clients?**  
A: Single goroutine serializes mutations — avoid races.

**Q: Deadlock risk?**  
A: Blocking `c <- msg` stalls all — use select default or buffered ch.

**Q: vs WebSocket?**  
A: Same fan-out; WS for browsers.

**Q: Scale?**  
A: Shard rooms, Redis pub/sub across nodes.

**Q: Complexity?**  
A: Broadcast O(clients) per message.
