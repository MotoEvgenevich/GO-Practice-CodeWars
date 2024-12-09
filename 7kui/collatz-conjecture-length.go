/*
The Collatz Conjecture states that for any positive natural number n, this process:

if n is even, divide it by 2
if n is odd, multiply it by 3 and add 1
repeat
will eventually reach n = 1.

For example, if n = 20, the resulting sequence will be:

[ 20, 10, 5, 16, 8, 4, 2, 1 ]
Write a program that will output the length of the Collatz Conjecture for any given n.
In the example above, the output would be 8.
*/

package main

import "fmt"

func main() {
	fmt.Println(Collatz(20))
}

func Collatz(n int) int {
	reps := 1
	i := n
	for i > 1 {
		switch {
		case i%2 == 0:
			i = i / 2
			reps += 1
		case i%2 != 0:
			i = i*3 + 1
			reps += 1
		}

	}
	return reps
}
