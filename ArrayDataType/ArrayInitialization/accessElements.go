package arrayinitialization

import "fmt"

func AccessElements() {
	arr := [7]int{1, 33, 44, 2, 5, 99, 67}
	fmt.Println("access elements :", arr[4])
	fmt.Println("access elements :", arr[3])
}
