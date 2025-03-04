package main

import (
	"strconv"
)

func ReverseOfDigit(num int) int {
	var ans int

	for num != 0 {
		rem := num % 10
		ans = ans*10 + rem
		num = num / 10
	}
	return ans
}

func SumOfDigite(n int) int {
	var ans int
	for n != 0 {
		rem := n % 10
		ans = ans + rem
		n = n / 10
	}
	return ans
}

func Ispalindrome(n int) bool {

	str := strconv.Itoa(n)
	reverse := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return str == reverse
}

func Ispalindrome1(n int) bool {
	origanal := n
	reverse := 0

	for n != 0 {
		rem := n % 10
		reverse = reverse*10 + rem
		n = n / 10
	}
	if origanal == reverse {
		return true
	}
	return false
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func GCD1(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	num1, num2 := a, b

	for b != 0 {
		a, b = b, a%b
	}
	gcd := a

	lcm := num1 * num2 / gcd
	return lcm
}
