package advancedarray

import "fmt"

func SumOfAllElements(arr [5]int) {
	sum := 0

	for _, v := range arr {
		sum = sum + v
	}
	fmt.Println("sum of all elements :", sum)
}
