package syncs

import (
	"fmt"
	"sync"
)

func SyncSendReceive(n int) {
	send := make(chan int)
	ack := make(chan bool)

	go func() {
		for i := 1; i <= n; i++ {
			send <- i // Block until value is received
			fmt.Println("send :", i)
			<-ack // Wait for acknowledgment from receiver
		}
		close(send)
	}()

	for val := range send {
		fmt.Println("received :", val)
		ack <- true // Acknowledge to sender that value is received

	}

}

// use with waitgroup
func SyncSendReceiveWg(n int) {
	var wg sync.WaitGroup
	send := make(chan int)

	wg.Add(1)
	// Sender
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i++ {
			fmt.Println("Sent:", i)
			send <- i
		}
		close(send)
	}()

	// Receiver
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range send {
			fmt.Println("Received:", val)
		}
	}()

	wg.Wait()
}

func ConcurrentSendReceive(n int) {
	data := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2) // one for sender, one for receiver

	// Sender Goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i++ {
			fmt.Println("send message:", i)
			data <- i
		}
		close(data) // signal that no more data is coming
	}()

	// Receiver Goroutine
	go func() {
		defer wg.Done()
		for val := range data {
			fmt.Println("recive message:", val)
		}
	}()

	wg.Wait()
}
