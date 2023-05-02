/*
In this Kata, you will be given a string that may have mixed uppercase and
lowercase letters and your task is to convert that
string to either lowercase only or uppercase only based on:

make as few changes as possible.
if the string contains equal number of uppercase and lowercase letters, convert the string to lowercase.
For example:

solve("coDe") = "code". Lowercase characters > uppercase. Change only the "D" to lowercase.
solve("CODe") = "CODE". Uppercase characters > lowecase. Change only the "e" to uppercase.
solve("coDE") = "code". Upper == lowercase. Change all to lowercase.
More examples in test cases. Good luck!
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(solve("COde"))
	fmt.Println(solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}

func solve(str string) string {
	upCaseCount := 0
	lowCaseCount := 0
	fixString := ""
	for _, v := range str {
		if v > 64 && v < 91 {
			upCaseCount += 1
		}
		if v > 96 && v < 123 {
			lowCaseCount += 1
		}
	}
	switch {
	case upCaseCount > lowCaseCount:
		for _, v := range str {
			if v > 96 && v < 123 {
				fixString += string(v - 32)
			} else {
				fixString += string(v)
			}
		}
	case upCaseCount <= lowCaseCount:
		for _, v := range str {
			if v > 64 && v < 91 {
				fixString += string(v + 32)
			} else {
				fixString += string(v)
			}
		}
	}

	return fixString
}

func bestPractice(str string) string {
	uppers := 0

	for _, rune := range str {
		if unicode.IsUpper(rune) {
			uppers += 1
		}
	}

	if uppers > len(str)/2 {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
