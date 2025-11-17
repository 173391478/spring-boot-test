package utils

func ReverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
package utils

import (
	"strings"
	"unicode"
)

// ReverseString returns the reverse of the input string
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountWords returns the number of words in a string
func CountWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

// IsPalindrome checks if a string is a palindrome
func IsPalindrome(s string) bool {
	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	return cleaned == ReverseString(cleaned)
}