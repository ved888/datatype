package main

import "fmt"

func main() {
	x := SumOfDigite(123)
	fmt.Println("sum of digit:", x)
	x = ReverseOfDigit(123)
	fmt.Println("reverse of digit:", x)

	num := 121
	if Ispalindrome(num) {
		fmt.Printf("%d is a palindrome\n", num)
	} else {
		fmt.Printf("%d is not a palindrome\n", num)
	}
	bool := Ispalindrome1(num)
	fmt.Println("num is palidrom:", bool)

	gcd := GCD(24, 15)
	fmt.Println("gcd", gcd)
	gcd1 := GCD1(15, 24)
	fmt.Println("gcd1", gcd1)
	lcm := LCM(12, 20)
	fmt.Println("lcm", lcm)
}
