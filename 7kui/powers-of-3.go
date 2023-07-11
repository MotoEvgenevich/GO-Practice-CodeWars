/* Given a positive integer N, return the largest integer k such that 3^k < N.

For example,

LargestPower(3) // returns 0
LargestPower(4) // returns 1
You may assume that the input to your function is always a positive integer.
TESTS:
in_ := [...]uint64{3, 5, 7, 81, 82}
    out := [...]int{0, 1, 1, 3, 4}
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(LargestPower(82))
}

func LargestPower(n uint64) int {
	border := int(n)
	const Num float64 = 3.0
	k := 0
	var result float64 = 0

	for i := 0; int(result) < border; i++ {
		k = i
		result = math.Pow(Num, float64(k))
	}
	return int(k - 1) // return the largest integer k such that 3^k < n
}

//BEST PRACTICE

func LargestPower1(n uint64) int {
	k := 0.0
	for math.Pow(3, k) < float64(n) {
		k++
	}
	return int(k - 1)
}
