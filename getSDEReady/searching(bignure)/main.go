package main

import (
	"fmt"
	"searching/question"
)

func main() {
	num := question.AnySqure(8)
	fmt.Println("return round down squre root :", num)
	result := question.IntersectionOfTwoArraysII([]int{2, 2}, []int{1, 2, 1, 2})
	fmt.Println("IntersectionOfTwoArraysII :", result)
	rotate := question.RotateString("abcde", "cdeab")
	fmt.Println("RotateString :", rotate)
}
