package main

import "fmt"

func main() {
	fmt.Println(Solve("dcbadcba"))
}
func Solve(s string) rune {
	m := make(map[int]int)
	for i, v := range s {
		m[i] = int(v)
	}
	fmt.Println(len(m))
	MaxCount := 0
	//result := 0
	for i := 0; i < (len(m) - 1); i++ {
		if m[i]-i > MaxCount {
			MaxCount = m[i]
		}

	}
	if len(m) == 1 {
		MaxCount = m[0]
	}
	return rune(MaxCount)
}

/* In this Kata, you will be given a string and your task is to return the most valuable character.
The value of a character is the difference between the index of its last occurrence and the index of its first occurrence.
Return the character that has the highest value.
If there is a tie, return the alphabetically lowest character. [For Golang return rune]

All inputs will be lower case.

For example:
solve('a') = 'a'
solve('ab') = 'a'. Last occurrence is equal to first occurrence of each character. Return lexicographically lowest.
solve("axyzxyz") = 'x'
"axyzxyz", 120
More examples in test cases. Good luck! */
