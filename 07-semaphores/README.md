# Implementing a Semaphore Using Channels

### Problem Statement:

Implement a semaphore in Go using channels. A semaphore is a concurrency primitive that limits the number of goroutines that can access a particular resource. Your semaphore should have methods for acquiring and releasing a token.

### Key Concepts:

- Using buffered channels to simulate semaphores.
- Ensuring correct release of tokens.
- Preventing goroutines from deadlocking.

### Hints:

- Create a buffered channel with capacity equal to the maximum number of concurrent accesses allowed.
- The acquire operation (Acquire()) should block until a token is available.
- The release operation (Release()) should return a token to the channel.
