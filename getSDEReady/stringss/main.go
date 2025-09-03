package main

import (
	"fmt"
	"stringss/question"
)

func main() {
	fmt.Println("", question.RepeatedSubstringPattern("abcabcabcabc"))
	fmt.Println("ValidPalindrome :", question.ValidPalindrome("abca"))
	fmt.Println("ReverseString :", question.ReverseString("hello"))
	fmt.Println("ValidPalindromeAfterDeletingOneChar :", question.ValidPalindromeAfterDeletingOneChar("abcbda"))
	fmt.Println("IsPalindrome :", question.IsPalindrome("A man, a plan, a canal: Panama"))                       // true
	fmt.Println("IsPalindrome :", question.IsPalindrome("race a car"))                                           // false
	fmt.Println("WithUnicodeIsPalindrome :", question.WithUnicodeIsPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println("WithUnicodeIsPalindrome :", question.WithUnicodeIsPalindrome("race a car"))                     // false
}
