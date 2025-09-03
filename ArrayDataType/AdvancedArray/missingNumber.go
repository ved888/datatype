package advancedarray

func MissingNumber(arr []int) int {
	n := len(arr) + 1
	totalSum := (n * (n + 1)) / 2
	achualSum := 0

	for _, num := range arr {
		achualSum += num
	}
	return totalSum - achualSum

}

func MissingAnyWayAnyNumber(arr []int) int {
	if len(arr) == 0 {
		return -1 // no elements
	}

	// Step 1: Find min and max in the array
	min, max := arr[0], arr[0]
	for _, num := range arr {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// Step 2: Calculate expected sum from min to max
	expectedSum := (max*(max+1))/2 - ((min - 1) * min / 2)

	// Step 3: Calculate actual sum
	actualSum := 0
	for _, num := range arr {
		actualSum += num
	}

	// Step 4: Return the missing number
	return expectedSum - actualSum
}

func Shrting1(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// much better shrting
func Shrting(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false // Flag to check if swapping happened

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap elements
				swapped = true
			}
		}

		// If no swapping happened, the array is already sorted
		if !swapped {
			break
		}
	}
}

func FindDuplicates(arr []int) []int {
	fre := make(map[int]int)
	dublicate := []int{}

	for _, num := range arr {
		fre[num]++
		if fre[num] == 2 {
			dublicate = append(dublicate, num)
		}
	}
	return dublicate
}
