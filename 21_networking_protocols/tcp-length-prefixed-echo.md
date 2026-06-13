# TCP length-prefixed echo

## Live interview task
Implement an echo server where each message is a 4-byte big-endian length followed by payload bytes.

## Required behavior
- Use `io.ReadFull` for both header and body.
- Reject frames above a configured maximum before allocating.
- Set connection deadlines and handle each connection independently.
- Use a write loop or `io.CopyN`; `Write` may be short.

## Interview notes / pitfalls
- TCP is a byte stream, not a message protocol.
- Never allocate directly from an untrusted length without a limit.
