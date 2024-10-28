package main

import "fmt"

func main() {
	fmt.Println(AbbrevName("Sam Harris"))
}

func AbbrevName(name string) string {
	slice := []rune{}
	surnameSlice := []rune{}
	totSlice := []rune{}
	result := ""
	word := ""

	for _, v := range name {
		slice = append(slice, v)
	}

	totSlice = append(totSlice, slice[0])
	totSlice = append(totSlice, 32)

	for _, v := range slice {
		if v != 32 {
			word += string(v)

		}
		if v == 32 {
			word = ""
		}
	}

	for _, v := range word {
		surnameSlice = append(surnameSlice, v)
	}

	totSlice = append(totSlice, surnameSlice[0])

	for _, v := range totSlice {
		switch {
		case v <= 122 && v >= 97:
			result += string(v - 32)
		case v == 32:
			result += string(v + 14)
		default:
			result += string(v)
		}

	}
	return result
}
