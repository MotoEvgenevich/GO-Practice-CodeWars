package kata

import (
	"strconv"
)

func Solve(st string, k int) int {
	n := len(st)
	maxValue := 0
	var generateSplits func(pos int, splits []int)
	generateSplits = func(pos int, splits []int) {
		if len(splits) == k {
			prev := 0
			for _, split := range splits {
				part, _ := strconv.Atoi(st[prev:split])
				if part > maxValue {
					maxValue = part
				}
				prev = split
			}
			lastPart, _ := strconv.Atoi(st[prev:])
			if lastPart > maxValue {
				maxValue = lastPart
			}
			return
		}
		for i := pos + 1; i < n; i++ {
			generateSplits(i, append(splits, i))
		}
	}
	generateSplits(0, []int{})
	return maxValue
}

/*
In this Kata, you will be given a number in form of a string and an integer k and your task is to insert k commas into the string and determine which of the partitions is the largest.

For example:
solve('1234',1) = 234 because ('1','234') or ('12','34') or ('123','4').
solve('1234',2) = 34 because ('1','2','34') or ('1','23','4') or ('12','3','4').
solve('1234',3) = 4
solve('2020',1) = 202
More examples in test cases. Good luck!

Please also try Simple remove duplicates

*/
