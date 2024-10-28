package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(IsUpperCase(MyString("JFHJKSDFH ")))
}

type MyString string

func IsUpperCase(s MyString) bool {
	// Your code here!
	BoolTmp := false
	slice := []byte{}
	for _, val := range s {
		slice = append(slice, byte(val))
	}
	fmt.Println(slice)
	for _, b := range slice {
		if b >= 65 && b <= 90 || b == 32 {
			BoolTmp = true
		} else {
			return false
		}

	}
	return BoolTmp
}

func (s MyString) IsUpperCase() bool {
	// Your code here!
	var status bool
	for _, r := range s {
		if unicode.IsLower(r) {
			status = false
			break
		} else {
			status = true
		}
	}
	return status
}
