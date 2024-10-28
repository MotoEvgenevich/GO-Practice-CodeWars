package main

import "fmt"

func main() {
	fmt.Println(Solution("Soon", "Me"))
}

func Solution(a, b string) string {
	res := ""
	if len(a) > len(b) {
		res = b + a + b
	}
	if len(a) < len(b) {
		res = a + b + a
	}
	return res
}
