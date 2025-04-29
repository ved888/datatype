package arrayinitialization

import "fmt"

// Partial Initialization (rest get zero values)

func PartialInitialization() {
	arr := [5]int{1, 2}
	fmt.Println("Partial Initialization :", arr)
}
