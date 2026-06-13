# Cryptographically random token

## Live interview task
Generate a URL-safe 256-bit reset token.

## Candidate solution

```go
func token() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil { return "", err }
	return base64.RawURLEncoding.EncodeToString(b), nil
}
```

Imports: `crypto/rand`, `encoding/base64`, `io`.

## Interview notes / pitfalls
- Do not use `math/rand` for secrets.
- Store a hash of bearer tokens when plaintext recovery is unnecessary.
- Token entropy, expiration, single use, and transport security are separate requirements.
