package kata

import (
	"strings"
)

// Deemojify converts a string of emoji representations back into digits
func Deemojify(input string) string {
	emojiToDigit := map[string]string{
		":)":  "0",
		":D":  "1",
		">( ": "2",
		">:C": "3",
		":/":  "4",
		":|":  "5",
		":O":  "6",
		";)":  "7",
		"^.^": "8",
		":(":  "9",
	}

	words := strings.Split(input, "  ")
	result := ""

	for _, word := range words {
		emotes := strings.Split(word, " ")
		number := ""

		for _, emote := range emotes {
			if digit, exists := emojiToDigit[emote]; exists {
				number += digit
			}
		}
		char := string(rune(atoi(number)))
		result += char
	}

	return result
}

// atoi is a helper function that converts a string of digits into an integer
func atoi(s string) int {
	result := 0
	for _, ch := range s {
		result = result*10 + int(ch-'0')
	}
	return result
}