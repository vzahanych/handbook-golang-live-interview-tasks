# Constant-time secret comparison

## Live interview task
Compare fixed-format authentication tags without early-exit timing leakage.

## Candidate solution

```go
func equalTag(got, want []byte) bool {
	return len(got) == len(want) && subtle.ConstantTimeCompare(got, want) == 1
}
```

## Interview notes / pitfalls
- Length still affects timing; fixed-length decoded values are simplest.
- Passwords require a password-hashing algorithm, not direct constant-time comparison.
- Do not expose whether a username or secret was the failing component.
