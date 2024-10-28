package main

import (
	"fmt"
)

func main() {

	StringToNumber("-123234345")
	fmt.Println(StringToNumber("-123234345"))

}

func StringToNumber(n string) int {
	slice := []int{}

	result := 0
	for _, value := range n {

		digit := string(value)
		//fmt.Println(digit)
		//fmt.Printf("%v %T\n", digit, digit)
		dig := digitStrToInt(digit)
		//fmt.Printf("%v %T\n", dig, dig)
		slice = append(slice, dig)
		//fmt.Println("ищем минус:", value)

	}
	for i := 0; i < len(slice); i++ {
		result += slice[len(slice)-1-i] * tenStepen(i)

		//fmt.Println("значение сллайса:", slice[i])
		//fmt.Println("значение степени:", tenStepen(i))
		//fmt.Println("result=", result)
	}

	//fmt.Println(slice)
	//fmt.Println(result)

	for _, value := range n {
		if value == 45 {
			return -(result)
		}
	}
	return result
}

func tenStepen(stepen int) int {
	if stepen == 0 {
		return 1
	}
	if stepen == 1 {
		return 10
	}
	if stepen == 2 {
		return 100
	}
	if stepen == 3 {
		return 1000
	}
	if stepen == 4 {
		return 10000
	}
	if stepen == 5 {
		return 100000
	}
	if stepen == 6 {
		return 1000000
	}
	if stepen == 7 {
		return 10000000
	}
	if stepen == 8 {
		return 100000000
	}
	if stepen == 9 {
		return 1000000000
	}

	panic(stepen)
}

func digitStrToInt(s string) int {

	if s == "1" {
		return 1
	}
	if s == "2" {
		return 2
	}
	if s == "3" {
		return 3
	}
	if s == "4" {
		return 4
	}
	if s == "5" {
		return 5
	}
	if s == "6" {
		return 6
	}
	if s == "7" {
		return 7
	}
	if s == "8" {
		return 8
	}
	if s == "9" {
		return 9
	}
	if s == "0" {
		return 0
	}
	if s == "-" {
		return -0
	}
	panic("bad value: " + s)

}
