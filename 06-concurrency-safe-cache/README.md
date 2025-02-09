# Concurrency-Safe Cache

### Problem Statement:

Create a concurrency-safe cache that supports concurrent read and write operations. The cache should support basic operations like Get, Set, and optionally Delete. Use either sync.RWMutex or channels to manage access to the underlying data structure (e.g., a map).

### Key Concepts:

- Protecting shared mutable state.
- Use of read-write locks to allow concurrent reads.
- Optionally, consider using Go channels to serialize access to the map.

### Hints:

- With sync.RWMutex, use RLock/RUnlock for reads and Lock/Unlock for writes.
- Alternatively, implement a cache manager goroutine that listens on a set of request channels (one for get and one for set) and serializes access to the map.
