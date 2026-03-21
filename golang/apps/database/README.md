# Database Internals — Storage Engine from Scratch in Go
Educational implementation of a **relational storage engine** in pure Go — heap storage, index structures, cost-based query planner, and benchmarks. Built to deeply understand how PostgreSQL works under the hood, not to replace it.
## What's Inside
```
                    ┌─────────────────────────────────┐
                    │          Storage Engine          │
                    │                                  │
  SearchEqual ──────▶  Planner (cost-based)            │
  SearchRange       │    │                             │
  SearchPrefix      │    ▼ analyze(index, conditions)  │
                    │  ┌──────────┐  ┌──────────────┐  │
                    │  │  SeqScan │  │  Hash Index   │  │
                    │  │  O(N)    │  │  O(1) lookup  │  │
                    │  └────┬─────┘  └──────┬───────┘  │
                    │       │    ctids[]     │          │
                    │       ▼               ▼          │
                    │  ┌──────────────────────────┐    │
                    │  │   Heap (map[ctid]Record)  │    │
                    │  └──────────────────────────┘    │
                    └─────────────────────────────────┘
```
### Heap
In-memory page heap modeled after PostgreSQL's heap storage. Records are addressed by `ctid` (tuple identifier). The heap is generic — `Heap[TRecord Record]` works with any type implementing the `Record` interface (`PK()`, `Fields()`, `Get()`, `ToMap()`).
### Indexes
Pluggable index system via the `Index` interface — same approach PostgreSQL uses for its access methods (Hash, B-Tree, GIN, BRIN):
```go
type Index interface {
    Insert(ctid, fieldName, value string)
    Search(fieldName, value string) []string
    Delete(fieldName, value string)
    DeleteCTID(ctid, fieldName, value string)
    Update(ctid, fieldName, oldValue, newValue string)
    Stats() *IndexStats
}
```
**Implemented:**
- **Hash Index** — O(1) equality lookup, thread-safe (`sync.RWMutex`), with CTID-level deletion and cleanup of empty buckets.
- **Sequential Scan** — full heap scan as the fallback path, cost = `1 × total_rows`.
  **Planned:** B-Tree (range queries, ORDER BY), GIN (arrays, full-text), BRIN (time-series, monotonic fields).
### Query Planner
Cost-based planner that chooses the cheapest access path for a given set of conditions:
1. Collects `[]Condition` (field, value, operator).
2. For each registered index, calls `analyze()` to get `{Cost, Rows}`.
3. Picks the index with the lowest cost.
4. Falls back to SeqScan if no index matches.
   The planner is designed for incremental evolution — from simple cost comparison to selectivity estimation (NDV, HyperLogLog, Count-Min Sketch, histograms). See the full roadmap in [`storage/planner/todo.md`](storage/planner/todo.md).
### Benchmarks
Comparative benchmarks on 100K records — SeqScan vs Hash Index across all field types:
```bash
go test -bench=. -benchmem ./...
```
The benchmark matrix covers: `id`, `email`, `age`, `favorite_colors`, `last_access_time` × `{SeqScan, Hash}` × `SearchEqual`. Planned: range queries, prefix search, multi-condition queries.
## Index Theory — Reference Notes
The [`indexes/readme.md`](indexes/readme.md) contains detailed notes on PostgreSQL index internals:
| Index | When to use | Complexity |
|---|---|---|
| **Hash** | Millions of point lookups (`=`), no ranges | O(1) avg |
| **B+Tree** | Ranges, sorting, uniqueness — the universal default | O(log N) |
| **GIN** | Full-text, arrays, JSONB (`@>`, `&&`, `@@`) | O(log K) + posting list |
| **BRIN** | Huge monotonic tables (logs, time-series) | O(1) per block range |
Plus: partial indexes, B-link-Tree concurrent access, parallel scan strategies, `fastupdate` / pending list in GIN.
## Project Structure
```
database/
├── indexes/
│   ├── hash/               # Hash index implementation (thread-safe)
│   └── readme.md           # PostgreSQL index internals reference
├── storage/
│   ├── core/
│   │   ├── heap.go         # Generic heap storage (map[ctid]Record)
│   │   └── record.go       # Record interface
│   ├── search/
│   │   ├── index.go        # Index interface (pluggable access methods)
│   │   ├── indexer.go       # Index builder
│   │   ├── seqsearcher.go  # Sequential scan (fallback)
│   │   └── condition.go    # Query conditions
│   ├── planner/
│   │   ├── planner.go      # Cost-based query planner
│   │   ├── analyser.go     # Cost analysis
│   │   └── todo.md         # Planner evolution roadmap
│   ├── stats/              # Index statistics
│   └── storage.go          # Storage engine facade
├── test/
│   ├── dal/                # Test record types (User)
│   └── fixtures/           # Data generators (100K users)
└── indexes_bench_test.go   # Comparative benchmarks
```
## Why
Understanding database internals is the difference between "add an index" and knowing *which* index, *why*, and *what it costs*. This project is a hands-on exploration of:
- How heap storage and ctid addressing work
- Why the planner picks SeqScan over an index (and when it shouldn't)
- The real cost difference between O(N) and O(1) on 100K+ rows
- How PostgreSQL's pluggable index API is designed
- Selectivity estimation and cost models
## License
MIT