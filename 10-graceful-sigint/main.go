package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	// Create a process
	proc := MockProcess{}

	// Run the process (blocking)
	go proc.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sig := <-c
	log.Printf("\ncaptured sigint %v, stopping application and exiting...\n", sig)

	go proc.Stop()

	// Wait for either graceful stop or second SIGINT
	select {
	case <-time.After(10 * time.Second):
		log.Println("Graceful shutdown completed.")
	case sig := <-c:
		log.Printf("Captured %v again, forcing shutdown...\n", sig)
	}

	os.Exit(0)
}

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
