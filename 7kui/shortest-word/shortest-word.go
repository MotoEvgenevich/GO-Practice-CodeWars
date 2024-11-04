package kata

import "strings"

func FindShort(s string) int {
	words := strings.Fields(s)
	shortest := len(words[0])
	for _, word := range words {
		if len(word) < shortest {
			shortest = len(word)
		}
	}
	return shortest
}
