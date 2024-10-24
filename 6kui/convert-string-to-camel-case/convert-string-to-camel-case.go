package kata

import (
	"strings"
)

func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(c rune) bool {
		return c == '-' || c == '_'
	})

	if len(words) == 0 {
		return ""
	}

	result := words[0]

	for i := 1; i < len(words); i++ {
		result += strings.Title(words[i])
		//result += cases.Title(language.English).String(words[i])
	}

	return result
}

/*
Complete the method/function so that it converts dash/underscore delimited words into camel casing.
The first word within the output should be capitalized only if the original word was capitalized
(known as Upper Camel Case, also often referred to as Pascal case). The next words should be always capitalized.

Examples

"the-stealth-warrior" gets converted to "theStealthWarrior"

"The_Stealth_Warrior" gets converted to "TheStealthWarrior"

"The_Stealth-Warrior" gets converted to "TheStealthWarrior"

*/
