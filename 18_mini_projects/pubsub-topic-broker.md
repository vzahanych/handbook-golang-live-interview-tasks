# pubsub topic broker

## Live interview task
Build in-memory pub/sub broker with topics and non-blocking publish.

## Concepts covered
- pub/sub
- channels
- fan-out
- slow consumer handling

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

type Broker struct {
    mu   sync.RWMutex
    subs map[string][]chan string
}

func NewBroker() *Broker {
    return &Broker{subs: make(map[string][]chan string)}
}

func (b *Broker) Subscribe(topic string) <-chan string {
    ch := make(chan string, 8)
    b.mu.Lock()
    b.subs[topic] = append(b.subs[topic], ch)
    b.mu.Unlock()
    return ch
}

func (b *Broker) Publish(topic, msg string) {
    b.mu.RLock()
    subs := append([]chan string(nil), b.subs[topic]...)
    b.mu.RUnlock()

    for _, ch := range subs {
        select {
        case ch <- msg:
        default:
            // drop if subscriber slow — or block, or close subscriber
        }
    }
}

func main() {
    b := NewBroker()
    sub := b.Subscribe("go")
    b.Publish("go", "hello")
    fmt.Println(<-sub)
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Copy subscriber slice under `RLock` — Publish doesn't hold lock during send.
- `select` + `default` drops messages — document policy; chat may block instead.
- Unsubscribe: remove chan from slice + close — prevent leaks.
- Single broker goroutine pattern alternative — serialize all Subscribe/Publish.

## Q&A

**Q: Blocking publish?**  
A: Remove `default` — slow consumer blocks others if shared loop.

**Q: Buffer size 8?**  
A: Trade burst tolerance vs memory.

**Q: Wildcard topics?**  
A: Not in demo — match prefix in Publish.

**Q: Production?**  
A: NATS, Kafka, Redis pub/sub.

**Q: Complexity?**  
A: Publish O(subscribers on topic).
