package main

import "fmt"

func main() {
	fmt.Println(Hero(10, 5))
}

func Hero(bullets, dragons int) bool {
	// your code
	if (bullets / 2) >= dragons {
		return true
	} else {
		return false
	}
	return true
}
