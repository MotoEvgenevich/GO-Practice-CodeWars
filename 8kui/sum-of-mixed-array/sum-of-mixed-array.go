package kata

import (
	"strconv"
)

func SumMix(arr []any) int {
	sum := 0
	strArr := []string{}
	for _, v := range arr {
		switch v := v.(type) {
		case int:
			sum += v
		case string:
			strArr = append(strArr, v)
		}
	}
	for _, v := range strArr {
		num, _ := strconv.Atoi(v)
		sum += num
	}
	return sum
}

/*
DESCRIPTION:
Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.
Return your answer as a number.
*/
