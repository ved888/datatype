package questionvise

func IsFactorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * IsFactorial(num-1)
}
