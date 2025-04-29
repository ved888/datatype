package advancedarray

func MergeArrays(arr1 [5]int, arr2 [5]int) []int {
	marge := []int{}

	for _, v := range arr1 {
		marge = append(marge, v)
	}

	for _, v := range arr2 {
		marge = append(marge, v)
	}
	return marge
}
