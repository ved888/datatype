package searchword

import "fmt"

func ReplaceChar(str string) {
	output := ""

	for i := 0; i < len(str); i++ {
		if str[i] == 'a' {
			output += "z"
		} else {
			output += string(str[i])
		}
	}
	fmt.Println("ReplaceChar :", output)
}
