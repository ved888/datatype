package question

func SubtractProductAndSum(num int) int {

	product, sum := 1, 0

	for num != 0 {
		rem := num % 10
		sum = rem + sum
		product = rem * product
		num = num / 10
	}
	return product - sum

}
