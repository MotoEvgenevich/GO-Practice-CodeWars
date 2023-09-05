package kata

import (
	"strings"
)

func ReverseWords(str string) string {
	slice := append(strings.Split(str, " "))
	m := make(map[int]string)
	result := ""
	for i, word := range slice {
		m[i] = word

	}
	for i := len(slice) - 1; i >= 0; i-- {
		result += m[i] + " "
	}
	result = result[:len(result)-1]

	return result
}

func BestPractice(str string) string {
	words := strings.Split(str, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

/* Complete the solution so that it reverses all of the words within the string passed in.

Words are separated by exactly one space and there are no leading or trailing spaces.

Example(Input --> Output):

"The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The" */
