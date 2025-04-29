package arrayinitialization

import "fmt"

func CompilerInferredLength() {
	arr := [...]float64{3.14, 1.61, 2.71}
	fmt.Println("Compiler Inferred Length :", arr)
}
