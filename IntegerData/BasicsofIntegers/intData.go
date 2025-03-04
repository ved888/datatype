package main

import "fmt"

func ExceedTheMaximumValue() {
	var x int8 = 127 // Maximum value for int8
	fmt.Println(x)
	x++
	fmt.Println(x) // This will wrap around to -128
}
