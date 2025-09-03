package advancedarray

func Permute(nums []int) [][]int {
	var result [][]int

	var backtrack func([]int, int)
	backtrack = func(arr []int, start int) {
		if start == len(arr) {
			perm := make([]int, len(arr))
			copy(perm, arr)
			result = append(result, perm)
			return
		}

		for i := start; i < len(arr); i++ {
			arr[start], arr[i] = arr[i], arr[start] // swap
			backtrack(arr, start+1)
			arr[start], arr[i] = arr[i], arr[start] // backtrack
		}
	}

	backtrack(nums, 0)
	return result
}

// func Permute(num []int) {
// 	var helper func([]int, int)
// 	helper = func(arr []int, n int) {
// 		if n == 1 {
// 			fmt.Println(arr)
// 			return
// 		}

// 		for i := 0; i < n; i++ {
// 			helper(arr, n-1)
// 			if n%2 == 1 {
// 				arr[0], arr[n-1] = arr[n-1], arr[0]
// 			} else {
// 				arr[i], arr[n-1] = arr[n-1], arr[i]
// 			}

// 		}
// 	}
// 	helper(num, len(num))
// }
//
//
//
//
//
//
//for understand

// package main

// import (
// 	"fmt"
// )

// func Permute(nums []int) [][]int {
// 	var result [][]int

// 	var backtrack func([]int, int)
// 	backtrack = func(arr []int, start int) {
// 		fmt.Printf("Start = %d, Current arr = %v\n", start, arr)

// 		if start == len(arr) {
// 			perm := make([]int, len(arr))
// 			copy(perm, arr)
// 			result = append(result, perm)
// 			fmt.Printf("✅ Added permutation to result: %v\n", perm)
// 			return
// 		}

// 		for i := start; i < len(arr); i++ {
// 			// Swap to try a new number at index `start`
// 			arr[start], arr[i] = arr[i], arr[start]
// 			fmt.Printf("🔁 Swapped arr[%d] with arr[%d] => %v\n", start, i, arr)

// 			// Recurse to fill the next position
// 			backtrack(arr, start+1)

// 			// Backtrack (swap back)
// 			arr[start], arr[i] = arr[i], arr[start]
// 			fmt.Printf("↩️ Backtracked arr[%d] with arr[%d] => %v\n", start, i, arr)
// 		}
// 	}

// 	backtrack(nums, 0)
// 	return result
// }

// func main() {
// 	nums := []int{1, 2, 3}
// 	result := Permute(nums)
// 	fmt.Println("\n✅ All permutations:")
// 	for i, perm := range result {
// 		fmt.Printf("%d: %v\n", i+1, perm)
// 	}
// }
