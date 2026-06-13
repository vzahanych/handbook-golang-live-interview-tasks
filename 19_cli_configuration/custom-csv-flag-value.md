# Custom comma-separated flag value

## Live interview task
Implement `-tags=go,backend` as a custom `flag.Value` that trims whitespace and rejects empty tags.

## Candidate solution

```go
type tagsFlag []string

func (t *tagsFlag) String() string { return strings.Join(*t, ",") }
func (t *tagsFlag) Set(raw string) error {
	for _, value := range strings.Split(raw, ",") {
		value = strings.TrimSpace(value)
		if value == "" { return errors.New("empty tag") }
		*t = append(*t, value)
	}
	return nil
}
```

## Interview notes / pitfalls
- Decide whether repeated flags append or replace.
- Deduplicate only if the command contract requires it.
