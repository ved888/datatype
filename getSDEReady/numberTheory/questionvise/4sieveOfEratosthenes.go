// Problem Statement:

// Given a number N, calculate the prime numbers up to N using Sieve of Eratosthenes.

package questionvise

func SieveOfEratosthenes(n int) []int {
	isPrime := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	prime := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prime = append(prime, i)
		}
	}
	return prime
}
