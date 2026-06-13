# type parameters not runtime polymorphism

## Live interview task
Explain that Go generics are compile-time monomorphization, not JVM-style runtime polymorphism.

## Concepts covered
- generics instantiation
- monomorphization
- interfaces vs type params

## Candidate solution

```go
package main

import "fmt"

func Print[T any](v T) {
    fmt.Println(v)
}

type Dog struct{}
func (Dog) Speak() string { return "woof" }

type Cat struct{}
func (Cat) Speak() string { return "meow" }

type Speaker interface{ Speak() string }

func main() {
    // Compile-time: compiler generates Print[int], Print[string], etc.
    Print(42)
    Print("hi")

    // Runtime polymorphism: interface holds dynamic type
    var s Speaker = Dog{}
    fmt.Println(s.Speak())

    // Generics cannot replace this without constraint listing all types:
    // func SpeakAll[T Speaker](v T) — each concrete T still separate instantiation
}
```

## Run

```bash
go run .
```

## Interview notes / pitfalls
- Each unique type argument creates specialized code (monomorphization) — can increase binary size.
- Interface: one code path, dynamic dispatch via itable — runtime cost, flexible.
- Use **generics** for algorithms on many types (slice ops); **interfaces** for behavior injection and mocking.
- Cannot do `[]Speaker` containing mixed types with generics alone — that's interface slice.

## Q&A

**Q: Reflection vs generics?**  
A: Reflection runtime slow, no type safety; generics compile-time checked.

**Q: Type assertion on type param?**  
A: Possible with constraint or `any` — loses generic benefit.

**Q: Binary bloat concern?**  
A: Real for many instantiations — use interfaces at hot boundaries if needed.

**Q: `interface{}` before generics?**  
A: Same heap boxing for non-pointer values in interfaces — generics often stack-allocated.

**Q: Interview one-liner?**  
A: "Generics are for static type lists; interfaces are for dynamic behavior."
