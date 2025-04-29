package advancedarray

func SortAnArrayInAscending(arr *[5]int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

func SortAnArrayInDescending(arr *[5]int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func RemoveDuplicatesFromArray(arr [5]int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func RemoveDuplicatesFromArrayWithoutReturn(arr *[5]int) {
	seen := make(map[int]bool)
	index := 0

	for i := 0; i < len(arr); i++ {
		if !seen[arr[i]] {
			seen[arr[i]] = true
			arr[index] = arr[i]
			index++
		}
	}
	// Fill the rest with zeros
	for i := index; i < len(arr); i++ {
		arr[i] = 0
	}
}
