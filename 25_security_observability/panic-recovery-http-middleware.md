# HTTP panic recovery middleware

## Live interview task
Recover handler panics, log a stack, and return a generic 500 response.

## Interview notes / pitfalls
- Recovery must run in the same goroutine as the panic.
- Headers may already be committed, so a clean 500 cannot always be sent.
- Do not return panic details to clients.
- A recovered process may still contain corrupted application state; classify panics carefully.
