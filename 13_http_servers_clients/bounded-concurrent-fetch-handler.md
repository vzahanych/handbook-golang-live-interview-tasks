# Bounded concurrent fetch handler

## Live interview task
Build `POST /fetch` that accepts JSON URLs, fetches them concurrently, preserves order, and never runs more than N outbound requests.

## Required behavior
- Decode with `http.MaxBytesReader` and `json.Decoder.DisallowUnknownFields`.
- Validate URL schemes and enforce an SSRF policy.
- Use the request context and an `http.Client` with a timeout.
- Store each response in the slot matching its input index.
- Return per-URL status and error data rather than failing the entire batch.

## Interview notes / pitfalls
- Do not use `http.DefaultClient` for untrusted or slow targets.
- Closing response bodies permits connection reuse; draining may also be needed.
- A semaphore per request does not enforce a process-wide limit.
- DNS rebinding can bypass validation performed only before dialing.
