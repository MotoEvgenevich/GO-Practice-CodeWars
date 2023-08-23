package kata

import (
	"strings"
)

func Capitalize(st string, arr []int) string {
	result := ""
	lowerCaseLetters := make(map[int]string)
	for key, letter := range st {
		lowerCaseLetters[key] = string(letter)
	}
	for _, v := range arr {
		lowerCaseLetters[v] = strings.ToUpper(lowerCaseLetters[v])
	}
	for i := 0; i < len(st); i++ {
		result += lowerCaseLetters[i]
	}
	return result
}

/*
Given a string and an array of integers representing indices, capitalize all letters at the given indices.

For example:

capitalize("abcdef",[1,2,5]) = "aBCdeF"
capitalize("abcdef",[1,2,5,100]) = "aBCdeF". There is no index 100.
The input will be a lowercase string with no spaces and an array of digits.

Good luck! */
