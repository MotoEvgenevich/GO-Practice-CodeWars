package main

import "fmt"

func main() {
	fmt.Println(ToAlternatingCase("HeLLO WORLD 12344"))
}

func ToAlternatingCase(str string) string {
	var slice []byte
	var altSlice []byte
	for _, char := range str {
		slice = append(slice, byte(char))
	}
	fmt.Println(string(slice))
	for _, v := range slice {
		if v >= 65 && v <= 90 {
			altSlice = append(altSlice, v+32)
		}
		if v >= 97 && v <= 122 {
			altSlice = append(altSlice, v-32)
		}
		if v == 32 {
			altSlice = append(altSlice, v)
		}
		if v >= 0 && v <= 31 || v >= 33 && v <= 64 {
			altSlice = append(altSlice, v)
		}
	}
	fmt.Println(string(altSlice))
	return string(altSlice)
}
