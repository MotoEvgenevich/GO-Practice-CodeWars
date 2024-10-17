package kata

import "sort"

func Solve(s string) bool {
	digits := make(map[rune]int)

	for _, char := range s {
		digits[char]++
		if digits[char] > 1 {
			return false
		}
	}
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

	for i := 1; i < len(runes); i++ {
		if runes[i]-runes[i-1] != 1 {
			return false
		}
	}

	return true
}

/*
In this Kata, we will check if a string contains consecutive letters as they appear in the English alphabet and if each letter occurs only once.

Rules are: (1) the letters are adjacent in the English alphabet, and (2) each letter occurs only once.

For example:
solve("abc") = True, because it contains a,b,c
solve("abd") = False, because a, b, d are not consecutive/adjacent in the alphabet, and c is missing.
solve("dabc") = True, because it contains a, b, c, d
solve("abbc") = False, because b does not occur once.
solve("v") = True
All inputs will be lowercase letters.

More examples in test cases. Good luck!
*/
