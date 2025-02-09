package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	numProducers := 5
	numConsumers := 2

	wg := sync.WaitGroup{}
	ch := make(chan int, 3)

	defer close(ch)

	for i := 0; i < numProducers; i++ {
		count := 5
		wg.Add(count)
		go producer(i+1, count, ch)
	}

	for i := 0; i < numConsumers; i++ {
		go consumer(i+1, ch, &wg)
	}

	wg.Wait()
}

func producer(id int, count int, ch chan<- int) {
	for i := 1; i <= count; i++ {
		num := i * count * rand.Intn(3000)
		fmt.Printf("(%d) sending num %d into ch\n", id, num)
		ch <- num
	}
}

func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		fmt.Printf("(%d) received num: %d\n", id, num)
		wg.Done()
	}
}
