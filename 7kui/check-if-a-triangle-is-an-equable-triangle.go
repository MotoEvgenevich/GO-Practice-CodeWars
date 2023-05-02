/* A triangle is called an equable triangle if its area equals its perimeter.
Return true, if it is an equable triangle, else return false.
You will be provided with the length of sides of the triangle. Happy Coding!
*/

package main

import "fmt"

func main() {
	fmt.Println(EquableTriangle(5, 12, 13))
}

func EquableTriangle(a, b, c int) bool {
	perimeter := a + b + c
	fmt.Println(perimeter)
	area := 0
	p := (a + b + c) / 2

	sBfSquare := p * (p - a) * (p - b) * (p - c)
	fmt.Println(sBfSquare)

	for i := 1; i <= sBfSquare; i++ {
		area = i * i
		if area == sBfSquare {
			area = i
			fmt.Println(area)
			break
		}

	}
	fmt.Println(area)
	if area == perimeter {
		return true
	} else {
		return false
	}
}
