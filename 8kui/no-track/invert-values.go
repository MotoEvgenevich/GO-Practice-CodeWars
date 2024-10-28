package main

import "fmt"

func main() {
	fmt.Println(Invert([]int{123, -232, 0, 432, -7353}))
}

func Invert(arr []int) []int {
	slice := []int{}
	for _, v := range arr {
		slice = append(slice, -v)
	}
	return slice
}
