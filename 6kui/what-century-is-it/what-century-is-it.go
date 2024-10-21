package kata

import "strconv"

func WhatCentury(year string) string {
	yearInt, _ := strconv.Atoi(year)
	cent := (yearInt + 99) / 100

	switch cent % 10 {
	case 1:
		if cent%100 != 11 {
			return strconv.Itoa(cent) + "st"
		}
	case 2:
		if cent%100 != 12 {
			return strconv.Itoa(cent) + "nd"
		}
	case 3:
		if cent%100 != 13 {
			return strconv.Itoa(cent) + "rd"
		}
	}
	return strconv.Itoa(cent) + "th"
}

/*
Return the century of the input year. The input will always be a 4 digit string, so there is no need for validation.

Examples

"1999" --> "20th"
"2011" --> "21st"
"2154" --> "22nd"
"2259" --> "23rd"
"1124" --> "12th"
"2000" --> "20th"
*/

/*
func century(year int) int {

}
*/
