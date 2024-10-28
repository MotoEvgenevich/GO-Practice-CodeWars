package main

import "fmt"

func main() {
	fmt.Println(FindMultiples(68, 11884))
}

func FindMultiples(integer, limit int) []int {
	var slice []int
	sum := 0
	var slice2 []int

	for i := 0; sum < limit; i++ {
		sum += integer
		slice = append(slice, sum)
	}
	if sum > limit {
		for _, num := range slice {
			if num < limit {
				slice2 = append(slice2, num)
			}
		}
	} else {
		return slice
	}

	return slice2
}
