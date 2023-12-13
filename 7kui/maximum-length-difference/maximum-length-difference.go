package kata

func MxDifLg(a1 []string, a2 []string) int {
	maxLenA1 := 0
	maxLenA2 := 0
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	for _, v := range a1 {
		if len(v) > maxLenA1 {
			maxLenA1 = len(v)
		}
	}
	for _, v := range a2 {
		if len(v) > maxLenA2 {
			maxLenA2 = len(v)
		}
	}
	minLenA1 := maxLenA1
	minLenA2 := maxLenA2
	for _, v := range a1 {
		if len(v) < minLenA1 {
			minLenA1 = len(v)
		}
	}
	for _, v := range a2 {
		if len(v) < minLenA2 {
			minLenA2 = len(v)
		}
	}
	compSum1 := maxLenA1 - minLenA2
	compSum2 := maxLenA2 - minLenA1

	if compSum1 >= compSum2 {
		return compSum1
	}
	if compSum2 >= compSum1 {
		return compSum2
	}

	return 0
}

/*
DESCRIPTION:

You are given two arrays a1 and a2 of strings. Each string is composed with letters from a to z.
Let x be any string in the first array and y be any string in the second array.

Find max(abs(length(x) − length(y)))

If a1 and/or a2 are empty return -1 in each language except in Haskell (F#) where you will return Nothing (None).

Example:

a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
mxdiflg(a1, a2) --> 13
Bash note:

input : 2 strings with substrings separated by ,
output: number as a string

*/
