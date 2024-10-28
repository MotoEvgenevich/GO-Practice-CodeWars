package main

import "fmt"

func main() {
	fmt.Println(Quadratic(2, 6))
}

func Quadratic(x1, x2 int) [3]int {
	a := 1
	b := -(x1 + x2)
	c := x1 * x2
	arr := [3]int{a, b, c}

	return arr
}

// ax^2 + bx + c = 0
