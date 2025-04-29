package advancedarray

func ReturnAnArray() [5]int {
	var arr [5]int
	arr = [5]int{2, 4, 5, 8, 1}
	return arr
}

func ModifyArray(arr *[3]int) {
	(*arr)[0] = 99
}
