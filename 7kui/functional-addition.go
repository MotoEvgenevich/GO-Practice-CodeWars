/* Create a function add(n)/Add(n) which returns a function that
always adds n to any number

Note for Java: the return type and methods have not
been provided to make it a bit more challenging.

var addOne = Add(1)
addOne(3) // 4 */

package main

import "fmt"

func main() {
	addOne := Add(1)
	result := addOne(3)
	fmt.Println(result)
}

func Add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
