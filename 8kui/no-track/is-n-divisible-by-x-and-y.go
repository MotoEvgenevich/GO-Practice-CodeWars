package main

import "fmt"

func main() {
	fmt.Println(IsDivisible(12, 3, 4))
}

func IsDivisible(n, x, y int) bool {

	return n%x == 0 && n%y == 0
}

func IsDivisible1(n, x, y int) bool {

	if n%x == 0 && n%y == 0 {
		return true
	} else {
		return false
	}
	// your code here
}
