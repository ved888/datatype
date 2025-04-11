package main

import "fmt"

func SafeFunc() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("About to panic...")
	panic("Something went wrong!")
}

func Add1(a, b int) int {
	return a + b
}

func main() {
	SafeFunc()
	fmt.Println("server is recover")
}
