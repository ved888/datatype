package advancedarray

func PerfectSqure(n int) bool {
	if n < 0 {
		return false
	}

	low, high := 0, n

	for low <= high {
		mid := (low + high) / 2

		squre := mid * mid

		if squre == n {
			return true
		} else if squre < n {
			low = mid + 1
		} else if squre > n {
			high = mid - 1
		}
	}
	return false
}
