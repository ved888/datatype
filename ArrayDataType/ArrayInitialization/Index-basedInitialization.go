package arrayinitialization

import "fmt"

func IndexBasedInitialization() {
	arr := [5]int{0: 10, 3: 40}
	fmt.Println("Index-based Initialization :", arr)
}
