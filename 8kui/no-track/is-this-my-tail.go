package main

import "fmt"

func main() {
	fmt.Println(CorrectTail("Fox", 'x'))
}

func CorrectTail(body string, tail rune) bool {
	fmt.Println(string(body[len(body)-1]))
	if rune(body[len(body)-1]) == tail {

		return true
	}
	fmt.Println(body[0])
	return false
}
