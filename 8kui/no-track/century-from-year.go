package main

import "fmt"

func main() {
	fmt.Println(century(1900))
}

func century(year int) int {
	var cent float64
	var centInt int
	cent = (float64(year) / 100)
	//fmt.Println(cent)
	centInt = int(cent)
	total := float64(centInt) - cent
	fmt.Println(total)
	if total != 0 {
		return int(cent) + 1
	}

	return int(cent)
}
