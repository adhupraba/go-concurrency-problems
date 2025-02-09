package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (this *Counter) mutexIncrement(wg *sync.WaitGroup) {
	defer wg.Done()

	this.mu.Lock()
	defer this.mu.Unlock()

	this.value++
}

func main() {
	counter := Counter{value: 0, mu: sync.Mutex{}}
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go counter.mutexIncrement(&wg)
	}

	wg.Wait()

	fmt.Println("mutex solution:", counter.value)
}
