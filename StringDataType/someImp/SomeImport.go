package someimp

import (
	"fmt"
	"strings"
	"unicode"
)

//	input := "@@hel12lo!!world"
//
// Output: helloworld
func CleanString1(s string) {
	result := []rune{}

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			result = append(result, ch)
		}
	}
	fmt.Println(string(result))
}

//	input := "@@hel12lo!!world"
//
// Output: hello world
func CleanStringBetter(s string) string {
	var words []string
	var currentWord []rune

	for _, ch := range s {
		if unicode.IsLetter(ch) {
			currentWord = append(currentWord, ch)
		} else if unicode.IsDigit(ch) {
			// Skip digits completely — do not break the word
			continue
		} else {
			// On symbol: treat it as a word boundary
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = nil
			}
		}
	}

	// Add the last word if it exists
	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	return strings.Join(words, " ")
}

// /"@@hel&&lo!!wor55ld%%"
// / hello world
func clearStringGood(str string) string {
	// strings.Builder is efficient for building strings character by character
	var resultBuilder strings.Builder

	// Iterate over each character (rune) in the input string
	for _, r := range str {
		if unicode.IsLetter(r) {
			// If it's a letter, append its lowercase version
			resultBuilder.WriteRune(unicode.ToLower(r))
		} else if unicode.IsDigit(r) {
			// If it's a digit, skip it (remove it)
			continue
		} else if r == '!' || r == ' ' { // Treat '!' and existing spaces as potential word separators
			// If the character is '!' or a space, append a space
			// (we'll normalize spaces later)
			resultBuilder.WriteRune(' ')
		} else {
			// For any other symbols (like '@', '&', '%', '.', '_', '$', etc.),
			// we simply ignore them, effectively removing them.
			continue
		}
	}

	// Convert the builder's content to a string
	cleaned := resultBuilder.String()

	// strings.Fields splits the string by one or more consecutive whitespace characters
	// and returns a slice of non-empty strings. It naturally handles:
	// 1. Removing leading/trailing spaces.
	// 2. Replacing multiple internal spaces with a single space.
	words := strings.Fields(cleaned)

	// Join the words back together with a single space
	return strings.Join(words, " ")
}

// --------------------OR---------------
// /"@@hel&&lo!!wor55ld%%"
// / hello world
func clearStringWithoutBuilder(str string) string {
	// Initialize a rune slice to store the filtered characters.
	// We can pre-allocate some capacity to minimize reallocations.
	filteredRunes := make([]rune, 0, len(str))

	for _, r := range str {
		if unicode.IsLetter(r) {
			// If it's a letter, append its lowercase version.
			filteredRunes = append(filteredRunes, unicode.ToLower(r))
		} else if unicode.IsDigit(r) {
			// If it's a digit, skip it (remove it).
			continue
		} else if r == '!' || r == ' ' {
			// If the character is '!' or an existing space, append a space.
			filteredRunes = append(filteredRunes, ' ')
		} else {
			// For any other symbols, skip them (remove them).
			continue
		}
	}

	// Convert the filtered rune slice to a string.
	cleaned := string(filteredRunes)

	// strings.Fields splits the string by any sequence of whitespace characters
	// and returns a slice of non-empty strings. This handles:
	// 1. Removing leading/trailing spaces.
	// 2. Reducing multiple internal spaces (e.g., from `!!` becoming two spaces) to a single space.
	words := strings.Fields(cleaned)

	// Join the words back together with a single space.
	return strings.Join(words, " ")
}

// //////////////or/////////////////
//
//	input := "@@hel12lo!!world"
//
// Output: hello world
func clearString(str string) string {
	result := []rune{}
	spaceAdd := false

	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
			result = append(result, ch)
			spaceAdd = false
		} else if ch >= '0' && ch <= '9' {
			continue
		} else if !spaceAdd && len(result) > 0 && result[len(result)-1] != ' ' {
			result = append(result, ' ')
			spaceAdd = true
		}
	}
	return string(result)
}

//	input := "@hello!!world"
//
// Output: hello world
func CleanString2(s string) {
	var result []string
	word := ""

	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			word += string(ch)
		} else {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		}
	}

	// Append the last word if any
	if word != "" {
		result = append(result, word)
	}

	// Join all words with a space
	fmt.Println(strings.Join(result, " "))
}

//	input := "@@hel12lo!!world"
//
// Output: hello world
func CleanString(s string) {
	var result strings.Builder
	inWord := false // track if we're inside a word

	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			// letter: add it
			result.WriteRune(ch)
			inWord = true
		} else if inWord {
			// non-letter after a word ends → add space once
			result.WriteRune(' ')
			inWord = false
		}
	}

	// Trim trailing space if any
	output := strings.TrimSpace(result.String())
	fmt.Println(output)
}

func ToLower(s string) {
	result := []rune{}

	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			result = append(result, ch+32)
		} else {
			result = append(result, ch)
		}
	}
	fmt.Println(string(result))
}

func ToUpper(s string) {
	result := []rune{}

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result = append(result, ch-32)
		} else {
			result = append(result, ch)
		}
	}
	fmt.Println(string(result))
}
