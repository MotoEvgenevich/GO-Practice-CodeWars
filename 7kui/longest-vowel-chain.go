/* The vowel substrings in the word codewarriors are o,e,a,io.
The longest of these has a length of 2.
Given a lowercase string that has alphabetic characters only
(both vowels and consonants) and no spaces,
return the length of the longest vowel substring.
Vowels are any of aeiou.

TESTS:
		 Expect(Solve("codewarriors")).To(Equal(2))
        Expect(Solve("suoidea")).To(Equal(3))
        Expect(Solve("ultrarevolutionariees")).To(Equal(3))
        Expect(Solve("strengthlessnesses")).To(Equal(1))
        Expect(Solve("cuboideonavicuare")).To(Equal(2))
        Expect(Solve("chrononhotonthuooaos")).To(Equal(5))
        Expect(Solve("iiihoovaeaaaoougjyaw")).To(Equal(8))
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solve("codewarriors"))
	fmt.Println(Solve("iiihoovaeaaaoougjyaw"))
}

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
