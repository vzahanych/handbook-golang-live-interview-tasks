# Colly safe target validation

## Live interview task
Validate user-supplied scrape targets and block common SSRF routes.

## Required checks
- Allow only `http` and `https`; reject credentials and malformed hosts.
- Resolve every hostname and reject loopback, private, link-local, multicast, and unspecified addresses.
- Recheck the connected IP in a custom `net.Dialer.Control` or transport dial hook.
- Apply the policy to redirects as well as the initial URL.
- Set response-size, timeout, redirect, and concurrency limits.

## Interview notes / pitfalls
- String-prefix checks do not understand IPv6, encoded hosts, or DNS rebinding.
- An allowlist is stronger than a denylist when targets are known.
- Scraping arbitrary public sites also requires legal, robots, and rate-limit considerations.
