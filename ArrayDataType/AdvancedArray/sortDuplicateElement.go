package advancedarray

import "sort"

func SortBasicOnDuplicateElement(arr []int) {

	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}

// convert array to slice
func ArrayToSlice(arr [5]int) []int {
	// Convert array to slice using slicing syntax
	slice := arr[:]
	// this function works on only slice
	sort.Ints(slice)
	return slice
}

func SortStringArray(arr []string) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}
