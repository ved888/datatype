package question

// Given a m x n matrix grid which is sorted in non-increasing order both row-wise and column-wise, return the number of negative numbers in grid.

// by binary search
func FindNegativeFromSortedMatrix(matrix [][]int) int {
	count := 0
	for _, row := range matrix {
		low, high := 0, len(row)

		for low < high {
			mid := low + (high-low)/2

			if row[mid] < 0 {
				high = mid
			} else {
				low = mid + 1
			}

		}
		count += len(row) - low
	}
	return count

}
