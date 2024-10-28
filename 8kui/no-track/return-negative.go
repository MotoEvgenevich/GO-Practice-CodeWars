package main

import "fmt"

func main() {
	fmt.Println(MakeNegative(0))
}

func MakeNegative(x int) int {
	switch {
	case x < 0:
		return x
	case x > 0:
		return -(x)
	}
	return 0
}
