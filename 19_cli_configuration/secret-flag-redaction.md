# Secret configuration redaction

## Live interview task
Define a configuration struct whose diagnostic output never exposes API keys or passwords.

## Required behavior
- Keep secret values out of `String`, logs, errors, and metrics labels.
- Return a redacted copy for diagnostics.
- Distinguish missing secret from a present secret without printing its value.
- Prefer environment or file descriptors over command-line secrets visible in process listings.

## Follow-up questions
- How should secret rotation work?
- Which configuration fields are safe to expose on a debug endpoint?
