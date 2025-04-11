package main

import (
	"fmt"
)

func main() {
	DeclareAChannelUnbuffer1()
	DeclareAChannelUnbuffer2()
	DeclareAChannelBuffer1()
	DeclareAChannelBuffer2()
	// DeadlockWithUnbufferWhenNoGorutineReadyForRecive()
	DeadlockWithUnbufferWhenNoGorutineReadyForReciveFix()
	// DeadlockWithUnbufferWhenNoGorutineReadyForSend()
	DeadlockWithUnbufferWhenNoGorutineReadyForSendFix()
	// DeadlockWithBufferWhenBufferIsFull()
	DeadlockWithBufferWhenBufferIsFullFix()
	// DeadlockWithBufferWhenBufferIsEmpty()
	DeadlockWithBufferWhenBufferIsEmptyFix()
	// UsingRangeOnAChannelThatIsNeverClosed()
	UsingRangeOnAChannelThatIsClosedFix()
	// NilChannel()
	CheckChannelCloseOrNot()
	ReceiveDataFromAClosedChannel()
	ch := make(chan int)
	go ChannelsBePassedAsFunctionArguments(ch)
	ReceiveData(ch)
	BufferChannelWith2Cap()
	SyncTwoChannelAndPrintInAcending()
	go SendOnlyChannel(ch) // Sending
	ReceiveOnlyChannel(ch) // Receiving
}

func DeclareAChannelUnbuffer1() {
	var ch chan int // declare

	ch = make(chan int) // initialize

	go func() {
		ch <- 5 // send
	}()
	value := <-ch // receive
	fmt.Println("receive data from channel :", value)
}

func DeclareAChannelUnbuffer2() {
	ch := make(chan int)

	go func() {
		ch <- 50
	}()
	value := <-ch
	fmt.Println("receive data from channel :", value)
}

func DeclareAChannelBuffer1() {
	var ch chan int

	ch = make(chan int, 5)
	ch <- 55
	value := <-ch
	fmt.Println("receive data from buffer channel :", value)
}

func DeclareAChannelBuffer2() {
	ch := make(chan int, 5)
	ch <- 89
	value := <-ch
	fmt.Println("receive data from buffered channel :", value)
}

func DeadlockWithUnbufferWhenNoGorutineReadyForRecive() {
	ch := make(chan int)
	ch <- 9 // 🔴 This will block forever because no one is receiving
	fmt.Println("receive data from unbuffere", ch)
}

// fix
func DeadlockWithUnbufferWhenNoGorutineReadyForReciveFix() {
	ch := make(chan int)
	go func() {
		ch <- 9
	}()
	fmt.Println("receive data from unbuffere", <-ch)
}

func DeadlockWithUnbufferWhenNoGorutineReadyForSend() {
	ch := make(chan int)
	value := <-ch // 🔴 This blocks forever because no one is sending
	fmt.Println("no reciver for ready", value)
}

// fix
func DeadlockWithUnbufferWhenNoGorutineReadyForSendFix() {
	ch := make(chan int)
	go func() {
		ch <- 42 // sender runs in goroutine
	}()
	value := <-ch
	fmt.Println("Received from unbuffered:", value)
}

func DeadlockWithBufferWhenBufferIsFull() {
	ch := make(chan int, 2)
	ch <- 9
	ch <- 5
	ch <- 7 // here occure deadlock because buffer is full
	fmt.Println("receive data from buffer :", <-ch)
}

func DeadlockWithBufferWhenBufferIsFullFix() {
	ch := make(chan int, 2)
	ch <- 9
	ch <- 5
	go func() {
		ch <- 7 // ✅ this won't block since another goroutine is reading
	}()
	fmt.Println("receive data from buffer :", <-ch)
}

func DeadlockWithBufferWhenBufferIsEmpty() {
	ch := make(chan int, 2)

	fmt.Println("receive data from buffer :", <-ch)
}

func DeadlockWithBufferWhenBufferIsEmptyFix() {
	ch := make(chan int, 2)
	go func() {
		ch <- 100
	}()

	fmt.Println("receive data from buffer :", <-ch)
}

func UsingRangeOnAChannelThatIsNeverClosed() {
	ch := make(chan int)
	go func() {
		ch <- 9
		ch <- 6
		ch <- 10
	}()

	for val := range ch {
		fmt.Println("receive data from range loop", val)
	}
}

func UsingRangeOnAChannelThatIsClosedFix() {
	ch := make(chan int)
	go func() {
		ch <- 3
		ch <- 10
		ch <- 9
		close(ch)
	}()
	for val := range ch {
		fmt.Println("receive value withy loop", val)
	}
}

func NilChannel() {
	var ch chan int // nil channel

	ch <- 10 // ❌ DEADLOCK
	fmt.Println("receive dat from nil channel :", ch)
}

func CheckChannelCloseOrNot() {
	ch := make(chan int)
	close(ch)
	val, ok := <-ch
	// ch <- 10 // ❌ PANIC: send on closed channel

	if !ok {
		fmt.Println("Channel is closed")
	} else {
		fmt.Println("Received:", val)
	}
}

func ReceiveDataFromAClosedChannel() {
	ch := make(chan int)
	close(ch)
	val, ok := <-ch
	fmt.Println("receive data from close channel :", val, ok)
}

func ChannelsBePassedAsFunctionArguments(ch chan int) {
	ch <- 10
}

func ReceiveData(ch chan int) {
	val := <-ch
	fmt.Println("Received:", val)
}

func BufferChannelWith2Cap() {
	ch := make(chan int, 2)

	for i := 0; i < 6; i++ {
		go func(val int) {
			ch <- 1 + val
		}(i)
	}
	for i := 0; i < 6; i++ {
		fmt.Println("return multiple time :", <-ch)
	}
	close(ch)
}

func SyncTwoChannelAndPrintInAcending() {
	odd := make(chan int)
	even := make(chan int)

	// ✅ One goroutine to send both odd and even numbers
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				even <- i
			} else {
				odd <- i
			}
		}
		close(odd)
		close(even)
	}()

	// ✅ Receive and print in ascending order
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even:", <-even)
		} else {
			fmt.Println("odd :", <-odd)
		}
	}
}

func SendOnlyChannel(ch chan<- int) {
	ch <- 555555 // ✅ Can only send
}

func ReceiveOnlyChannel(ch <-chan int) {
	val := <-ch // ✅ Only receiving allowed
	fmt.Println("this channel only Received:", val)
}
