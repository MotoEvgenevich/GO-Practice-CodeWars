package kata

import (
	"strconv"
)

func Calc(s string) int {
	numStr := ""
	total1 := 0
	total2 := 0
	for _, v := range s {
		numStr += strconv.Itoa(int(v))
	}
	for _, v := range numStr {
		num, _ := strconv.Atoi(string(v))
		total1 += num
		if num == 7 {
			total2 += 1
			continue
		}
		total2 += num
	}
	return total1 - total2
}

/*
Given a string, turn each character into its ASCII character code and join them together to create a number - let's call this number total1:

'ABC' --> 'A' = 65, 'B' = 66, 'C' = 67 --> 656667
Then replace any incidence of the number 7 with the number 1, and call this number 'total2':

total1 = 656667
              ^
total2 = 656661
              ^
Then return the difference between the sum of the digits in total1 and total2:

  (6 + 5 + 6 + 6 + 6 + 7)
- (6 + 5 + 6 + 6 + 6 + 1)
-------------------------
                       6
*/
