package swapstring

import "strings"

func SwapWordCharatureInPara(str string) string {
	runes := []rune(str)
	n := len(runes)

	for i := 0; i*i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]

	}
	return string(runes)
}

func SwapPara(text string) string {
	split := strings.Split(text, " ")
	var result []string
	for _, v := range split {
		result = append(result, SwapWordCharatureInPara(v))
	}
	return strings.Join(result, " ")
}

// reverse the order of words in a sentence
func ReverseWordOrder(str string) string {
	words := strings.Fields(str)
	n := len(words)

	for i := 0; i < n/2; i++ {
		words[i], words[n-1-i] = words[n-1-i], words[i]
	}

	return strings.Join(words, " ")
}
