package question

func FindPeakElement(arr []int) int {
	ans := arr[0]
	n := len(arr)
	for i := 0; i < n; i++ {
		if ans < arr[i] {
			ans = arr[i]
		}
	}

	for index, val := range arr {
		if ans == val {
			return index
		}
	}
	return -1
}

// by binary search
func FindPeakElementNeighbors(arr []int) int {
	low, high := 0, len(arr)-1

	for low < high {
		mid := low + (high-low)/2

		if arr[mid] > arr[mid+1] {
			high = mid // Peak is on left side including mid
		} else {
			low = mid + 1 // Peak is on right side
		}

	}
	return low

}
