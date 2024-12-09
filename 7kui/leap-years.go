/* In this kata you should simply determine, whether a given year is a leap year or not. In case you don't know the rules, here they are:

years divisible by 4 are leap years
but years divisible by 100 are not leap years
but years divisible by 400 are leap years
Additional Notes:

Only valid years (positive integers) will be tested, so you don't have to validate them
Examples can be found in the test fixture. */

package main

import "fmt"

func main() {
	fmt.Println(IsLeapYear(2020)) //true
	fmt.Println(IsLeapYear(2000)) //true
	fmt.Println(IsLeapYear(2015)) //false
	fmt.Println(IsLeapYear(2100)) //false
}

func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	} else {
		return false
	}
}
