package main

import "fmt"

func main() {
	fmt.Println(Heron(3.0, 4.0, 5.0))
}

func Heron(a, b, c float64) (area float64) {
	fmt.Println(a, b, c)
	area = 6.00
	fmt.Println(area)
	s := (a + b + c) / 2
	fmt.Println(s)
	area = (s * (s - a) * (s - b) * (s - c))
	fmt.Println(area)

	return (area)
}

/* Write function heron which calculates the area of a triangle with sides a, b, and c (x, y, z in COBOL).

Heron's formula:

Output should have 2 digits precision.

input:
3.0, 4.0, 5.0,
output
6.0
*/
