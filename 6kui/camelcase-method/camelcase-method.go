package kata

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CamelCase(s string) string {
	words := strings.Split(s, " ")
	key := cases.Title(language.English, cases.NoLower).String(words[0])
	// but codewars don't allow to use package cases use it: key := strings.Title(words[0])
	for _, word := range words[1:] {
		key += cases.Title(language.English, cases.NoLower).String(word)
		// but codewars don't allow to use package cases use it: key += strings.Title(word)
	}
	return key
}

/*
Write a method (or function, depending on the language) that converts a string to camelCase,
that is, all words must have their first letter capitalized and spaces must be removed.

Examples (input --> output):

"hello case" --> "HelloCase"
"camel case word" --> "CamelCaseWord"
Don't forget to rate this kata! Thanks :)

*/
