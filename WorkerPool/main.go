package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const newJob = 5
	const newWorker = 3
	var wg sync.WaitGroup
	jobs := make(chan int, newJob)
	results := make(chan int, newJob)

	// Start worker goroutines
	for w := 1; w <= newWorker; w++ {
		wg.Add(1)
		go Worker(w, jobs, results, &wg)
	}

	// send jobs
	for j := 1; j <= newJob; j++ {
		jobs <- j
	}

	close(jobs) // no more jobs
	// Wait for all workers to finish
	wg.Wait()
	close(results)
	// Read results
	for result := range results {
		fmt.Println("Result:", result)
	}

}

func Worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // simulate work
		fmt.Printf("worker %d finished job %d\n", id, job)
		result <- job * 2
	}

}
