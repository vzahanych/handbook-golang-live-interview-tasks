# Task catalog

Generated list of 228 live interview tasks. Run `go run ./tools/generate_catalog.go` after adding or removing task files.

## 01_language_basics - Language basics and source organization

- [blank-identifier-for-compile-time-interface-check](./01_language_basics/blank-identifier-for-compile-time-interface-check.md)
- [blank-import-side-effect](./01_language_basics/blank-import-side-effect.md)
- [exported-identifiers-and-package-api](./01_language_basics/exported-identifiers-and-package-api.md)
- [hello-main-package-and-init-order](./01_language_basics/hello-main-package-and-init-order.md)
- [iota-bitmask-permissions](./01_language_basics/iota-bitmask-permissions.md)
- [loop-variable-closure-capture](./01_language_basics/loop-variable-closure-capture.md)
- [shadowing-short-variable-trap](./01_language_basics/shadowing-short-variable-trap.md)
- [type-alias-vs-new-defined-type](./01_language_basics/type-alias-vs-new-defined-type.md)

## 02_control_flow - Control flow and statements

- [defer-in-loop-performance-trap](./02_control_flow/defer-in-loop-performance-trap.md)
- [defer-lifo-resource-cleanup](./02_control_flow/defer-lifo-resource-cleanup.md)
- [for-range-over-integer-go122](./02_control_flow/for-range-over-integer-go122.md)
- [for-range-over-iter-seq-go123](./02_control_flow/for-range-over-iter-seq-go123.md)
- [for-range-stdlib-iterators-go123](./02_control_flow/for-range-stdlib-iterators-go123.md)
- [labeled-break-out-of-nested-loop](./02_control_flow/labeled-break-out-of-nested-loop.md)
- [select-default-nonblocking-receive](./02_control_flow/select-default-nonblocking-receive.md)
- [switch-fallthrough-trap](./02_control_flow/switch-fallthrough-trap.md)
- [switch-without-expression-classifier](./02_control_flow/switch-without-expression-classifier.md)
- [type-switch-format-any](./02_control_flow/type-switch-format-any.md)

## 03_types_values_memory - Types, values, identity, assignability, memory model

- [array-copy-vs-slice-sharing](./03_types_values_memory/array-copy-vs-slice-sharing.md)
- [assignability-with-underlying-types](./03_types_values_memory/assignability-with-underlying-types.md)
- [method-set-pointer-vs-value-receiver](./03_types_values_memory/method-set-pointer-vs-value-receiver.md)
- [nil-interface-vs-interface-holding-nil-pointer](./03_types_values_memory/nil-interface-vs-interface-holding-nil-pointer.md)
- [slice-append-shared-backing-trap](./03_types_values_memory/slice-append-shared-backing-trap.md)
- [struct-padding-order-fields](./03_types_values_memory/struct-padding-order-fields.md)
- [zero-values-reference-types](./03_types_values_memory/zero-values-reference-types.md)

## 04_arrays_slices_strings - Arrays, slices, strings and bytes

- [binary-search-lower-bound](./04_arrays_slices_strings/binary-search-lower-bound.md)
- [chunk-slice](./04_arrays_slices_strings/chunk-slice.md)
- [copy-overlapping-slices](./04_arrays_slices_strings/copy-overlapping-slices.md)
- [dedupe-sorted-slice-in-place](./04_arrays_slices_strings/dedupe-sorted-slice-in-place.md)
- [is-palindrome-ignore-nonletters](./04_arrays_slices_strings/is-palindrome-ignore-nonletters.md)
- [merge-sorted-slices](./04_arrays_slices_strings/merge-sorted-slices.md)
- [remove-element-no-order](./04_arrays_slices_strings/remove-element-no-order.md)
- [remove-element-preserve-order](./04_arrays_slices_strings/remove-element-preserve-order.md)
- [reverse-slice-in-place](./04_arrays_slices_strings/reverse-slice-in-place.md)
- [rotate-slice-left-k](./04_arrays_slices_strings/rotate-slice-left-k.md)
- [rune-safe-reverse-string](./04_arrays_slices_strings/rune-safe-reverse-string.md)
- [slices-package-contains-clone-go121](./04_arrays_slices_strings/slices-package-contains-clone-go121.md)
- [sliding-window-max-sum](./04_arrays_slices_strings/sliding-window-max-sum.md)
- [string-builder-join](./04_arrays_slices_strings/string-builder-join.md)
- [string-immutable-byte-slice-trap](./04_arrays_slices_strings/string-immutable-byte-slice-trap.md)

## 05_maps_sets - Maps, sets and grouping

- [first-non-repeating-rune](./05_maps_sets/first-non-repeating-rune.md)
- [group-users-by-role](./05_maps_sets/group-users-by-role.md)
- [invert-map-detect-duplicates](./05_maps_sets/invert-map-detect-duplicates.md)
- [lru-cache-with-list-and-map](./05_maps_sets/lru-cache-with-list-and-map.md)
- [map-concurrent-write-panic](./05_maps_sets/map-concurrent-write-panic.md)
- [map-delete-while-iterate](./05_maps_sets/map-delete-while-iterate.md)
- [set-with-map-struct](./05_maps_sets/set-with-map-struct.md)
- [sort-map-keys-for-deterministic-output](./05_maps_sets/sort-map-keys-for-deterministic-output.md)
- [two-sum-map](./05_maps_sets/two-sum-map.md)
- [word-frequency-counter](./05_maps_sets/word-frequency-counter.md)

## 06_functions_methods_interfaces - Functions, methods, interfaces

- [accept-interfaces-return-concrete](./06_functions_methods_interfaces/accept-interfaces-return-concrete.md)
- [closure-counter](./06_functions_methods_interfaces/closure-counter.md)
- [functional-options-pattern](./06_functions_methods_interfaces/functional-options-pattern.md)
- [interface-small-storage](./06_functions_methods_interfaces/interface-small-storage.md)
- [io-reader-line-counter](./06_functions_methods_interfaces/io-reader-line-counter.md)
- [method-expression-vs-method-value](./06_functions_methods_interfaces/method-expression-vs-method-value.md)
- [sort-custom-structs](./06_functions_methods_interfaces/sort-custom-structs.md)
- [variadic-append-trap](./06_functions_methods_interfaces/variadic-append-trap.md)

## 07_generics - Generics and type constraints

- [constraint-with-method-and-underlying-type](./07_generics/constraint-with-method-and-underlying-type.md)
- [generic-channel-fan-in](./07_generics/generic-channel-fan-in.md)
- [generic-comparable-map-key-cache](./07_generics/generic-comparable-map-key-cache.md)
- [generic-linked-list](./07_generics/generic-linked-list.md)
- [generic-map-filter-reduce](./07_generics/generic-map-filter-reduce.md)
- [generic-set-operations](./07_generics/generic-set-operations.md)
- [generic-stack](./07_generics/generic-stack.md)
- [ordered-min-max-constraint](./07_generics/ordered-min-max-constraint.md)
- [slices-sortfunc-generic-go121](./07_generics/slices-sortfunc-generic-go121.md)
- [type-parameters-not-runtime-polymorphism](./07_generics/type-parameters-not-runtime-polymorphism.md)

## 08_errors_panic_defer - Errors, panic, recover and defer

- [custom-error-type-with-errors-as](./08_errors_panic_defer/custom-error-type-with-errors-as.md)
- [defer-argument-evaluation-time](./08_errors_panic_defer/defer-argument-evaluation-time.md)
- [defer-modifies-named-result](./08_errors_panic_defer/defer-modifies-named-result.md)
- [defer-recover-must-be-func-literal](./08_errors_panic_defer/defer-recover-must-be-func-literal.md)
- [errors-join-go120](./08_errors_panic_defer/errors-join-go120.md)
- [panic-safe-parser](./08_errors_panic_defer/panic-safe-parser.md)
- [recover-at-goroutine-boundary](./08_errors_panic_defer/recover-at-goroutine-boundary.md)
- [wrap-and-match-errors](./08_errors_panic_defer/wrap-and-match-errors.md)

## 09_concurrency_channels - Goroutines, channels and select

- [channel-generator-and-range](./09_concurrency_channels/channel-generator-and-range.md)
- [fan-out-fan-in-pipeline](./09_concurrency_channels/fan-out-fan-in-pipeline.md)
- [first-result-wins](./09_concurrency_channels/first-result-wins.md)
- [or-done-channel-combinator](./09_concurrency_channels/or-done-channel-combinator.md)
- [rate-limiter-with-ticker](./09_concurrency_channels/rate-limiter-with-ticker.md)
- [safe-channel-closing-owner](./09_concurrency_channels/safe-channel-closing-owner.md)
- [select-fairness-and-default](./09_concurrency_channels/select-fairness-and-default.md)
- [semaphore-bounded-concurrency](./09_concurrency_channels/semaphore-bounded-concurrency.md)
- [send-on-closed-channel-panic](./09_concurrency_channels/send-on-closed-channel-panic.md)
- [tee-channel-values](./09_concurrency_channels/tee-channel-values.md)
- [waitgroup-basic-worker-start](./09_concurrency_channels/waitgroup-basic-worker-start.md)
- [worker-pool-jobs-results](./09_concurrency_channels/worker-pool-jobs-results.md)

## 10_context_cancellation - Context, cancellation and timeouts

- [cancellable-worker-loop](./10_context_cancellation/cancellable-worker-loop.md)
- [context-afterfunc-cleanup-go121](./10_context_cancellation/context-afterfunc-cleanup-go121.md)
- [context-never-pass-nil](./10_context_cancellation/context-never-pass-nil.md)
- [context-timeout-http-request](./10_context_cancellation/context-timeout-http-request.md)
- [context-values-request-id](./10_context_cancellation/context-values-request-id.md)
- [errgroup-like-cancel-on-error](./10_context_cancellation/errgroup-like-cancel-on-error.md)
- [pipeline-with-context-cancellation](./10_context_cancellation/pipeline-with-context-cancellation.md)

## 11_sync_atomics - sync package and atomics

- [atomic-counter](./11_sync_atomics/atomic-counter.md)
- [atomic-vs-mutex-when](./11_sync_atomics/atomic-vs-mutex-when.md)
- [mutex-protected-counter](./11_sync_atomics/mutex-protected-counter.md)
- [rwmutex-cache](./11_sync_atomics/rwmutex-cache.md)
- [sync-cond-broadcast](./11_sync_atomics/sync-cond-broadcast.md)
- [sync-map-use-case](./11_sync_atomics/sync-map-use-case.md)
- [sync-once-lazy-init](./11_sync_atomics/sync-once-lazy-init.md)
- [sync-pool-buffer-reuse](./11_sync_atomics/sync-pool-buffer-reuse.md)

## 12_parallel_slice_calculations - Parallel slice calculations

- [parallel-any-match-cancel](./12_parallel_slice_calculations/parallel-any-match-cancel.md)
- [parallel-dedupe-sharded-map](./12_parallel_slice_calculations/parallel-dedupe-sharded-map.md)
- [parallel-dot-product](./12_parallel_slice_calculations/parallel-dot-product.md)
- [parallel-filter-stable](./12_parallel_slice_calculations/parallel-filter-stable.md)
- [parallel-histogram-by-buckets](./12_parallel_slice_calculations/parallel-histogram-by-buckets.md)
- [parallel-map-preserve-order](./12_parallel_slice_calculations/parallel-map-preserve-order.md)
- [parallel-min-max](./12_parallel_slice_calculations/parallel-min-max.md)
- [parallel-prefix-sum-two-pass](./12_parallel_slice_calculations/parallel-prefix-sum-two-pass.md)
- [parallel-processing-with-context](./12_parallel_slice_calculations/parallel-processing-with-context.md)
- [parallel-sort-and-merge](./12_parallel_slice_calculations/parallel-sort-and-merge.md)
- [parallel-sum-configurable-workers](./12_parallel_slice_calculations/parallel-sum-configurable-workers.md)
- [parallel-sum-int-slice](./12_parallel_slice_calculations/parallel-sum-int-slice.md)
- [parallel-transform-error-cancel](./12_parallel_slice_calculations/parallel-transform-error-cancel.md)
- [parallel-word-count](./12_parallel_slice_calculations/parallel-word-count.md)
- [when-parallel-not-worth-it](./12_parallel_slice_calculations/when-parallel-not-worth-it.md)

## 13_http_servers_clients - HTTP servers and clients

- [bounded-concurrent-fetch-handler](./13_http_servers_clients/bounded-concurrent-fetch-handler.md)
- [file-upload-handler](./13_http_servers_clients/file-upload-handler.md)
- [graceful-shutdown-server](./13_http_servers_clients/graceful-shutdown-server.md)
- [http-client-timeout-and-status-check](./13_http_servers_clients/http-client-timeout-and-status-check.md)
- [http-server-production-timeouts](./13_http_servers_clients/http-server-production-timeouts.md)
- [httptest-server-client](./13_http_servers_clients/httptest-server-client.md)
- [in-memory-rate-limit-middleware](./13_http_servers_clients/in-memory-rate-limit-middleware.md)
- [json-api-handler](./13_http_servers_clients/json-api-handler.md)
- [middleware-logging-status-recorder](./13_http_servers_clients/middleware-logging-status-recorder.md)
- [minimal-http-server](./13_http_servers_clients/minimal-http-server.md)
- [repeated-url-command-line-flag](./13_http_servers_clients/repeated-url-command-line-flag.md)
- [request-scoped-context-value-middleware](./13_http_servers_clients/request-scoped-context-value-middleware.md)
- [responsewriter-writeheader-once](./13_http_servers_clients/responsewriter-writeheader-once.md)
- [streaming-response-flusher](./13_http_servers_clients/streaming-response-flusher.md)

## 14_json_files_io - JSON, files and I/O

- [copy-file-with-buffer](./14_json_files_io/copy-file-with-buffer.md)
- [csv-reader-to-structs](./14_json_files_io/csv-reader-to-structs.md)
- [filepath-walk-extension-filter](./14_json_files_io/filepath-walk-extension-filter.md)
- [gzip-compress-decompress](./14_json_files_io/gzip-compress-decompress.md)
- [json-marshal-unmarshal-struct-tags](./14_json_files_io/json-marshal-unmarshal-struct-tags.md)
- [json-omitempty-nil-pointer-trap](./14_json_files_io/json-omitempty-nil-pointer-trap.md)
- [read-file-lines-scanner](./14_json_files_io/read-file-lines-scanner.md)
- [stream-json-lines-decoder](./14_json_files_io/stream-json-lines-decoder.md)

## 15_testing_benchmarking - Testing, benchmarking and fuzzing

- [benchmark-string-concat](./15_testing_benchmarking/benchmark-string-concat.md)
- [fuzz-reverse-twice](./15_testing_benchmarking/fuzz-reverse-twice.md)
- [golden-file-test](./15_testing_benchmarking/golden-file-test.md)
- [httptest-handler-test](./15_testing_benchmarking/httptest-handler-test.md)
- [race-detector-shared-map-test](./15_testing_benchmarking/race-detector-shared-map-test.md)
- [t-cleanup-and-helpers](./15_testing_benchmarking/t-cleanup-and-helpers.md)
- [table-driven-unit-test](./15_testing_benchmarking/table-driven-unit-test.md)

## 16_runtime_performance - Runtime and performance interview tasks

- [avoid-large-range-value-copy](./16_runtime_performance/avoid-large-range-value-copy.md)
- [bounds-check-elimination-hint](./16_runtime_performance/bounds-check-elimination-hint.md)
- [clip-slice-capacity-before-append](./16_runtime_performance/clip-slice-capacity-before-append.md)
- [escape-analysis-example](./16_runtime_performance/escape-analysis-example.md)
- [map-increment-efficient](./16_runtime_performance/map-increment-efficient.md)
- [object-pool-with-reset](./16_runtime_performance/object-pool-with-reset.md)
- [pprof-allocs-quick-start](./16_runtime_performance/pprof-allocs-quick-start.md)
- [preallocate-slice-before-append](./16_runtime_performance/preallocate-slice-before-append.md)
- [strings-builder-grow-once](./16_runtime_performance/strings-builder-grow-once.md)

## 17_colly_scraping - Colly scraping examples

- [colly-basic-title-scraper](./17_colly_scraping/colly-basic-title-scraper.md)
- [colly-concurrent-multiple-urls](./17_colly_scraping/colly-concurrent-multiple-urls.md)
- [colly-coursera-course-card-style](./17_colly_scraping/colly-coursera-course-card-style.md)
- [colly-error-handling](./17_colly_scraping/colly-error-handling.md)
- [colly-hackernews-comments-style](./17_colly_scraping/colly-hackernews-comments-style.md)
- [colly-local-html-file](./17_colly_scraping/colly-local-html-file.md)
- [colly-login-post-before-scrape](./17_colly_scraping/colly-login-post-before-scrape.md)
- [colly-max-depth-crawler](./17_colly_scraping/colly-max-depth-crawler.md)
- [colly-multipart-form-submit](./17_colly_scraping/colly-multipart-form-submit.md)
- [colly-parallel-async-scraper](./17_colly_scraping/colly-parallel-async-scraper.md)
- [colly-proxy-switcher](./17_colly_scraping/colly-proxy-switcher.md)
- [colly-queue-crawl](./17_colly_scraping/colly-queue-crawl.md)
- [colly-rate-limit-delay](./17_colly_scraping/colly-rate-limit-delay.md)
- [colly-reddit-posts-style](./17_colly_scraping/colly-reddit-posts-style.md)
- [colly-request-context-metadata](./17_colly_scraping/colly-request-context-metadata.md)
- [colly-request-timeout-and-context](./17_colly_scraping/colly-request-timeout-and-context.md)
- [colly-robots-txt-etiquette](./17_colly_scraping/colly-robots-txt-etiquette.md)
- [colly-safe-target-validation](./17_colly_scraping/colly-safe-target-validation.md)
- [colly-scraper-server-endpoint](./17_colly_scraping/colly-scraper-server-endpoint.md)
- [colly-shopify-sitemap-style](./17_colly_scraping/colly-shopify-sitemap-style.md)
- [colly-url-filter](./17_colly_scraping/colly-url-filter.md)
- [colly-xkcd-store-items-style](./17_colly_scraping/colly-xkcd-store-items-style.md)

## 18_mini_projects - Small integrated projects

- [bounded-crawler-standard-library](./18_mini_projects/bounded-crawler-standard-library.md)
- [chat-broadcast-server-skeleton](./18_mini_projects/chat-broadcast-server-skeleton.md)
- [cli-todo-json-file](./18_mini_projects/cli-todo-json-file.md)
- [flag-driven-concurrent-colly-service](./18_mini_projects/flag-driven-concurrent-colly-service.md)
- [in-memory-url-shortener-http](./18_mini_projects/in-memory-url-shortener-http.md)
- [pubsub-topic-broker](./18_mini_projects/pubsub-topic-broker.md)
- [ttl-cache-with-cleanup-goroutine](./18_mini_projects/ttl-cache-with-cleanup-goroutine.md)

## 19_cli_configuration - CLI and configuration

- [config-precedence-flags-env-defaults](./19_cli_configuration/config-precedence-flags-env-defaults.md)
- [custom-csv-flag-value](./19_cli_configuration/custom-csv-flag-value.md)
- [read-lines-from-file-or-stdin](./19_cli_configuration/read-lines-from-file-or-stdin.md)
- [secret-flag-redaction](./19_cli_configuration/secret-flag-redaction.md)
- [subcommands-with-flagsets](./19_cli_configuration/subcommands-with-flagsets.md)

## 20_database_sql - Database and SQL

- [connection-pool-tuning](./20_database_sql/connection-pool-tuning.md)
- [optimistic-locking-version-column](./20_database_sql/optimistic-locking-version-column.md)
- [query-context-and-rows-lifecycle](./20_database_sql/query-context-and-rows-lifecycle.md)
- [sql-null-values](./20_database_sql/sql-null-values.md)
- [transaction-transfer-funds](./20_database_sql/transaction-transfer-funds.md)

## 21_networking_protocols - Networking and protocols

- [dns-lookup-with-timeout](./21_networking_protocols/dns-lookup-with-timeout.md)
- [netip-private-address-check](./21_networking_protocols/netip-private-address-check.md)
- [tcp-length-prefixed-echo](./21_networking_protocols/tcp-length-prefixed-echo.md)
- [udp-request-response](./21_networking_protocols/udp-request-response.md)
- [url-resolution-and-normalization](./21_networking_protocols/url-resolution-and-normalization.md)

## 22_time_scheduling - Time and scheduling

- [debounce-events](./22_time_scheduling/debounce-events.md)
- [injectable-clock-tests](./22_time_scheduling/injectable-clock-tests.md)
- [ticker-worker-clean-shutdown](./22_time_scheduling/ticker-worker-clean-shutdown.md)
- [time-zone-day-boundary](./22_time_scheduling/time-zone-day-boundary.md)
- [timer-reset-correctly](./22_time_scheduling/timer-reset-correctly.md)

## 23_reflection_encoding - Reflection and encoding

- [copy-matching-struct-fields](./23_reflection_encoding/copy-matching-struct-fields.md)
- [custom-json-time-format](./23_reflection_encoding/custom-json-time-format.md)
- [type-switch-vs-reflection](./23_reflection_encoding/type-switch-vs-reflection.md)
- [unsafe-string-byte-conversion](./23_reflection_encoding/unsafe-string-byte-conversion.md)
- [validate-required-struct-tags](./23_reflection_encoding/validate-required-struct-tags.md)

## 24_design_architecture - Design and architecture

- [dependency-injection-small-interface](./24_design_architecture/dependency-injection-small-interface.md)
- [explicit-state-machine](./24_design_architecture/explicit-state-machine.md)
- [idempotent-command-handler](./24_design_architecture/idempotent-command-handler.md)
- [repository-transaction-boundary](./24_design_architecture/repository-transaction-boundary.md)
- [retry-with-backoff-policy](./24_design_architecture/retry-with-backoff-policy.md)

## 25_security_observability - Security, reliability, and observability

- [constant-time-secret-compare](./25_security_observability/constant-time-secret-compare.md)
- [crypto-random-token](./25_security_observability/crypto-random-token.md)
- [liveness-readiness-shutdown](./25_security_observability/liveness-readiness-shutdown.md)
- [panic-recovery-http-middleware](./25_security_observability/panic-recovery-http-middleware.md)
- [structured-slog-context](./25_security_observability/structured-slog-context.md)

## 26_algorithms_data_structures - Algorithms and data structures

- [bfs-shortest-path-grid](./26_algorithms_data_structures/bfs-shortest-path-grid.md)
- [lru-cache-generic](./26_algorithms_data_structures/lru-cache-generic.md)
- [merge-overlapping-intervals](./26_algorithms_data_structures/merge-overlapping-intervals.md)
- [top-k-with-min-heap](./26_algorithms_data_structures/top-k-with-min-heap.md)
- [union-find-components](./26_algorithms_data_structures/union-find-components.md)

## 27_os_processes - OS, files, and processes

- [atomic-file-replacement](./27_os_processes/atomic-file-replacement.md)
- [embed-static-assets](./27_os_processes/embed-static-assets.md)
- [exec-command-context](./27_os_processes/exec-command-context.md)
- [file-permission-check](./27_os_processes/file-permission-check.md)
- [signal-notify-context](./27_os_processes/signal-notify-context.md)

