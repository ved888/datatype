package main

import (
	searchword "StringDataType/searchWord"
	swapstring "StringDataType/swapString"
	"fmt"
)

func main() {
	str := swapstring.SwapPara("Hello world, this is Go")
	fmt.Println("word revers in para is :", str)
	word := swapstring.ReverseWordOrder("Hello world, this is Go")
	fmt.Println("Reverse the words :", word)
	match := searchword.SearchWordFromText("Hello world, this is Go", "Hello")
	fmt.Println("word is match :", match)
	fmt.Println("FindSubString :", searchword.FindSubString("GolangIsGood", "Is"))
	fmt.Println("CountCharFromString :", searchword.CountCharFromString("GoLangIsGreat"))
	searchword.ReplaceChar("abacada")
}
