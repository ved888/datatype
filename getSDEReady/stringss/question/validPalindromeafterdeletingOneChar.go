package question

func ValidPalindromeAfterDeletingOneChar(s string) bool {

	isPalindrom := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return isPalindrom(left+1, right) || isPalindrom(left, right-1)
		}
		left++
		right--
	}
	return true
}
