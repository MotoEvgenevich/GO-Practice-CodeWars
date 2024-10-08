package kata

import (
	"fmt"
	"strings"
)

func Histogram(results [6]int) string {
	numb := make(map[int]int)

	for i, v := range results {
		numb[i+1] = v

	}
	var histogram string
	for i := 6; i > 0; i-- {
		if numb[i] == 0 {
			histogram += fmt.Sprintf("%d|%s\n", i, strings.Repeat("#", numb[i]))
		} else {
			histogram += fmt.Sprintf("%d|%s %d\n", i, strings.Repeat("#", numb[i]), numb[i])
		}

	}
	return histogram
}

/*
Background

A 6-sided die is rolled a number of times and the results are plotted as a character-based histogram.

Example:

6|##### 5
5|
4|# 1
3|########## 10
2|### 3
1|####### 7
Task

You will be passed the dice value frequencies, and your task is to write the code to return a string representing a histogram, so that when it is printed it has the same format as the example.

Notes

There are no trailing spaces on the lines
All lines (including the last) end with a newline \n
A count is displayed beside each bar except where the count is 0
The number of rolls may vary but there are never more than 100
*/

/* var histogram string
for k, v := range numb {
	histogram += fmt.Sprintf("%d|%s\n", k, strings.Repeat("#", v))
} */
