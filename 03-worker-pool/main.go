package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numWorkers := 5
	taskChan := make(chan int, numWorkers)
	resultChan := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, taskChan, resultChan, &wg)
	}

	go giveWork(taskChan, 20)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Printf("received doubled num: %d\n", result)
	}
}

func worker(id int, taskChan <-chan int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		time.Sleep(1 * time.Second)
		fmt.Printf("(%d) doubling number %d\n", id, task)
		resultChan <- task + task
	}
}

func giveWork(taskChan chan<- int, workCount int) {
	for i := 0; i < workCount; i++ {
		taskChan <- i + 1
	}

	close(taskChan)
}
