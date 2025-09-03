package main

import (
	"fmt"
	"sync"
)

func main() {
	RunNGoroutines(5)
}

// Run N Goroutines and Wait using sync.WaitGroup
func RunNGoroutines(n int) {
	var wg sync.WaitGroup

	wg.Add(n)

	for i := 0; i <= n; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Println("Goroutine", id, "finished")
		}(i)
	}
	wg.Wait()
}
