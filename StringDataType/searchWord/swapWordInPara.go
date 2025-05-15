package searchword

import "strings"

func SearchWordFromText(text, target string) bool {
	words := strings.Fields(text)
	n := len(words)

	for i := 0; i < n; i++ {
		if words[i] == target {
			return true
		}
	}
	return false
}
