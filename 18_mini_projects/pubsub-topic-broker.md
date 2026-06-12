# pubsub topic broker

## Live interview task
Build a simple in-memory pub/sub broker with topics.

## Concepts covered
- pub/sub
- channels
- non-blocking send

## Candidate solution

```go
package main

import (
    "fmt"
    "sync"
)

type Broker struct { mu sync.RWMutex; subs map[string][]chan string }
func NewBroker() *Broker { return &Broker{subs: make(map[string][]chan string)} }
func (b *Broker) Subscribe(topic string) <-chan string { ch := make(chan string, 8); b.mu.Lock(); b.subs[topic] = append(b.subs[topic], ch); b.mu.Unlock(); return ch }
func (b *Broker) Publish(topic, msg string) { b.mu.RLock(); subs := append([]chan string(nil), b.subs[topic]...); b.mu.RUnlock(); for _, ch := range subs { select { case ch <- msg: default: } } }

func main() { b := NewBroker(); sub := b.Subscribe("go"); b.Publish("go", "hello"); fmt.Println(<-sub) }
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
