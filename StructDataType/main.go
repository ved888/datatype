package main

import (
	"fmt"
)

func main() {
	AccessAndModifyStructFields()
	AnonymousStruct()
	CompareTwoStructs()
}

func AccessAndModifyStructFields() {
	type User struct {
		Name string
		Age  int
	}
	p := User{"ved", 28}
	fmt.Println("before modification user name :", p.Name)
	// modify struct field
	p.Name = "Rahul"
	fmt.Println("after modification user name :", p.Name)
}

func AnonymousStruct() {
	person := struct {
		Name string
		Age  int
	}{
		Name: "ved prakash",
		Age:  28,
	}
	fmt.Println("struct data :", person)
}

func CompareTwoStructs() {
	type User struct {
		Name string
		Age  int
	}

	type Student struct {
		Name  string
		Class int
	}

	user := User{"ved", 18}
	student := Student{"ved", 12}

	if user.Name == student.Name {
		fmt.Println("Names are the same!")
	} else {
		fmt.Println("Names are different!")
	}

}
