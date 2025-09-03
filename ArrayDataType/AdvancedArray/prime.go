package advancedarray

import (
	"fmt"
)

func IsPrimeNumber2(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// For small datasets: First approach (direct prime check) is better.
// printPrimes takes an array and prints prime numbers
func PrintPrimes(arr []int) {
	fmt.Println("Prime numbers in the array:")
	for _, num := range arr {
		if num < 2 {
			continue
		}
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(num, " ")
		}
	}
	fmt.Println()
}

// For large datasets: Second approach (Sieve of Eratosthenes) is better due to its efficient prime computation.
// printPrimes filters and prints prime numbers from an array
func PrintPrimesWithLargDataset(arr []int) {
	// Find the maximum number in the array to define sieve size
	maxNum := 0
	for _, num := range arr {
		if num > maxNum {
			maxNum = num
		}
	}

	// Create a boolean array for the sieve
	isPrime := make([]bool, maxNum+1)
	for i := 2; i <= maxNum; i++ {
		isPrime[i] = true
	}

	// Implement the Sieve of Eratosthenes
	for i := 2; i*i <= maxNum; i++ {
		if isPrime[i] {
			for j := i * i; j <= maxNum; j += i { // Mark multiples as false
				isPrime[j] = false
			}
		}
	}

	// Print prime numbers from the given array
	fmt.Println("Prime numbers in the array:")
	for _, num := range arr {
		if num >= 2 && isPrime[num] {
			fmt.Print(num, " ")
		}
	}
	fmt.Println()
}
