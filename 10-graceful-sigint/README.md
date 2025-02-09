# Graceful SIGINT Shutdown for a Long-Running Process

### Problem Statement:

Implement a graceful shutdown for a process that runs indefinitely. On receiving the first SIGINT (Ctrl-C), attempt a graceful stop by calling `Stop()`. If a second SIGINT is received before the process shuts down, force an immediate termination.

### Key Concepts:

- Handling OS signals with the os/signal package.
- Implementing graceful shutdowns with timeouts.
- Using goroutines and select for concurrent event handling.

### Hints:

- Use a channel to receive SIGINT signals.
- Launch the process in its own goroutine and call `Stop()` on the first SIGINT.
- Use a select statement with a timeout (e.g., time.After) and a secondary signal read to decide between waiting for graceful shutdown or forcing termination.

### Example:

```go
// MockProcess for example
type MockProcess struct {
	mu        sync.Mutex
	isRunning bool
}

// Run will start the process
func (m *MockProcess) Run() {
	m.mu.Lock()
	m.isRunning = true
	m.mu.Unlock()

	fmt.Print("Process running..")
	for {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
}

// Stop tries to gracefully stop the process, in this mock example
// this will not succeed
func (m *MockProcess) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.isRunning {
		log.Fatalln("Cannot stop a process which is not running")
	}

	fmt.Println("\nStopping process..")
	for {
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// implement graceful shutdown of the mock process
}

```
