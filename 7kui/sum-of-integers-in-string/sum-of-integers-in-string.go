package kata

import (
	"strconv"
)

func SumOfIntegersInString(strng string) int {
	sum := 0
	digits := ""
	dig := 0
	for _, v := range strng {
		if v != '1' && v != '2' && v != '3' && v != '4' && v != '5' && v != '6' && v != '7' && v != '8' && v != '9' && v != '0' {
			dig, _ = strconv.Atoi(digits)
			sum += dig
			dig = 0
			digits = ""

		} else {
			digits += string(v)
		}

	}
	dig, _ = strconv.Atoi(digits)
	sum += dig
	return sum
}

/*
Your task in this kata is to implement a function that calculates the sum of the integers inside a string.
For example, in the string "The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", the sum of the integers is 3635.

Note: only positive integers will be tested.

*/
