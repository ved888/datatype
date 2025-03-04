package main

import "fmt"

func main() {
	Fibonacci(9)
	fib := FibonacciByRecursion(9)
	fmt.Println("fibonacci searies", fib)
	fib1 := FibonacciDP(10)
	fmt.Println("fibonacci searies:", fib1)
}
