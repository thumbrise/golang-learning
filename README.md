# Demo — Go Monorepo
Production-grade applications and from-scratch implementations of CS fundamentals — all in one repo. Built to demonstrate **systems thinking**: from raw memory management and index internals to event-driven microservices with full-stack observability.
## Repository Map
```
golang/
├── apps/
│   ├── demo/           High-throughput comment service (Redis Streams, OTel, k6)
│   ├── database/       Storage engine from scratch (heap, indexes, query planner)
│   └── worker_basic/   Goroutine worker pool with graceful shutdown
├── education/
│   ├── structures/     Array (CGo malloc/free) & Hash Table (chaining, linear probing)
│   ├── datarace/       Data race demos: mutex vs channel-based locking
│   └── channels/       Channel semantics: buffering, close behavior, deadlocks
└── leetcode/
    └── tower_of_hanoi/ Classic recursion
```
## Applications
### [`apps/demo`](golang/apps/demo) — High-Throughput Comment Service
Event-driven write path: **HTTP → Redis Stream (TxPipeline) → Batch Consumer (XReadGroup) → PostgreSQL**.
- **3 000 rps** write throughput (k6 load tested)
- Redis Streams as write buffer with consumer groups (at-least-once, pending reprocessing)
- **OpenTelemetry** across all three pillars — traces, metrics, logs — plus Pyroscope continuous profiling
- HTTP metrics via `httpconv` semconv, Redis auto-instrumented via `redisotel`
- Business metrics: observable gauges pulled from PostgreSQL on scrape
- Infrastructure: vmagent remote-write, postgres_exporter, redis_exporter
  **Stack:** Go 1.24 · Gin · Redis 8 Streams · PostgreSQL 17 · GORM · OpenTelemetry SDK · Pyroscope · Google Wire · Cobra · k6 · Docker Compose
### [`apps/database`](golang/apps/database) — Storage Engine from Scratch
In-memory relational storage engine modeled after PostgreSQL internals.
- **Heap** storage with ctid addressing (generic `Heap[TRecord]`)
- **Pluggable indexes** via `Index` interface — same pattern as PostgreSQL access methods
- **Hash Index** with O(1) lookup, thread-safe, CTID-level deletion
- **Cost-based query planner** — analyzes index stats, picks cheapest access path
- **100K-record benchmarks** — SeqScan vs Hash across all field types
- Reference notes on B+Tree, GIN, BRIN, partial indexes, B-link-Tree concurrency
### [`apps/worker_basic`](golang/apps/worker_basic) — Goroutine Worker Pool
Minimal worker pool pattern: channel-based task distribution, context cancellation, graceful shutdown. Clean separation: `main` (orchestration) → `workers` (channel consumer) → `task` (business logic).
## Education
### [`structures/`](golang/education/structures) — Data Structures from Scratch
**Array** — generic fixed-size array on raw memory. `C.malloc` / `C.free` via CGo, `unsafe.Pointer` arithmetic, manual deallocation. Zero Go slices.
**Hash Table** — pluggable collision resolution:

| Strategy       | Implementation                              |
|----------------|---------------------------------------------|
| Chaining       | Per-bucket `sync.RWMutex` (sharded locking) |
| Linear Probing | Tombstone deletion, auto-grow ×2, rehash    |
Two hashers: production `maphash` and intentionally bad `FirstRuneReturn` for collision stress-testing. Concurrent correctness verified with 1 000 goroutines.
### [`datarace/`](golang/education/datarace) — Concurrency Primitives
Side-by-side comparison of `sync.RWMutex` vs channel-based locking for shared state protection. Safe vs unsafe read/write methods on the same struct.
### [`channels/`](golang/education/channels) — Channel Semantics
Interactive demo: buffered channel lifecycle, close behavior, deadlock conditions. What happens when you read from a closed channel vs an open empty one.
## License
MIT