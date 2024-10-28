package main

import "fmt"

func main() {
	fmt.Println(Grow([]int{1, 2, 3, 4}))
}

func Grow(arr []int) int {
	sum := 1
	for _, v := range arr {
		sum *= v
	}
	return sum
}
