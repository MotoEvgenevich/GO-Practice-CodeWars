package main

import (
	"fmt"
	"strings"
)

func main() {
	RepeatStr1(4, "hello")
	fmt.Println(RepeatStr1(5, "hello"))
}

func RepeatStr(repetitions int, value string) string {
	slice := []string{}
	for i := 0; i < repetitions; i++ {
		slice = append(slice, value)
	}
	fmt.Println(slice)
	result2 := strings.Join(slice, "")
	return result2
}

func RepeatStr1(repetitions int, value string) string {
	result := ""
	for i := 0; i <= repetitions; i++ {
		result += value
	}
	return result
}
