/*
Return the Nth Even Number

Example(Input --> Output)

1 --> 0 (the first even number is 0)
3 --> 4 (the 3rd even number is 4 (0, 2, 4))
100 --> 198
1298734 --> 2597466
The input will not be 0. */

package main

import "fmt"

func main() {
	fmt.Println(NthEven(1298734))
}

func NthEven(n int) int {
	return n*2 - 2
}
