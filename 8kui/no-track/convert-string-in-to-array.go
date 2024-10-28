package main

import "fmt"

func main() {
	fmt.Println(StringToArray("Robin Singh"))
}

func StringToArray(str string) []string {
	arr := []rune{}
	for _, v := range str {
		arr = append(arr, v)
	}
	word := ""
	newArr := []string{}
	for _, v := range arr {
		if v != 32 {
			word += string(v)

		}
		if v == 32 {
			newArr = append(newArr, word)
			word = ""
		}
	}
	newArr = append(newArr, word)
	return newArr
}
