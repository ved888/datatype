package basicarray

import "fmt"

func MultiDimensionalArray() {
	// 14. Declare a multi-dimensional array (5 rows, 4 columns)
	var arr [5][4]int

	// 15. Initialize the multi-dimensional array with specific values
	arr = [5][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	// Print the entire multi-dimensional array
	fmt.Println("multi-dimensional array:", arr)

	// 16. Access an element at the 4th row and 3rd column (zero-based indexing)
	fmt.Println("access elements in a multi-dimensional array:", arr[3][2]) // Output: 15

	// 17. Loop through the multi-dimensional array
	fmt.Println("Looping through the multi-dimensional array:")
	for i, row := range arr {
		// Print the full row
		fmt.Printf("Row %d: %v\n", i, row)
		// Loop through each value in the row
		for j, val := range row {
			fmt.Printf("arr[%d][%d] = %d\n", i, j, val)
		}
	}

	// 18. Check if a specific value exists in the array
	valueToFind := 15
	found := false
	for _, row := range arr {
		for _, val := range row {
			if val == valueToFind {
				found = true
				break
			}
		}
	}
	if found {
		fmt.Println("Value found!")
	} else {
		fmt.Println("Value not found.")
	}

	// 19. Compare two multi-dimensional arrays for equality
	// Must be same size and type
	arr2 := [5][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	if arr == arr2 {
		fmt.Println("Both arrays are equal.")
	} else {
		fmt.Println("Arrays are not equal.")
	}
}
