package question

func singleNonDuplicate(nums []int) int {
	seen := make(map[int]int)

	for _, val := range nums {
		seen[val]++
	}

	for _, val := range nums {
		if seen[val] == 1 {
			return val
		}
	}
	return -1
}

// by binary search approch
