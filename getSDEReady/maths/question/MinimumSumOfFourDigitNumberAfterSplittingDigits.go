package question

import "sort"

func MinimumSumOf4Digit(num int) int {
	res := make([]int, 4)
	i := 0

	for num != 0 {
		res[i] = num % 10
		num = num / 10
		i++
	}

	sort.Ints(res)
	return res[0]*10 + res[2] + res[1]*10 + res[3]

}

func MinimumSumOfDigit(n int) int {
	res := []int{}

	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}
	sort.Ints(res)
	return ((res[0]*10 + res[2]) + (res[1]*10 + res[3]))
}
