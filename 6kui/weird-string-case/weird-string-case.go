package kata

import (
	"strings"
)

func ToWeirdCase(str string) string {
	pointer := 0
	slice := make([]string, 0)
	for i, v := range str {
		if v == ' ' {
			pointer = 0
			slice = append(slice, string(v))
			continue
		}
		if pointer%2 == 0 {
			str = str[:i] + strings.ToUpper(string(v)) + str[i+1:]
			pointer++
		} else {
			str = str[:i] + strings.ToLower(string(v)) + str[i+1:]
			pointer++
		}
	}
	return str
}

/*
Write a function that accepts a string, and returns the same string with all even indexed characters in each word upper cased,
and all odd indexed characters in each word lower cased. The indexing just explained is zero based, so the zero-ith index is even,
therefore that character should be upper cased and you need to start over for each word.

The passed in string will only consist of alphabetical characters and spaces(' ').
Spaces will only be present if there are multiple words. Words will be separated by a single space(' ').

Examples:

"String" => "StRiNg"
"Weird string case" => "WeIrD StRiNg CaSe"
*/
