package kata

import (
	"fmt"
	"strconv"
)

func RangeBitCount(a, b int) int {
	count := 0
	binaryArr := []string{}
	for i := a; i <= b; i++ {
		binary := strconv.FormatInt(int64(i), 2)
		binaryArr = append(binaryArr, binary)
	}
	fmt.Println("binaryArr:", binaryArr)
	for _, v := range binaryArr {
		for _, char := range v {
			if string(char) == "1" {
				count++
			}
		}
	}
	return count
}

/*
Task

You are given two numbers a and b where 0 ≤ a ≤ b. Imagine you construct an array of all the integers from a to b inclusive.
You need to count the number of 1s in the binary representations of all the numbers in the array.

Example

For a = 2 and b = 7, the output should be 11

Given a = 2 and b = 7 the array is: [2, 3, 4, 5, 6, 7]. Converting the numbers to binary,
we get [10, 11, 100, 101, 110, 111], which contains 1 + 2 + 1 + 2 + 2 + 3 = 11 1s.

Input/Output

[input] integer a
Constraints: 0 ≤ a ≤ b.

[input] integer b
Constraints: a ≤ b ≤ 100.

[output] an integer
*/
