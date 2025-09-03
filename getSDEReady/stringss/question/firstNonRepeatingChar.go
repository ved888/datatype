package question

func FirstNonRepeatingChar(s string) rune {
	seen := make(map[rune]int)

	for _, ch := range s {
		seen[ch]++
	}

	for _, ch := range s {
		if seen[ch] == 1 {
			return ch
		}
	}
	return -1
}

func NonRepeatingChar(s string) []rune {
	seen := make(map[rune]int)

	result := []rune{}

	for _, ch := range s {
		seen[ch]++
	}

	for _, ch := range s {
		if seen[ch] == 1 {
			result = append(result, ch)
		}
	}
	return result
}
