/* Consider the word "abode". We can see that the letter a is in position 1 and b is in position 2.
In the alphabet, a and b are also in positions 1 and 2.
Notice also that d and e in abode occupy the positions
they would occupy in the alphabet, which are positions 4 and 5.

Given an array of words, return an array of the number of letters that
occupy their positions in the alphabet for each word. For example,

solve(["abode","ABc","xyzD"]) = [4, 3, 1]
See test cases for more examples.

Input will consist of alphabet characters,
both uppercase and lowercase. No spaces.

*/

package kata

import (
	"strings"
)

func Solve(slice []string) []int {
	lowelSlice := []string{}
	count := 0
	result := []int{}
	for _, v := range slice {
		lowelSlice = append(lowelSlice, strings.ToLower(v))
	}

	for _, letter := range lowelSlice {
		for i, char := range letter {
			if int(char-96) == i+1 {
				count += 1
			}
		}
		result = append(result, count)
		count = 0
	}
	return result

}
