package main

import (
	"fmt"
	"sync"
)

func main() {
	WithoutmutexrRaceCondication()
}

func WithoutmutexrRaceCondication() {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i <= 1000; i++ {
		wg.Add(1)

		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

func UseMutexForPreventRaceCondicaton() {
	counter := 0
	var wg sync.WaitGroup
	var mg sync.Mutex

	for i := 0; i <= 1000; i++ {
		wg.Add(1)

		go func() {
			mg.Lock()
			counter++
			mg.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
