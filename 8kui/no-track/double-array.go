package main

import "fmt"

func main() {
	fmt.Println(123)
	Maps([]int{1, 2, 3})
}

func Maps(x []int) []int {
	slice := []int{}
	for _, value := range x {
		slice = append(slice, value*2)
	}
	fmt.Println(slice)
	return slice
}
