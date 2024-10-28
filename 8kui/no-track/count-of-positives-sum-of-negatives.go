package main

import "fmt"

func main() {
	fmt.Println(CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
}

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	sumPos := 0
	sumNeg := 0
	for _, v := range numbers {
		switch {
		case v > 0:
			sumPos += 1
		case v < 0:
			sumNeg += v
		}
	}
	res = append(res, sumPos, sumNeg)
	return res // your code here
}
