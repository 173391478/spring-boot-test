package utils

import (
	"regexp"
)

// IsValidEmail checks if the provided string is a valid email address
func IsValidEmail(email string) bool {
	// Regular expression for basic email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}
package utils

import (
	"strings"
	"unicode"
)

// CapitalizeWords capitalizes the first letter of each word in the input string.
// It considers words as sequences of characters separated by spaces.
func CapitalizeWords(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			runes := []rune(word)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}