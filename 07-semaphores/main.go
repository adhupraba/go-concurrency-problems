package main

import (
	"log"
	"sync"
	"time"
)

type Semaphore struct {
	sem chan struct{}
}

func (this *Semaphore) Acquire() {
	this.sem <- struct{}{}
}

func (this *Semaphore) Release() {
	<-this.sem
}

func NewSemaphore(maxConcurrency int) *Semaphore {
	return &Semaphore{
		sem: make(chan struct{}, maxConcurrency),
	}
}

func main() {
	semaphore := NewSemaphore(2)
	numChan := make(chan int, 5)
	wg := sync.WaitGroup{}

	defer close(numChan)

	for i := 1; i <= 15; i++ {
		semaphore.Acquire()
		wg.Add(1)

		go func() {
			time.Sleep(1 * time.Second)

			defer wg.Done()
			defer semaphore.Release()

			val := <-numChan
			log.Println("val:", val)
		}()

		numChan <- i
	}

	wg.Wait()
}
