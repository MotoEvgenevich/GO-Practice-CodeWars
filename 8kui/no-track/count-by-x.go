package main

import "fmt"

func main() {
	fmt.Println(CountBy(2, 5))
}

func CountBy(x, n int) []int {
	slice := []int{}
	var sum int
	for i := 1; i <= n; i++ {
		sum += x
		slice = append(slice, sum)
	}

	return slice
}
