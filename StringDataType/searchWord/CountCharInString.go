package searchword

func CountCharFromString(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' {
			count++
		}
	}
	return count
}
