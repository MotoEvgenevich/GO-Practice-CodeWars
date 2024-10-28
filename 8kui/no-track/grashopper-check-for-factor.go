package main

import "fmt"

func main() {
	fmt.Println(CheckForFactor(24617, 3))
}

func CheckForFactor(base int, factor int) bool {
	if base%factor == 0 {
		return true
	} else {
		return false
	}

}
