package question

import "strings"

func RepeatedSubstringPattern(str string) bool {
	n := len(str)

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			pattern := str[:i]
			repeated := ""
			for j := 0; j < n/i; j++ {
				repeated += pattern
			}
			if repeated == str {
				return true
			}
		}

	}

	return false

}

func ReverseString(str string) string {
	rune := []rune(str)
	n := len(rune)

	for i := 0; i < n/2; i++ {
		rune[i], rune[n-i-1] = rune[n-i-1], rune[i]
	}
	return string(rune)
}

func Para(text string) string {
	split := strings.Split(text, " ")
	result := []string{}
	for _, v := range split {
		result = append(result, ReverseString(v))
	}
	return string.Join(result, " ")

}

// ----------------- OR ------------------

// func repeatedSubstringPattern(s string) bool {
// 	n := len(s)
// 	for i := 1; i <= n/2; i++ {
// 		if n%i == 0 {
// 			pattern := s[:i]
// 			repeated := strings.Repeat(pattern, n/i)
// 			if repeated == s {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
