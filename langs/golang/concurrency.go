package main


import (
	"fmt"
	"sync"
	"time"
)

func concuurency() {
	fmt.Println("=== 1. Goroutines ===")
	// A goroutine is a lightweight thread managed by Go runtime
	go sayHello()
	time.Sleep(1 * time.Second) // wait for goroutine to finish

	fmt.Println("\n=== 2. Channels ===")
	// A channel is used to communicate between goroutines
	ch := make(chan string)
	go sendMessage(ch)
	msg := <-ch // receive from channel
	fmt.Println("Received from channel:", msg)

	fmt.Println("\n=== 3. Buffered Channels ===")
	// Buffered channel allows sending without waiting for a receiver
	bufCh := make(chan int, 2)
	bufCh <- 10
	bufCh <- 20
	fmt.Println("Buffered values:", <-bufCh, <-bufCh)

	fmt.Println("\n=== 4. Select Statement ===")
	// Select lets you wait on multiple channel operations
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch1 <- "message from ch1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "message from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}

	fmt.Println("\n=== 5. WaitGroup ===")
	// WaitGroup waits for a collection of goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2) // we're waiting for 2 goroutines

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 1 done")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Goroutine 2 done")
	}()

	wg.Wait()
	fmt.Println("All goroutines finished")
}

// sayHello runs as a goroutine
func sayHello() {
	fmt.Println("Hello from goroutine!")
}

// sendMessage sends a string to a channel
func sendMessage(ch chan string) {
	time.Sleep(500 * time.Millisecond)
	ch <- "Hi from channel"
}
