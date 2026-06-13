# race detector shared map test

## Live interview task
Demonstrate a data race on a shared map that `go test -race` catches.

## Concepts covered
- race detector
- concurrent map writes
- fixing with mutex

## Buggy test (race)

```go
func TestSharedMapRace(t *testing.T) {
    m := map[int]int{}
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            m[i] = i // concurrent map write — race
        }(i)
    }
    wg.Wait()
}
```

## Fixed pattern (for interview discussion)

```go
type Safe struct {
    mu sync.Mutex
    m  map[int]int
}

func (s *Safe) Set(k, v int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m[k] = v
}
```

## Run

```bash
go test -race ./...
```

## Interview notes / pitfalls
- Race detector instruments memory accesses — ~5-10x slowdown, use in CI not every edit.
- Concurrent map read+write panics at runtime even without -race — race detector finds subtler races.
- Test documents **intentional** race for teaching — mark fixed version in same package test file separately.

## Q&A

**Q: What is a race?**  
A: Two goroutines access same memory, one write, without synchronization.

**Q: False positives?**  
A: Rare — usually real bugs.

**Q: CI?**  
A: `go test -race ./...` on Linux/macOS — not all platforms.

**Q: Atomic enough?**  
A: For single int yes; not for map.

**Q: Channel vs mutex?**  
A: Single goroutine owns map via channel ops — race-free.
