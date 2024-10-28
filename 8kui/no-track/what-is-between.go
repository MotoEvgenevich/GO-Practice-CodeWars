package main

import "fmt"

func main() {
	fmt.Println(Between(-2, 2))
}

func Between(a, b int) []int {
	slice := []int{}
	for i := a; i <= b; i++ {
		slice = append(slice, i)
	}

	return slice
}
