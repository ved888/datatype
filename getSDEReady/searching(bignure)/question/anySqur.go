package question

// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.

func AnySqure(n int) int {
	low, high := 0, n
	ans := 0

	for low <= high {
		mid := low + (high-low)/2
		squre := mid * mid

		if squre <= n {
			ans = mid
			low = mid + 1
		} else if squre > n {
			high = mid - 1
		}

	}
	return ans
}
