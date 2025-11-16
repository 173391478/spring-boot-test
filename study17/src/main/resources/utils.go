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

// CapitalizeWords capitalizes the first letter of each word in a string
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

// IsPalindrome checks if a string reads the same forwards and backwards
func IsPalindrome(s string) bool {
	cleaned := strings.ToLower(strings.Join(strings.Fields(s), ""))
	return cleaned == ReverseString(cleaned)
}
package utils

import (
	"regexp"
	"time"
)

// ValidateEmail checks if the provided string is a valid email address
func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// FormatTimestamp formats a time.Time to RFC3339 string
func FormatTimestamp(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseTimestamp parses an RFC3339 formatted string to time.Time
func ParseTimestamp(timestamp string) (time.Time, error) {
	return time.Parse(time.RFC3339, timestamp)
}