package advancedarray

func ShiftLeft(arr *[5]int) {

	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = 0

}

func ShiftRight(arr *[5]int) {
	for i := len(arr) - 1; i >= 1; i-- {
		arr[i] = arr[i-1]
	}

	arr[0] = 0
}
