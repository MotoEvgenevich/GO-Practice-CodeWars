package kata

import (
	"regexp"
)

func Disemvowel(comment string) string {
	fix := ""
	for _, v := range comment {
		switch v {
		case 65:
		case 97:
		case 69:
		case 73:
		case 105:
		case 101:
		case 79:
		case 111:
		case 85:
		case 117:
		default:
			fix += string(v)
		}
	}

	return fix
}

func BestPractice(comment string) string {
	return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(comment, "")
}

func DisemvowelV2(comment string) string {
	fix := ""
	for _, v := range comment {
		switch v {
		case 65, 97, 69, 73, 105, 101, 79, 111, 85, 117:
		default:
			fix += string(v)
		}
	}

	return fix
}

func DisemvowelV3(comment string) string {
	fix := ""
	for _, v := range comment {
		switch string(v) {
		case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		default:
			fix += string(v)
		}
	}

	return fix
}

/* Trolls are attacking your comment section!

A common way to deal with this situation is to remove all of the vowels
from the trolls' comments, neutralizing the threat.

Your task is to write a function that takes a string and return a new string
with all vowels removed.

For example, the string "This website is for losers LOL!"
would become "Ths wbst s fr lsrs LL!".

Note: for this kata y isn't considered a vowel. */
