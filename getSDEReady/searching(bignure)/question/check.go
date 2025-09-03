package question

// Check if Array Is Sorted and Rotated
// Input: nums = [3,4,5,1,2]
// Output: true
// Explanation: [1,2,3,4,5] is the original sorted array.
// You can rotate the array by x = 3 positions to begin on the the element of value 3: [3,4,5,1,2]

func Check(num []int) bool {
	n := len(num)
	count := 0
	for i := 0; i < n; i++ {
		if num[i] > num[(i+1)%n] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
