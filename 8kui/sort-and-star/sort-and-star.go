package kata

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	totalWord := strings.Join(strings.Split(arr[0], ""), "***")
	return totalWord
}

/* DESCRIPTION:

You will be given a list of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars) and then return the first value.
The returned value must be a string, and have "***" between each of its letters.
You should not remove or add elements from/to the array.

I don't uderstand that bugging when I made attempt, and write function wuth st-lib. Maybe later i will understand.
Expected
    <string>: "H***o***x*..."
to equal           |
    <string>: "H***c***u*..."
*/

/* func TwoSort(arr []string) string {
	firstWord := arr[0]
	twoWordCompare := arr[0]
	firstLetter := firstWord[0]
	totalWord := ""
	for _, v := range arr {
		if v[0] < firstLetter {
			firstLetter = v[0]
			firstWord = v
		}
		if v[0] == firstLetter {
			twoWordCompare = v
		}

	}
	for i, v := range twoWordCompare {
		if byte(v) < firstWord[i] {
			firstWord = twoWordCompare
			break
		}

	}

	for _, v := range firstWord {
		totalWord += string(v) + "***"
	}
	totalWord = totalWord[0 : len(totalWord)-3]
	return totalWord
}
*/
