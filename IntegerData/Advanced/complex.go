package main

import "fmt"

func Fibonacci(n int) {
	if n < 0 {
		fmt.Errorf("give positive number")
		return
	}
	a, b := 0, 1
	fmt.Printf("febonasi number")
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		next := a + b
		a = b
		b = next
	}
	fmt.Println()

}

func FibonacciDP(n int) []int {
	if n < 1 {
		return []int{}
	}
	fib := make([]int, n)
	fib[0], fib[1] = 0, 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

func FibonacciByRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciByRecursion(n-1) * FibonacciByRecursion(n-2)
}
