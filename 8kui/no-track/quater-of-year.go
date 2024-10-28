package main

import "fmt"

func main() {
	fmt.Println(QuarterOf1(7))
}

func QuarterOf(month int) int {

	if 1 <= month && month <= 3 {
		return 1
	}
	if 3 <= month && month <= 6 {
		return 2
	}
	if 6 <= month && month <= 9 {
		return 3
	} else {
		return 4
	}
}

func QuarterOf1(month int) int {
	var res int
	switch {
	case month <= 3:
		res = 1
	case month <= 6:
		res = 2
	case month <= 9:
		res = 3
	default:
		res = 4
		return res
	}

}
