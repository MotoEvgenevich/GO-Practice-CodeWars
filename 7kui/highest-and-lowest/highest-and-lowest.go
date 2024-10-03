package kata

import (
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	var slice []int
	var val int
	var err error
	numbers := strings.Fields(in)
	for _, v := range numbers {
		val, err = strconv.Atoi(string(v))
		if err != nil {
			continue
		}
		slice = append(slice, val)
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	str := strconv.Itoa(slice[len(slice)-1]) + " " + strconv.Itoa(slice[0])

	return str
}

/*
In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

Examples

HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"
Notes

All numbers are valid Int32, no need to validate them.
There will always be at least one number in the input string.
Output string must be two numbers separated by a single space, and highest number is first.
*/
