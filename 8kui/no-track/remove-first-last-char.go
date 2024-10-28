package main

import "fmt"

func main() {
	RemoveChar("maybe")
	fmt.Println(RemoveChar("Satana"))
}

func RemoveChar(word string) string {
	slice := []byte{}

	for i, value := range word {
		if i > 0 && i < len(word)-1 {
			slice = append(slice, byte(value))
		}

	}
	return string(slice)
}
