package kata

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToJadenCase(str string) string {
	words := strings.Split(str, " ")
	var result string
	for _, word := range words {
		result += cases.Upper(language.English).String(word[0:1]) + word[1:] + " "
	}

	return result[:len(result)-1]
}

/* Jaden Smith, the son of Will Smith, is the star of films such as The Karate Kid (2010) and After Earth (2013).
Jaden is also known for some of his philosophy that he delivers via Twitter.
When writing on Twitter, he is known for almost always capitalizing every word.
For simplicity, you'll have to capitalize each word,
check out how contractions are expected to be in the example below.

Your task is to convert strings to how they would be written by Jaden Smith.
The strings are actual quotes from Jaden Smith, but they are not capitalized in the same way he originally typed them.

Example:

Not Jaden-Cased: "How can mirrors be real if our eyes aren't real"
Jaden-Cased:     "How Can Mirrors Be Real If Our Eyes Aren't Real" */
