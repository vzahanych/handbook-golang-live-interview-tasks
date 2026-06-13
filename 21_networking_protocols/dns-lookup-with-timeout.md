# DNS lookup with timeout

## Live interview task
Resolve all IPs for a hostname with a one-second deadline.

## Candidate solution

```go
ctx, cancel := context.WithTimeout(parent, time.Second)
defer cancel()
addrs, err := net.DefaultResolver.LookupNetIP(ctx, "ip", host)
```

## Interview notes / pitfalls
- DNS results may contain several IPv4 and IPv6 addresses.
- Resolution success does not imply a connection will succeed.
- Caching must respect changing records and application policy.
