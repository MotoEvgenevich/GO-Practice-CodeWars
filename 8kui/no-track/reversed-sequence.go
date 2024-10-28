package main

import "fmt"

func main() {
	fmt.Println(ReverseSeq(7))
}

func ReverseSeq(n int) []int {
	slice := []int{}
	for i := n; i > 0; i-- {
		slice = append(slice, i)
	}

	return slice
}
