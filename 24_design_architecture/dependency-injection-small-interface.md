# Dependency injection with a small interface

## Live interview task
Make a registration service testable without a real email provider.

## Candidate design

```go
type Mailer interface { SendWelcome(context.Context, string) error }
type Service struct { mailer Mailer }
func NewService(m Mailer) *Service { return &Service{mailer: m} }
```

## Interview notes / pitfalls
- Define the interface near the consumer.
- Constructors should reject invalid mandatory dependencies.
- Avoid interfaces that simply mirror an entire third-party client.
