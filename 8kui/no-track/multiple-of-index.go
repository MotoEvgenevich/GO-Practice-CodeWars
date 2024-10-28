package main

import "fmt"

func main() {
	fmt.Println(multipleOfIndex([]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68}))
}

func multipleOfIndex(ints []int) []int {
	// good luck
	sliceInts := []int{}
	for i, value := range ints {
		if i > 0 && value%i == 0 {
			sliceInts = append(sliceInts, value)
		}

	}
	return sliceInts
}
