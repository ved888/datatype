package main

import (
	"fmt"
	"numberTheory/questionvise"
)

func main() {
	x := questionvise.SieveOfEratosthenes(20)
	fmt.Println("prime number :", x)
	isFactorial := questionvise.IsFactorial(10)
	fmt.Println("IsFactorial :", isFactorial)
	gcd := questionvise.GCD(14, 8)
	fmt.Println("GCD :", gcd)
}
