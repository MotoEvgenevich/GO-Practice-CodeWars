package kata

import "strings"

func FindLongestCommonPrefix(slice []string) string {
	if len(slice) == 0 {
		return ""
	}

	for i := 0; i < len(slice[0]); i++ {
		for j := 1; j < len(slice); j++ {
			if i >= len(slice[j]) || strings.ToLower(string(slice[0][i])) != strings.ToLower(string(slice[j][i])) {
				return strings.ToLower(slice[0][:i])
			}
		}
	}
	return strings.ToLower(slice[0])
}

/*
Write a function to find the longest common prefix string amongst an array of strings.

Examples:
Input: ["FLOWER", "Flow", "fly"]
Output: "fl"

Input: ["dog", "racecar", "car"]
Output: "" (No common prefix)

Input: ["apple", "ape", "app"]
Output: "ap"

Input: ["prefix", "preference", "president"]
Output: "pre"


Inputs:
Array of strings that contain ASCII && Non-ASCII Lower && Upper case Characters. "Codepoints/Unicode chars && Non-Printable characters!"
Array Length: 0-10K
String/Item Length: 0-10K

Output:
Lower-Case string represents the longest common prefix for all strings provided.

Note:
Prefixes are supposed to be considered and returned case-insensitively. Notice provided sample test cases 📢

Cheers,

*/
