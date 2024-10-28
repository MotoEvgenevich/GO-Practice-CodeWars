package main

import "fmt"

func main() {
	fmt.Println(LoveFunc(3, 2))
}

func LoveFunc(flower1, flower2 int) bool {
	if flower1%2 == 0 && flower2%2 != 0 {
		return true
	}
	if flower1%2 != 0 && flower2%2 == 0 {
		return true
	} else {
		return false
	}

}
