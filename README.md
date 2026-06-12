# Golang live interview task examples

Generated from the handbook categories and organized as one Markdown file per example.

Total examples: **144**.

## Folder layout
- `01_language_basics/` — Language basics and source organization (6 examples)
- `02_control_flow/` — Control flow and statements (6 examples)
- `03_types_values_memory/` — Types, values, identity, assignability, memory model (6 examples)
- `04_arrays_slices_strings/` — Arrays, slices, strings and bytes (13 examples)
- `05_maps_sets/` — Maps, sets and grouping (8 examples)
- `06_functions_methods_interfaces/` — Functions, methods, interfaces (6 examples)
- `07_generics/` — Generics and type constraints (8 examples)
- `08_errors_panic_defer/` — Errors, panic, recover and defer (6 examples)
- `09_concurrency_channels/` — Goroutines, channels and select (10 examples)
- `10_context_cancellation/` — Context, cancellation and timeouts (5 examples)
- `11_sync_atomics/` — sync package and atomics (6 examples)
- `12_parallel_slice_calculations/` — Parallel slice calculations (9 examples)
- `13_http_servers_clients/` — HTTP servers and clients (10 examples)
- `14_json_files_io/` — JSON, files and I/O (7 examples)
- `15_testing_benchmarking/` — Testing, benchmarking and fuzzing (6 examples)
- `16_runtime_performance/` — Runtime and performance interview tasks (8 examples)
- `17_colly_scraping/` — Colly scraping examples (18 examples)
- `18_mini_projects/` — Small integrated projects (6 examples)

## How to use

Each example is intentionally short enough for live coding. A file contains:

1. the task prompt,
2. the Go concepts being tested,
3. a candidate solution,
4. run command,
5. pitfalls and follow-up questions.

Most examples are standard-library only. The Colly examples require:

```bash
go get github.com/gocolly/colly/v2
```

## Full index

## 01_language_basics — Language basics and source organization

- [hello-main-package-and-init-order](./01_language_basics/hello-main-package-and-init-order.md)
- [iota-bitmask-permissions](./01_language_basics/iota-bitmask-permissions.md)
- [shadowing-short-variable-trap](./01_language_basics/shadowing-short-variable-trap.md)
- [blank-identifier-for-compile-time-interface-check](./01_language_basics/blank-identifier-for-compile-time-interface-check.md)
- [exported-identifiers-and-package-api](./01_language_basics/exported-identifiers-and-package-api.md)
- [type-alias-vs-new-defined-type](./01_language_basics/type-alias-vs-new-defined-type.md)

## 02_control_flow — Control flow and statements

- [for-range-over-integer-go122](./02_control_flow/for-range-over-integer-go122.md)
- [labeled-break-out-of-nested-loop](./02_control_flow/labeled-break-out-of-nested-loop.md)
- [switch-without-expression-classifier](./02_control_flow/switch-without-expression-classifier.md)
- [type-switch-format-any](./02_control_flow/type-switch-format-any.md)
- [select-default-nonblocking-receive](./02_control_flow/select-default-nonblocking-receive.md)
- [defer-lifo-resource-cleanup](./02_control_flow/defer-lifo-resource-cleanup.md)

## 03_types_values_memory — Types, values, identity, assignability, memory model

- [array-copy-vs-slice-sharing](./03_types_values_memory/array-copy-vs-slice-sharing.md)
- [struct-padding-order-fields](./03_types_values_memory/struct-padding-order-fields.md)
- [nil-interface-vs-interface-holding-nil-pointer](./03_types_values_memory/nil-interface-vs-interface-holding-nil-pointer.md)
- [method-set-pointer-vs-value-receiver](./03_types_values_memory/method-set-pointer-vs-value-receiver.md)
- [assignability-with-underlying-types](./03_types_values_memory/assignability-with-underlying-types.md)
- [zero-values-reference-types](./03_types_values_memory/zero-values-reference-types.md)

## 04_arrays_slices_strings — Arrays, slices, strings and bytes

- [reverse-slice-in-place](./04_arrays_slices_strings/reverse-slice-in-place.md)
- [rotate-slice-left-k](./04_arrays_slices_strings/rotate-slice-left-k.md)
- [dedupe-sorted-slice-in-place](./04_arrays_slices_strings/dedupe-sorted-slice-in-place.md)
- [remove-element-preserve-order](./04_arrays_slices_strings/remove-element-preserve-order.md)
- [remove-element-no-order](./04_arrays_slices_strings/remove-element-no-order.md)
- [merge-sorted-slices](./04_arrays_slices_strings/merge-sorted-slices.md)
- [binary-search-lower-bound](./04_arrays_slices_strings/binary-search-lower-bound.md)
- [chunk-slice](./04_arrays_slices_strings/chunk-slice.md)
- [rune-safe-reverse-string](./04_arrays_slices_strings/rune-safe-reverse-string.md)
- [is-palindrome-ignore-nonletters](./04_arrays_slices_strings/is-palindrome-ignore-nonletters.md)
- [string-builder-join](./04_arrays_slices_strings/string-builder-join.md)
- [sliding-window-max-sum](./04_arrays_slices_strings/sliding-window-max-sum.md)
- [copy-overlapping-slices](./04_arrays_slices_strings/copy-overlapping-slices.md)

## 05_maps_sets — Maps, sets and grouping

- [word-frequency-counter](./05_maps_sets/word-frequency-counter.md)
- [group-users-by-role](./05_maps_sets/group-users-by-role.md)
- [set-with-map-struct](./05_maps_sets/set-with-map-struct.md)
- [two-sum-map](./05_maps_sets/two-sum-map.md)
- [first-non-repeating-rune](./05_maps_sets/first-non-repeating-rune.md)
- [invert-map-detect-duplicates](./05_maps_sets/invert-map-detect-duplicates.md)
- [sort-map-keys-for-deterministic-output](./05_maps_sets/sort-map-keys-for-deterministic-output.md)
- [lru-cache-with-list-and-map](./05_maps_sets/lru-cache-with-list-and-map.md)

## 06_functions_methods_interfaces — Functions, methods, interfaces

- [closure-counter](./06_functions_methods_interfaces/closure-counter.md)
- [functional-options-pattern](./06_functions_methods_interfaces/functional-options-pattern.md)
- [sort-custom-structs](./06_functions_methods_interfaces/sort-custom-structs.md)
- [interface-small-storage](./06_functions_methods_interfaces/interface-small-storage.md)
- [method-expression-vs-method-value](./06_functions_methods_interfaces/method-expression-vs-method-value.md)
- [io-reader-line-counter](./06_functions_methods_interfaces/io-reader-line-counter.md)

## 07_generics — Generics and type constraints

- [generic-stack](./07_generics/generic-stack.md)
- [generic-map-filter-reduce](./07_generics/generic-map-filter-reduce.md)
- [ordered-min-max-constraint](./07_generics/ordered-min-max-constraint.md)
- [generic-set-operations](./07_generics/generic-set-operations.md)
- [generic-linked-list](./07_generics/generic-linked-list.md)
- [constraint-with-method-and-underlying-type](./07_generics/constraint-with-method-and-underlying-type.md)
- [generic-comparable-map-key-cache](./07_generics/generic-comparable-map-key-cache.md)
- [generic-channel-fan-in](./07_generics/generic-channel-fan-in.md)

## 08_errors_panic_defer — Errors, panic, recover and defer

- [wrap-and-match-errors](./08_errors_panic_defer/wrap-and-match-errors.md)
- [custom-error-type-with-errors-as](./08_errors_panic_defer/custom-error-type-with-errors-as.md)
- [recover-at-goroutine-boundary](./08_errors_panic_defer/recover-at-goroutine-boundary.md)
- [defer-modifies-named-result](./08_errors_panic_defer/defer-modifies-named-result.md)
- [defer-argument-evaluation-time](./08_errors_panic_defer/defer-argument-evaluation-time.md)
- [panic-safe-parser](./08_errors_panic_defer/panic-safe-parser.md)

## 09_concurrency_channels — Goroutines, channels and select

- [waitgroup-basic-worker-start](./09_concurrency_channels/waitgroup-basic-worker-start.md)
- [channel-generator-and-range](./09_concurrency_channels/channel-generator-and-range.md)
- [worker-pool-jobs-results](./09_concurrency_channels/worker-pool-jobs-results.md)
- [fan-out-fan-in-pipeline](./09_concurrency_channels/fan-out-fan-in-pipeline.md)
- [semaphore-bounded-concurrency](./09_concurrency_channels/semaphore-bounded-concurrency.md)
- [or-done-channel-combinator](./09_concurrency_channels/or-done-channel-combinator.md)
- [tee-channel-values](./09_concurrency_channels/tee-channel-values.md)
- [rate-limiter-with-ticker](./09_concurrency_channels/rate-limiter-with-ticker.md)
- [safe-channel-closing-owner](./09_concurrency_channels/safe-channel-closing-owner.md)
- [first-result-wins](./09_concurrency_channels/first-result-wins.md)

## 10_context_cancellation — Context, cancellation and timeouts

- [context-timeout-http-request](./10_context_cancellation/context-timeout-http-request.md)
- [cancellable-worker-loop](./10_context_cancellation/cancellable-worker-loop.md)
- [pipeline-with-context-cancellation](./10_context_cancellation/pipeline-with-context-cancellation.md)
- [context-values-request-id](./10_context_cancellation/context-values-request-id.md)
- [errgroup-like-cancel-on-error](./10_context_cancellation/errgroup-like-cancel-on-error.md)

## 11_sync_atomics — sync package and atomics

- [mutex-protected-counter](./11_sync_atomics/mutex-protected-counter.md)
- [rwmutex-cache](./11_sync_atomics/rwmutex-cache.md)
- [sync-once-lazy-init](./11_sync_atomics/sync-once-lazy-init.md)
- [atomic-counter](./11_sync_atomics/atomic-counter.md)
- [sync-pool-buffer-reuse](./11_sync_atomics/sync-pool-buffer-reuse.md)
- [sync-cond-broadcast](./11_sync_atomics/sync-cond-broadcast.md)

## 12_parallel_slice_calculations — Parallel slice calculations

- [parallel-sum-int-slice](./12_parallel_slice_calculations/parallel-sum-int-slice.md)
- [parallel-map-preserve-order](./12_parallel_slice_calculations/parallel-map-preserve-order.md)
- [parallel-filter-stable](./12_parallel_slice_calculations/parallel-filter-stable.md)
- [parallel-min-max](./12_parallel_slice_calculations/parallel-min-max.md)
- [parallel-word-count](./12_parallel_slice_calculations/parallel-word-count.md)
- [parallel-histogram-by-buckets](./12_parallel_slice_calculations/parallel-histogram-by-buckets.md)
- [parallel-prefix-sum-two-pass](./12_parallel_slice_calculations/parallel-prefix-sum-two-pass.md)
- [parallel-dot-product](./12_parallel_slice_calculations/parallel-dot-product.md)
- [parallel-processing-with-context](./12_parallel_slice_calculations/parallel-processing-with-context.md)

## 13_http_servers_clients — HTTP servers and clients

- [minimal-http-server](./13_http_servers_clients/minimal-http-server.md)
- [json-api-handler](./13_http_servers_clients/json-api-handler.md)
- [middleware-logging-status-recorder](./13_http_servers_clients/middleware-logging-status-recorder.md)
- [http-client-timeout-and-status-check](./13_http_servers_clients/http-client-timeout-and-status-check.md)
- [graceful-shutdown-server](./13_http_servers_clients/graceful-shutdown-server.md)
- [request-scoped-context-value-middleware](./13_http_servers_clients/request-scoped-context-value-middleware.md)
- [streaming-response-flusher](./13_http_servers_clients/streaming-response-flusher.md)
- [file-upload-handler](./13_http_servers_clients/file-upload-handler.md)
- [in-memory-rate-limit-middleware](./13_http_servers_clients/in-memory-rate-limit-middleware.md)
- [httptest-server-client](./13_http_servers_clients/httptest-server-client.md)

## 14_json_files_io — JSON, files and I/O

- [json-marshal-unmarshal-struct-tags](./14_json_files_io/json-marshal-unmarshal-struct-tags.md)
- [stream-json-lines-decoder](./14_json_files_io/stream-json-lines-decoder.md)
- [read-file-lines-scanner](./14_json_files_io/read-file-lines-scanner.md)
- [csv-reader-to-structs](./14_json_files_io/csv-reader-to-structs.md)
- [filepath-walk-extension-filter](./14_json_files_io/filepath-walk-extension-filter.md)
- [gzip-compress-decompress](./14_json_files_io/gzip-compress-decompress.md)
- [copy-file-with-buffer](./14_json_files_io/copy-file-with-buffer.md)

## 15_testing_benchmarking — Testing, benchmarking and fuzzing

- [table-driven-unit-test](./15_testing_benchmarking/table-driven-unit-test.md)
- [httptest-handler-test](./15_testing_benchmarking/httptest-handler-test.md)
- [benchmark-string-concat](./15_testing_benchmarking/benchmark-string-concat.md)
- [fuzz-reverse-twice](./15_testing_benchmarking/fuzz-reverse-twice.md)
- [race-detector-shared-map-test](./15_testing_benchmarking/race-detector-shared-map-test.md)
- [golden-file-test](./15_testing_benchmarking/golden-file-test.md)

## 16_runtime_performance — Runtime and performance interview tasks

- [preallocate-slice-before-append](./16_runtime_performance/preallocate-slice-before-append.md)
- [avoid-large-range-value-copy](./16_runtime_performance/avoid-large-range-value-copy.md)
- [clip-slice-capacity-before-append](./16_runtime_performance/clip-slice-capacity-before-append.md)
- [strings-builder-grow-once](./16_runtime_performance/strings-builder-grow-once.md)
- [map-increment-efficient](./16_runtime_performance/map-increment-efficient.md)
- [bounds-check-elimination-hint](./16_runtime_performance/bounds-check-elimination-hint.md)
- [object-pool-with-reset](./16_runtime_performance/object-pool-with-reset.md)
- [escape-analysis-example](./16_runtime_performance/escape-analysis-example.md)

## 17_colly_scraping — Colly scraping examples

- [colly-basic-title-scraper](./17_colly_scraping/colly-basic-title-scraper.md)
- [colly-error-handling](./17_colly_scraping/colly-error-handling.md)
- [colly-login-post-before-scrape](./17_colly_scraping/colly-login-post-before-scrape.md)
- [colly-max-depth-crawler](./17_colly_scraping/colly-max-depth-crawler.md)
- [colly-parallel-async-scraper](./17_colly_scraping/colly-parallel-async-scraper.md)
- [colly-rate-limit-delay](./17_colly_scraping/colly-rate-limit-delay.md)
- [colly-queue-crawl](./17_colly_scraping/colly-queue-crawl.md)
- [colly-request-context-metadata](./17_colly_scraping/colly-request-context-metadata.md)
- [colly-url-filter](./17_colly_scraping/colly-url-filter.md)
- [colly-proxy-switcher](./17_colly_scraping/colly-proxy-switcher.md)
- [colly-multipart-form-submit](./17_colly_scraping/colly-multipart-form-submit.md)
- [colly-scraper-server-endpoint](./17_colly_scraping/colly-scraper-server-endpoint.md)
- [colly-local-html-file](./17_colly_scraping/colly-local-html-file.md)
- [colly-hackernews-comments-style](./17_colly_scraping/colly-hackernews-comments-style.md)
- [colly-shopify-sitemap-style](./17_colly_scraping/colly-shopify-sitemap-style.md)
- [colly-coursera-course-card-style](./17_colly_scraping/colly-coursera-course-card-style.md)
- [colly-reddit-posts-style](./17_colly_scraping/colly-reddit-posts-style.md)
- [colly-xkcd-store-items-style](./17_colly_scraping/colly-xkcd-store-items-style.md)

## 18_mini_projects — Small integrated projects

- [in-memory-url-shortener-http](./18_mini_projects/in-memory-url-shortener-http.md)
- [ttl-cache-with-cleanup-goroutine](./18_mini_projects/ttl-cache-with-cleanup-goroutine.md)
- [pubsub-topic-broker](./18_mini_projects/pubsub-topic-broker.md)
- [cli-todo-json-file](./18_mini_projects/cli-todo-json-file.md)
- [bounded-crawler-standard-library](./18_mini_projects/bounded-crawler-standard-library.md)
- [chat-broadcast-server-skeleton](./18_mini_projects/chat-broadcast-server-skeleton.md)
