package main

import (
	"fmt"
	"sync"
)

func main() {
	numChan := make(chan int)
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(2)

	go odd(numChan, done, &wg)
	go even(numChan, done, &wg)

	for i := 1; i <= 10; i++ {
		numChan <- i
		<-done
	}

	close(numChan)
	close(done)

	wg.Wait()
}

func odd(numChan chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range numChan {
		if num%2 != 0 {
			fmt.Println("odd received:", num)
			done <- struct{}{}
		} else {
			numChan <- num
		}
	}
}

func even(numChan chan int, done chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range numChan {
		if num%2 == 0 {
			fmt.Println("even received:", num)
			done <- struct{}{}
		} else {
			numChan <- num
		}
	}
}
