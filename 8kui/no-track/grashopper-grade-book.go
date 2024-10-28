package main

import "fmt"

func main() {
	fmt.Println(GetGrade(98, 95, 89))
}

func GetGrade(a, b, c int) rune {

	grade := (a + b + c) / 3

	switch {
	case 90 <= grade && grade <= 100:
		return 'A'
	case 80 <= grade && grade <= 90:
		return rune(66)
	case 70 <= grade && grade <= 80:
		return rune(67)
	case 60 <= grade && grade <= 70:
		return rune(68)
	case 0 <= grade && grade <= 60:
		return rune(70)

	}
	return 0
}
