package main

import "fmt"

func main() {
	fmt.Println(Factorial(5))
}

func Factorial(n int) int {
	result := 1
	for i := n; i > 0; i-- {
		result *= i
	}
	return result
}
