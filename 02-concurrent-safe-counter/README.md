# Concurrent Safe Counter

### Problem Statement:

Design and implement a counter that can be safely incremented and read from multiple goroutines concurrently. Implement at least two versions: one using sync.Mutex (or sync.RWMutex) and one using the sync/atomic package.

### Key Concepts:

- Protect shared state using mutexes or atomic operations.
- Compare the performance and code complexity of both approaches.

### Hints:

For the mutex version, embed a mutex in your counter struct and lock/unlock around increments and reads.
For the atomic version, use functions from the sync/atomic package (e.g., atomic.AddInt64).
