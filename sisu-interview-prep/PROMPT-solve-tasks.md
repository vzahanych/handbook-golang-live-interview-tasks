# Prompt: solve a simplesurance live-coding task

Use this prompt to turn a raw interview task in `sisu-interview-prep/` into a
worked solution that follows the exact conventions established in
`task-01-csv-summarizer-solution/`. Paste it to the assistant and name the task
folder you want solved.

---

## The prompt

> You are helping me prepare for a Go live-coding interview. The reference for
> **how** to work is `sisu-interview-prep/task-01-csv-summarizer-solution/` —
> study `main.go`, `main_test.go`, and `sample.csv` there first, then apply the
> same method to the task I name below.
>
> **Task to solve:** `sisu-interview-prep/<TASK-FOLDER>`
>
> Follow these rules exactly.
>
> ### 1. Folder layout — never edit the verbatim copy
> Each task exists twice: `<task>` (the email code, untouched) and
> `<task>-solution` (the worked copy). Only ever change files under
> `-solution/`. If the `-solution` copy does not exist yet, duplicate the
> verbatim folder into it first.
>
> ### 2. Answer the task's questions inline, as comments
> The task code opens with numbered questions/instructions in comments. Answer
> each one **directly below it**, as a comment, before writing any code. Write
> the answers top-down: the first sentence states the point, each next sentence
> specializes the previous one — no disconnected one-liners. When a question has
> two valid strategies (e.g. fail-fast vs. accumulate), name both, then say which
> is better and why.
>
> ### 3. Correct the code surgically — comment the old line, add the new one
> Do **not** rewrite functions or move code into new files. For every change,
> leave the original line commented out and put the new line right beside it, so
> at the interview I can reproduce the fix by commenting a few lines and typing a
> few new ones. Keep the same functions in the same place. Annotate each edit
> with which question it answers (e.g. `// answer 2: ...`).
>
> ### 4. Make design choices explicit with a switch
> When the task admits two behaviors, implement both behind a switch (a bool
> parameter, and a query parameter like `?fail_fast=true` at the HTTP layer)
> rather than picking one silently.
>
> ### 5. Add the simplest graceful shutdown
> For any HTTP server, add graceful shutdown the minimal way: keep
> `http.HandleFunc` on `DefaultServeMux`, create a named `srv := &http.Server{
> Addr: ":8080"}` (nil handler ⇒ DefaultServeMux), run `srv.ListenAndServe()` in
> a goroutine, block on `signal.NotifyContext(..., os.Interrupt, syscall.SIGTERM)`,
> then `srv.Shutdown(ctx)` with a `context.WithTimeout` deadline. Treat
> `http.ErrServerClosed` as success. Do **not** introduce a custom `mux`.
>
> ### 6. Provide realistic sample data with deliberate errors
> Add a `sample.<ext>` file (e.g. `sample.csv`) containing both good rows and a
> few intentionally bad cells, so both switch paths are demonstrable.
>
> ### 7. Add a "How to run and test" comment block at the bottom of main.go
> Document the **real** interface — this is an HTTP server, not a CLI with flags.
> Give `go run .` plus accurate `curl` commands for each path (note
> `--data-binary` to preserve newlines, and that the switch is a query
> parameter). State the method restriction (POST only ⇒ GET returns 405).
>
> ### 8. Keep tests minimal — this is live coding, not a test suite
> In a sibling `main_test.go` (same `package main`, so unexported functions are
> reachable) write exactly:
> - **one unit test** of the core function via `strings.NewReader(...)`,
>   asserting the sums and that a bad cell is collected (not dropped);
> - **one integration test** of the handler via `httptest.NewRecorder()` (no
>   real server), asserting the status code;
> - **one benchmark** that builds ~1 MB of input once, then times the parser with
>   `for b.Loop()` and `b.ReportAllocs()` / `b.SetBytes(...)`. Make the sample
>   size a named constant (`const MB = 1 << 20; const targetSize = 1 * MB`) so
>   it is obvious how to resize — never a bare `1<<20` in the loop condition.
>
> ### 9. Verify for real, then report honestly
> Build it, run the server, hit every documented `curl` path, and confirm the
> graceful shutdown by sending `SIGTERM` (`kill -TERM <pid>`) and checking the
> logs and that the port is freed. Run `go test -v` and `go test -bench=.
> -benchmem`. If a test surfaces a bug (as the label-column bug surfaced in
> task-01), say so plainly, fix it in the surgical style, and re-verify. Report
> failures with their output; never claim green without running it.
>
> ### 10. Prepare interview talking points
> After each change, give me 2–3 short talking points I can say out loud:
> the trade-off I chose, why, and the obvious next optimization (e.g. streaming
> keeps memory O(one row); `reader.ReuseRecord = true` cuts allocs but you must
> copy any retained slice).

---

## Why these rules (context for the assistant)

The interview is a ~45-minute live-coding session with screen sharing where I
narrate my thinking and **may not use AI**. So the `-solution/` folder is study
material, not something I paste: the value is that each fix is a small,
memorable diff (comment a line, add a line) and each design question already has
a spoken answer attached. The verbatim `<task>` folder is the starting point I
reproduce from scratch on the call; the `-solution` folder is the destination I
have rehearsed. Keeping edits surgical and functions in place is what makes the
solution reproducible under time pressure.
