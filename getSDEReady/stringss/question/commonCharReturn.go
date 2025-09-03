package question

// Input:
// words = ["bella","label","roller"]
// Output:
// ["e","l","l"]
func CommonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	// Initialize minFreq with frequency map of the first word
	minFreq := make([]int, 26)
	for i := range minFreq {
		minFreq[i] = int(^uint(0) >> 1) // Max int
	}

	for _, word := range words {
		freq := make([]int, 26)
		for _, char := range word {
			freq[char-'a']++
		}
		for i := 0; i < 26; i++ {
			if freq[i] < minFreq[i] {
				minFreq[i] = freq[i]
			}
		}
	}

	result := []string{}
	for i := 0; i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			result = append(result, string('a'+i))
		}
	}
	return result
}
