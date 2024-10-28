package main

import "fmt"

func main() {
	fmt.Println(Contamination("_3ebzgh4", "&"))
}

func Contamination(text, char string) string {
	res := ""
	for i := 0; i < (len(text)); i++ {
		res += char
	}
	return res
}
