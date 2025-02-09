# Worker Pool

### Problem Statement:

Implement a worker pool that processes a stream of tasks concurrently. The main goroutine should send tasks (e.g., integers representing work units) to a task channel. A fixed number of worker goroutines should pull tasks from the channel, process them (simulate work with a sleep or computation), and then send results on a results channel.

### Key Concepts:

- Use channels to distribute work.
- Use goroutines to implement workers.
- Use sync.WaitGroup to wait for all workers to complete.
- Handle closing channels and ensure no goroutine is left hanging.

### Hints:

- Create a function for your worker that runs in a loop reading from the task channel.
- Use a WaitGroup to wait for all workers to finish once the task channel is closed.
- Make sure to close the results channel when all workers have completed their work.
