# Concurrent Odd and Even Number Processing Using Channels

### Problem Statement:

Design a concurrent system where one goroutine processes odd numbers and another processes even numbers from a shared channel. Each number is handled by the appropriate goroutine, which signals completion before the next number is sent.

### Key Concepts:

- Using channels for inter-goroutine communication.
- Filtering data based on conditions (odd vs. even).
- Synchronizing tasks with a signaling channel.

### Hints:

- Use one channel to pass numbers and another to signal when a number has been processed.
- Utilize a WaitGroup to ensure all goroutines finish before program exit.
