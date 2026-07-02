# simplesurance (SiSu) — Go Live-Coding Interview Prep

Runnable practice tasks reconstructed from the interview briefing. Each task
lives in its own subfolder as a single `main.go` that holds **both the original
task code (commented at the top) and the proposed solution (active below it)**,
plus tests and a README. The layout mirrors the live session: you are handed the
starter, and turning it into the solution is a matter of commenting a few lines
and adding new code — the commented block is exactly what to comment out, and the
active code is exactly what to add. The interview is a ~45-minute live-coding
session with full screen sharing where you narrate your thinking; the READMEs'
"talking points" sections are written for that.

## Format of the interview (from the briefing)

- Brief intro (**keep it under a minute** — candidates run out of coding time).
- ~45 min live coding with screen sharing; you may Google and use normal tools,
  but they now ask that you **do not use AI**.
- Interleaved questions about scaling to many simultaneous transactions, race
  conditions (wait groups, mutexes), error handling, HTTP, and memory.
- You run the code against sample data they paste into the meeting chat.
- Your questions for them at the end (team, stack, rituals, first-weeks
  expectations).

## The tasks

| # | Folder | Origin | Focus |
|---|--------|--------|-------|
| 01 | [task-01-csv-summarizer](task-01-csv-summarizer) | Code example #1 in the email | Streaming large uploads, typed errors, partial results, resume, benchmark |
| 02 | [task-02-top-pages](task-02-top-pages) | Code example #2 in the email | JSON decode, counting, ranking, deterministic sort |
| 03 | [task-03-multi-sequence-goroutines](task-03-multi-sequence-goroutines) | The exercise the email says is *almost certainly asked* | Single→many→concurrent; WaitGroup, Mutex, bounded fan-out |
| 04 | [task-04-concurrency-patterns](task-04-concurrency-patterns) | The recurring throughput discussion | Data race, WaitGroup misuse, worker pool, context cancellation |
| 05 | [task-05-backend-fanout](task-05-backend-fanout) | The CIL "many back-ends / http requests" thread | Fan-out/fan-in, context deadline, body-close leaks, bounded concurrency |

Tasks 01–02 are the two code snippets the interviewers shared verbatim; each now
lives as a `<task>` starter (the pasted code, untouched) alongside a worked
`<task>-solution` produced with [PROMPT-solve-tasks.md](PROMPT-solve-tasks.md).
Tasks 03–05 are additional starters in the same style — numbered questions atop
intentionally incomplete or buggy code with inline sample data, ready to be
solved the same way. Task 03 is the sequence exercise the email states they
*will almost certainly ask* (single integer → many → concurrent). Task 04 is the
"very large number of simultaneous transactions" discussion as runnable code: a
data race and a WaitGroup bug to fix under `-race`, then a worker pool. Task 05
covers the CIL's fan-out to many back-ends: concurrency, per-request timeouts,
leak-free resource handling, and bounded concurrency.

## Running everything

Every folder is a standalone `package main` under the repo's `examples` module,
so from this directory:

```bash
go build ./...              # compile all tasks
go vet ./...                # static checks
go test ./...               # all tests
go test -race ./...         # race detector (matters for 03 and 04)
go test -bench=. -benchmem ./...
```

To run one task's server or program, `cd` into its folder and `go run .` (see
each README for curl commands and sample data).

## Dry-run checklist before the call

- Share your **entire screen** — close the meeting invite and anything mentioning
  Kazzlo/Zfort first (the briefing calls this out explicitly).
- Represent yourself as a contractor placed through **Kazzlo**; simplesurance is
  not to hear about Zfort.
- Have your one-minute intro and a few questions for the team ready.
- Editor + terminal + browser laid out so you can type, run, and talk at once.
