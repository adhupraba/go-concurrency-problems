package main

import (
	"fmt"
)

func main() {
	numChan := make(chan int)
	evenChan := make(chan int)
	squareChan := make(chan int)
	resultChan := make(chan int)

	go generateNums(numChan)
	go filterEvenNums(numChan, evenChan)
	go squareNum(evenChan, squareChan)
	go pipeResult(squareChan, resultChan)

	for result := range resultChan {
		fmt.Println("result:", result)
	}
}

func generateNums(numChan chan<- int) {
	for i := 0; i <= 100; i++ {
		numChan <- i
	}

	close(numChan)
}

func filterEvenNums(numChan <-chan int, evenChan chan<- int) {
	for num := range numChan {
		if num%2 == 0 {
			evenChan <- num
		}
	}

	close(evenChan)
}

func squareNum(evenChan <-chan int, squareChan chan<- int) {
	for num := range evenChan {
		squareChan <- num * num
	}

	close(squareChan)
}

func pipeResult(squareChan <-chan int, resultChan chan<- int) {
	for num := range squareChan {
		resultChan <- num
	}

	close(resultChan)
}
