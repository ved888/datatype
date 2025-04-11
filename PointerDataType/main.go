package main

import "fmt"

func main() {
	DeclareAPointer()
	ModifyTheValueOfAVariableThroughAPointer()
	x := 10
	PassAPointerToAFunction(&x)
	ptr := ReturnAPointerFromAFunction()                    // Store the returned pointer
	fmt.Println("return a pointer from a function :", *ptr) // Dereferencing pointer to get the value
	APointerToAPointer()
}

func DeclareAPointer() {
	var name = "ved praksh"

	var po *string // Declare a pointer variable
	po = &name     // Assign the address of 'name' to 'po'

	fmt.Println("address of name variable :", po) // Prints the address
	fmt.Println("Value stored at pointer:", *po)  // Dereferencing the pointer to get value

}

func ModifyTheValueOfAVariableThroughAPointer() {
	x := 20 // Step 1: Declare a variable

	var ptr *int

	ptr = &x // Step 2: Create a pointer to the variable

	fmt.Println("before update values :", x)

	*ptr = 50 // Step 3: Modify the value through the pointer
	fmt.Println("after update values :", x)

}

func PassAPointerToAFunction(ptr *int) {
	fmt.Println("before modified values :", *ptr)

	*ptr = 55
	fmt.Println("after modified values :", *ptr)
}

func ReturnAPointerFromAFunction() *int {
	ptr := 20   // Local variable
	return &ptr // Returning the address of num
}

func APointerToAPointer() {
	x := 5
	var ptr *int
	var ptr1 **int
	ptr = &x
	ptr1 = &ptr
	fmt.Println("a pointer to a pointer :", ptr1)
}
