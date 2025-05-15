package main

import (
	"fmt"
	"maths/question"
)

func main() {
	x := question.MinimumSumOf4Digit(2932)
	fmt.Println("MinimumSumOf4Digit :", x)
	dayOfmonth := question.DayOfYear("2004-03-01")
	fmt.Println("DayOfYear :", dayOfmonth)
	SubtractProductAndSum := question.SubtractProductAndSum(234)
	fmt.Println("SubtractProductAndSum :", SubtractProductAndSum)
}
