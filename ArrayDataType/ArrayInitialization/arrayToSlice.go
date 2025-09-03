package arrayinitialization

import "fmt"

func ArrayToSliceAndRemoveElement() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4] // [20 30 40]

	fmt.Println("Original Slice:", slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	// Remove element at index 1 (30)
	slice = append(slice[:1], slice[2:]...)

	fmt.Println("After Removal:", slice)
	fmt.Println("Length after removal:", len(slice))
	fmt.Println("Capacity after removal:", cap(slice))
}
