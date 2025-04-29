package basicarray

import "fmt"

func CopyArray(arr [5]int) {
	var arr1 [5]int
	arr1 = arr // Copies all elements

	fmt.Println(" copy one array to another :", arr1)
}
