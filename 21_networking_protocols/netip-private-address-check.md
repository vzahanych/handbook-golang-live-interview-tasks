# Private address check with netip

## Live interview task
Implement an outbound-network policy using `net/netip`.

## Candidate solution

```go
func blocked(addr netip.Addr) bool {
	return !addr.IsValid() || addr.IsUnspecified() || addr.IsLoopback() ||
		addr.IsPrivate() || addr.IsLinkLocalUnicast() ||
		addr.IsLinkLocalMulticast() || addr.IsMulticast()
}
```

## Interview notes / pitfalls
- Call `Unmap` when IPv4-mapped IPv6 addresses matter.
- Address classification must be enforced at dial time to resist DNS rebinding.
