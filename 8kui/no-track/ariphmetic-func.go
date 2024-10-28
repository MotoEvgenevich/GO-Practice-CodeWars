package main

import "fmt"

func main() {

	fmt.Println(Arithmetic(10, 8, "subtract"))
}

func Arithmetic(a int, b int, operator string) int {
	result := 0

	if operator == "add" {
		result = a + b
	}
	if operator == "subtract" {
		result = a - b
	}
	if operator == "multiply" {
		result = a * b
	}
	if operator == "divide" {
		result = a / b
	}
	return result
}
