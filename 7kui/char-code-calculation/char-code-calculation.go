package kata

import (
	"fmt"
	"strconv"
)

func Calc(s string) int {
	numStr := ""
	numStr2 := ""
	var total2 int
	for _, v := range s {
		numStr += strconv.Itoa(int(v))
		fmt.Println("V:", v)
	}
	fmt.Println("numStr:", numStr)
	total1, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(total1)
	for _, v := range numStr {
		if v == '7' {
			numStr2 += "1"
			continue
		}
		numStr2 += string(v)
	}
	fmt.Println("numStr2", numStr2)
	total2, _ = strconv.Atoi(numStr2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("total2:", total2)
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
