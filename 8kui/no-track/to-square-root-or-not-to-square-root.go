package main

import "fmt"

func main() {
	fmt.Println(SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1}))
	fmt.Println(SquareRoot(99))
}

func SquareOrSquareRoot(arr []int) []int {
	slice := []int{}
	var sum int
	for _, v := range arr {
		sum = SquareRoot(v)
		if sum == 0 {
			sum = v * v
		}
		slice = append(slice, sum)

	}
	fmt.Println(slice)
	return slice
}

func SquareRoot(dig int) int {
	for i := 1; i <= (dig / 2); i++ {
		x := i * i
		if x == dig {
			return i
		}
	}
	return 0
}
