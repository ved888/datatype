package syncs

import "fmt"

func EvenOdd(n int) {
	odd := make(chan int)
	even := make(chan int)

	go func() {
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(odd)
		close(even)
	}()

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			val := <-even
			fmt.Println("even :", val)
		} else {
			val := <-odd
			fmt.Println("odd :", val)
		}
	}
}

func OddEven(n int) {
	oddChan := make(chan int)
	evenChan := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i <= n; i++ {
			if i%2 == 1 {
				oddChan <- i
			} else {
				evenChan <- i
			}
		}
		close(oddChan)
		close(evenChan)
	}()

	go func() {
		for {
			val, ok := <-oddChan
			if !ok {
				break
			}
			fmt.Println("odd:", val)

			val, ok = <-evenChan
			if ok {
				fmt.Println("even:", val)
			}
		}
		done <- true
	}()

	<-done
}

func EvenOddInSenDReciveWay(n int) {
	even := make(chan int)
	odd := make(chan int)
	signal := make(chan bool)

	go func() {
		for i := 1; i <= n; i++ {
			if i%2 == 0 {
				even <- i
				fmt.Println("send even :", i)
				<-signal
			} else {
				odd <- i
				fmt.Println("send odd :", i)
				<-signal
			}
		}
		close(even)
		close(odd)
	}()

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			val := <-even
			fmt.Println("recive even :", val)
			signal <- true
		} else {
			fmt.Println("recive odd :", <-odd)
			signal <- true

		}
	}

}
