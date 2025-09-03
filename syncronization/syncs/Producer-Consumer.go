package syncs

import (
	"fmt"
	"time"
)

func producer(jobs chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println("Produced:", i)
		jobs <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(jobs)
}

func consumer(jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		fmt.Println("Consumed:", job)
		time.Sleep(500 * time.Millisecond)
	}
	done <- true
}

func StartProducerConsumer(count int) {
	jobs := make(chan int, 5) // Buffered job queue
	done := make(chan bool)   // Completion signal

	go producer(jobs, count)
	go consumer(jobs, done)

	<-done
	fmt.Println("All jobs processed.")
}
