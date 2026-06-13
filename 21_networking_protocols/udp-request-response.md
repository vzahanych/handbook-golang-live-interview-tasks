# UDP request and response

## Live interview task
Read datagrams, uppercase each payload, and send the response to the originating address.

## Interview notes / pitfalls
- UDP delivery, ordering, and uniqueness are not guaranteed.
- A fixed read buffer truncates oversized datagrams.
- There is no per-client connection state unless the application creates it.
- Add request IDs when clients need to match retries and responses.
