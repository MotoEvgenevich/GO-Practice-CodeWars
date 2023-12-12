package kata

import (
	"strings"
)

func RemoveDuplicateWords(str string) string {
	wordMap := make(map[string]bool)
	words := strings.Fields(str)
	var result []string
	for _, word := range words {
		if _, exists := wordMap[word]; !exists {
			wordMap[word] = true
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}

/* Your task is to remove all duplicate words from a string, leaving only single (first) words entries.

Example:

Input:

'alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta'

Output:

'alpha beta gamma delta' */
