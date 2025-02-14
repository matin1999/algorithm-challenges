package main

import (
	"fmt"
	"sync"
)

func sumRange(a, b int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done()
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	resultChan <- sum
}

func main() {
	a := 1
	b := 100
	n := 5

	var wg sync.WaitGroup
	resultChan := make(chan int, n)

	rangeSize := (b - a + 1) / n
	for i := 0; i < n; i++ {
		start := a + i*rangeSize
		end := start + rangeSize - 1

		if i == n-1 {
			end = b
		}

		if start <= end {
			wg.Add(1)
			go sumRange(start, end, &wg, resultChan)
		}
	}

	wg.Wait()
	close(resultChan)

	totalSum := 0
	for sum := range resultChan {
		totalSum += sum
	}

	fmt.Printf("Sum of numbers from %d to %d is: %d\n", a, b, totalSum)
}
