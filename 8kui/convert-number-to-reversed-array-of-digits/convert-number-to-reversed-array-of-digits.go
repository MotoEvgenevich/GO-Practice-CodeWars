package kata

import (
	"strconv"
)

func Digitize(n int) []int {
	numStr := strconv.Itoa(n)
	strSlice := []string{}
	numSlice := []int{}
	result := []int{}

	for _, v := range numStr {
		strSlice = append(strSlice, string(v))
	}

	for _, num := range strSlice {
		if s, err := strconv.Atoi(num); err == nil {
			numSlice = append(numSlice, s)
		}
	}

	for i := len(numSlice) - 1; i >= 0; i-- {
		result = append(result, numSlice[i])
	}
	return result
}

func DigitizeClever(n int) []int {
	var ret []int
	for {
		ret = append(ret, n%10)
		n /= 10
		if n == 0 {
			return ret
		}
	}
}

/* Convert number to reversed array of digits

Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

Example(Input => Output):

35231 => [1,3,2,5,3]
0 => [0] */
