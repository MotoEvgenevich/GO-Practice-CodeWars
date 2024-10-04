package kata

import (
	"strings"
)

func Solve(s string) int {
	vowel := "aeiyuo"
	maxChain := 0
	chain := 0
	//strings.ContainsRune(vowel, 97)
	for _, v := range s {
		if strings.ContainsRune(vowel, v) {
			chain += 1
			if chain > maxChain {
				maxChain = chain
			}
		} else {
			chain = 0
		}

	}
	return maxChain
}

/* The vowel substrings in the word codewarriors are o,e,a,io.
The longest of these has a length of 2.
Given a lowercase string that has alphabetic characters only
(both vowels and consonants) and no spaces,
return the length of the longest vowel substring.
Vowels are any of aeiou.

*/
