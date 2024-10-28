package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Digitize(35231))
}

func Digitize(n int) []int {
	convStringToNum := ""
	convStringToNum = strconv.Itoa(n)
	stringSlice := []string{}
	intSlice := []int{}
	fmt.Println(convStringToNum)
	for _, v := range convStringToNum {
		stringSlice = append(stringSlice, string(v))

	}
	fmt.Println(stringSlice)

	return nil
}
