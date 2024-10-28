package main

import (
	"fmt"
)

func main() {
	word := "Word word word"
	complete := NoSpace(word)
	fmt.Println(complete)
}

func NoSpace(word string) string {
	slice := []byte{}
	space := " "
	for i := 0; i < len(word); i++ {
		if string(word[i]) != space {
			slice = append(slice, word[i])
		}
	}
	return string(slice)
}
