package main

import "fmt"

func main() {
	fmt.Println(MinMax([]int{2334454, 5}))
}

func MinMax(arr []int) [2]int {
	var min int = arr[0]
	var max int = arr[0]
	fmt.Println(min)
	fmt.Println(max)
	for _, nmb := range arr {
		if nmb < min {
			min = nmb
		}
	}
	for _, nmb := range arr {
		if nmb > max {
			max = nmb
		}
	}
	result := [2]int{min, max}

	return result
}
