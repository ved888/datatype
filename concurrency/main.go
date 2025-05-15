package main

import (
	"fmt"
	"sync"
)

// Can you create a simple example where you spawn 5 goroutines and wait for them to finish using a WaitGroup?

func main() {
	var wg sync.WaitGroup

	wg.Add(5)
	go PrintHello(&wg)
	go PrintHello(&wg)
	go PrintHello(&wg)
	go PrintHello(&wg)
	wg.Wait()
	fmt.Println("Completed!")

}

func PrintHello(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("hello world")
}

// 🔁 Goroutines & WaitGroup
// What is a goroutine in Go, and how is it different from a thread?

// How do you wait for multiple goroutines to finish execution? Can you explain how sync.WaitGroup works?

// What happens if you call wg.Done() more times than wg.Add()?

// Can you create a simple example where you spawn 5 goroutines and wait for them to finish using a WaitGroup?

// 🔄 Channels
// What are buffered and unbuffered channels in Go? When would you use each?

// What happens if you try to send data on a channel but no one is receiving?

// How can you use channels to implement a simple worker pool pattern?

// What’s the purpose of the select statement in Go? Can you give an example?

// 🔒 Mutex
// Why would you use a sync.Mutex instead of channels? Can you show a simple use case where a mutex is needed?

// What happens if you forget to unlock a mutex? How can defer help avoid that mistake?
