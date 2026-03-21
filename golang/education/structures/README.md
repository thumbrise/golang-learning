# Data Structures — Low-Level Implementations in Go
From-scratch implementations of fundamental data structures with **manual memory management** (CGo malloc/free) and **pluggable collision resolution strategies**. No standard slices, no `map` — raw pointers, unsafe arithmetic, and explicit allocation.
## Array — Manual Memory via CGo
A generic fixed-size array that bypasses Go's runtime allocator entirely. Memory is allocated with `C.malloc` and freed with `C.free` — the same lifecycle as C/C++ heap objects.
```
  NewArray[T](5)
       │
       ▼
  C.malloc(5 × sizeof(T))  ──▶  raw memory block
       │
       ├── Get(i)  →  *((*T)(base + i × sizeof(T)))     pointer arithmetic
       ├── Set(i,v) → *((*T)(base + i × sizeof(T))) = v
       └── Clear()  →  C.free(ptr)                       manual deallocation
```
**Key details:**
- **Zero Go slices/arrays** — addressing is done via `unsafe.Pointer` + `uintptr` arithmetic
- **Generic** — `Array[T any]` works with `int`, `string`, structs, anything
- **Bounds checking** — panics with descriptive message on out-of-range access
- **Explicit lifecycle** — caller must call `Clear()` to avoid memory leaks (no GC)
- **Size tracking** — `SizeBytes()` reports exact allocation size
### Why malloc?
Go's `unsafe.Pointer` lets you do pointer arithmetic, but the GC still manages the backing memory. By using `C.malloc`, the array lives completely outside Go's heap — it demonstrates how real allocators work, how `unsafe.Sizeof` maps to memory layout, and why manual deallocation matters.
## Hash Table — Pluggable Collision Strategies
A generic hash table with **strategy pattern** for both hashing and collision resolution:
```
                    ┌───────────────────┐
  Set/Get/Delete ──▶│    HashTable[T]    │
                    │  hasher: Hasher    │
                    │  store:  Store[T]  │
                    └────────┬──────────┘
                             │
              ┌──────────────┼──────────────┐
              ▼              ▼              ▼
      ┌──────────┐   ┌────────────┐   ┌──────────┐
      │  Chain   │   │ Linear     │   │  (next)  │
      │ Buckets  │   │ Probing    │   │          │
      │ RWMutex  │   │ Tombstones │   │          │
      │ per-buck │   │ Auto-grow  │   │          │
      └──────────┘   └────────────┘   └──────────┘
```
### Collision Resolution Strategies
| Strategy | How it works | Trade-offs |
|---|---|---|
| **Chaining** (`chain`) | Each bucket is a linked list with `sync.RWMutex` | Simple, stable under high load factor, per-bucket locking |
| **Open Addressing / Linear Probing** (`linearprob`) | Probe sequentially until free slot; tombstones on delete; auto-grow ×2 | Cache-friendly, but clustering under bad hash; needs grow |
### Hashers
Pluggable via the `Hasher` interface (`Hash(key string) uint64`):
- **`MapHashHasher`** — production-grade `hash/maphash`, good distribution
- **`FirstRuneReturnHasher`** — intentionally terrible (returns first rune as hash), used to stress-test collision handling
### Concurrency
Both stores are thread-safe:
- **Chain** — `sync.RWMutex` per bucket (sharded locking)
- **Linear Probing** — global `sync.RWMutex` (simpler but coarser)
  Concurrent correctness is verified with 1 000 parallel goroutines writing unique keys.
### Benchmarks
Matrix benchmark: `{table_size} × {fill_level} × {hasher} × {strategy}`:
```bash
go test -bench=. -benchmem ./hashtable/...
```
Measures insert latency after pre-filling the table — shows how collision rate and fill factor affect performance with good vs. bad hash functions.
## Project Structure
```
structures/
├── array/
│   ├── array.go            # Generic array with unsafe.Pointer arithmetic
│   ├── allocator.go        # CGo malloc/free allocator
│   ├── array_test.go       # Bounds, types (int/string/struct), lifecycle
│   └── todo.md             # Roadmap: Slice, SyncArray, iterators
├── hashtable/
│   ├── hashtable.go        # Facade: pluggable hasher + store
│   ├── hasher.go           # Hasher interface
│   ├── hashers/
│   │   ├── maphash_hasher.go         # Production hasher (hash/maphash)
│   │   └── first_rune_return_hasher.go  # Bad hasher (collision demo)
│   ├── store/
│   │   ├── store.go        # Store interface (Set/Get/Delete)
│   │   ├── item.go         # Item interface
│   │   ├── hasheditem.go   # Concrete item (key + hash + value)
│   │   ├── zeroitem.go     # Null object (empty slot)
│   │   └── types/
│   │       ├── chain/
│   │       │   ├── store.go    # Chaining strategy
│   │       │   └── bucket.go   # Per-bucket RWMutex linked list
│   │       └── openaddr/
│   │           └── linearprob/
│   │               ├── store.go         # Linear probing + auto-grow
│   │               ├── tombstoneitem.go # Tombstone for lazy deletion
│   │               └── store_test.go
│   ├── hashtable_test.go   # Set/Get/Delete, overwrite, concurrent, benchmarks
│   └── todo.md             # Roadmap: xxhash, Robin Hood, cuckoo hashing
```
## Concepts Demonstrated
| Concept | Where |
|---|---|
| Manual memory management (malloc/free) | `array/allocator.go` |
| Pointer arithmetic (`unsafe.Pointer` + `uintptr`) | `array/array.go:64` |
| Strategy pattern (pluggable stores & hashers) | `hashtable/hashtable.go` |
| Sharded locking (per-bucket mutex) | `chain/bucket.go` |
| Tombstone deletion in open addressing | `linearprob/tombstoneitem.go` |
| Dynamic resizing (grow ×2 + rehash) | `linearprob/store.go:105` |
| Null Object pattern | `store/zeroitem.go` |
| Generics (`[T any]`) across all structures | everywhere |
| CGo interop | `array/allocator.go` |
## License
MIT