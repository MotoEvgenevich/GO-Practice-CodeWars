/* Given the triangle of consecutive odd numbers:

             1
          3     5
       7     9    11
   13    15    17    19
21    23    25    27    29
...
Calculate the sum of the numbers in the nth row of this triangle (starting at index 1) e.g.: (Input --> Output)

1 -->  1
2 --> 3 + 5 = 8 */

package main

import "fmt"

func main() {
	fmt.Println(RowSumOddNumbers(1))
	fmt.Println(RowSumOddNumbers(2))
	fmt.Println(RowSumOddNumbers(3))
	fmt.Println(RowSumOddNumbers(4))
	fmt.Println(bestPractice(1))
	fmt.Println(bestPractice(2))
	fmt.Println(bestPractice(3))
	fmt.Println(bestPractice(4))
}

func RowSumOddNumbers(n int) int {
	if n == 1 {
		return 1
	}
	startOddNumberPosition := 0
	startOddNumber := 1
	sum := 0
	for i := 0; i < n; i++ {
		startOddNumberPosition += i
	}
	for i := startOddNumber; i < startOddNumberPosition; i++ {
		startOddNumber += 2
	}
	for i := n; i > 0; i-- {
		sum += startOddNumber + 2
		startOddNumber += 2
	}
	return sum
}

func bestPractice(n int) int {
	return n * n * n
}
