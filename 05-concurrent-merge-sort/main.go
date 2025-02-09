package main

import (
	"fmt"
	"math/rand"
)

// threshold defines the minimum slice length to perform concurrent sorting.
// Below this threshold, the sort is done sequentially.
const SEQUENTIAL_SORT_THRESHOLD = 50

func main() {
	arr := generateRandomSlice(100)
	fmt.Printf("pre sort: %#v\n\n", arr)

	sorted := concurrentMergeSort(arr)
	fmt.Printf("post sort: %#v\n\n", sorted)
}

func generateRandomSlice(size int) []int {
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(5000)
	}

	return arr
}

func concurrentMergeSort(arr []int) []int {
	n := len(arr)

	if n <= 1 {
		return arr
	}

	// For small slices, fall back to sequential merge sort.
	if n < SEQUENTIAL_SORT_THRESHOLD {
		return sequentialMergeSort(arr)
	}

	mid := n / 2
	leftChan := make(chan []int, 1)
	rightChan := make(chan []int, 1)

	defer close(leftChan)
	defer close(rightChan)

	go func() {
		leftChan <- concurrentMergeSort(arr[:mid])
	}()

	go func() {
		rightChan <- concurrentMergeSort(arr[mid:])
	}()

	sortedLeft := <-leftChan
	sortedRight := <-rightChan

	return merge(sortedLeft, sortedRight)
}

func sequentialMergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	l := sequentialMergeSort(arr[0:mid])
	r := sequentialMergeSort(arr[mid:])

	return merge(l, r)
}

func merge(left, right []int) []int {
	m, n := len(left), len(right)
	res := make([]int, 0, m+n)
	i, j := 0, 0

	for i < m && j < n {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}
