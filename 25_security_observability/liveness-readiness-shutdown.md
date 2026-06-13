# Liveness and readiness during shutdown

## Live interview task
Implement `/livez` and `/readyz` and make readiness fail before graceful shutdown begins.

## Required behavior
- Liveness reports whether the process event loop is functioning.
- Readiness reports whether the process should receive new traffic.
- Atomically mark unready, allow load balancers time to observe it, then call `Shutdown`.
- Bound dependency checks and avoid turning liveness into a cascading failure detector.
