package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	const newJobs = 5
	const newWorker = 3
	jobs := make(chan int, newJobs)
	results := make(chan int, newJobs)
	// Create context with cancel
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// start worker
	for w := 1; w <= newWorker; w++ {
		wg.Add(1)
		go Worker(ctx, w, jobs, results, &wg)
	}
	// send jobs
	for j := 1; j <= newJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Optional: Cancel context after 3 seconds (simulate early shutdown)
	// go func() {
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("Cancelling context...")
	// 	cancel()
	// }()

	// Wait for workers to complete
	wg.Wait()
	close(results)
	// Read results
	for result := range results {
		fmt.Println("Result:", result)
	}

	// Cancel context explicitly (clean-up)
	cancel()

}

func Worker(ctx context.Context, id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d recive cancellation signal\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return // jobs channel closed
			}
			fmt.Printf("worker %d start job %d\n", id, job)
			time.Sleep(time.Second)
			fmt.Printf("worker %d finished job %d\n", id, job)

			select {
			case results <- job * 2:
			case <-ctx.Done():
				fmt.Printf("worker %d could not send result due to cancellation\n", id)
				return
			}

		}
	}
}
