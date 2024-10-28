package main

import "fmt"

func main() {
	numbers := []int{1, 2, 2}
	fmt.Println(SquareSum(numbers))
}

func SquareSum(numbers []int) int {
	totalsum := 0
	for i := 0; i < len(numbers); i++ {
		sum := numbers[i] * numbers[i]
		totalsum += sum
	}

	return totalsum
}
