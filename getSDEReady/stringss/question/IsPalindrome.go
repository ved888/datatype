package question

import "unicode"

func isAlphanumeric(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32 // Convert uppercase to lowercase
	}
	return ch
}

// or
func ToLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func ToApper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Move left pointer until it points to alphanumeric or crosses right
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		// Move right pointer until it points to alphanumeric or crosses left
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func WithUnicodeIsPalindrome(s string) bool {
	runes := []rune{}

	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			runes = append(runes, unicode.ToLower(ch))
		}
	}

	n := len(runes)

	for i := 0; i <= n/2; i++ {
		if runes[i] != runes[n-i-1] {
			return false
		}
	}
	return true
}
