package kata

import (
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	switch {
	case len(str) == 0:
		return true
	case len(str) == 1:
		return true
	case len(str)%2 == 0:
		i := (len(str) / 2)
		j := (len(str) / 2) - 1

		for ; j >= 0; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
	case len(str)%2 != 0:
		i := (len(str) / 2) + 1
		j := (len(str) / 2) - 1
		for ; j >= 0; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
	}
	return true

}

/*
DESCRIPTION:

Write a function that checks if a given string (case insensitive) is a palindrome.
A palindrome is a word, number, phrase,
or other sequence of symbols that reads the same backwards as forwards,
such as madam or racecar, the date and time 12/21/33 12:21,
and the sentence: "A man, a plan, a canal – Panama". */

/*
TEST CASES:

Expect(IsPalindrome("a")).To(Equal(true))
Expect(IsPalindrome("aba")).To(Equal(true))
Expect(IsPalindrome("Abba")).To(Equal(true))
Expect(IsPalindrome("hello")).To(Equal(false))
*/
