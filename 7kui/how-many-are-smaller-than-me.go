/* Write a function that given, an array arr,
returns an array containing at each index i the amount of
numbers that are smaller than arr[i] to the right.

For example:

* Input [5, 4, 3, 2, 1] => Output [4, 3, 2, 1, 0]
* Input [1, 2, 0] => Output [1, 1, 0] */

package main

import "fmt"

func main() {
	fmt.Println(Smaller([]int{1, 2, 0}))
}

func Smaller(arr []int) []int {
	result := []int{}
	sum := 0
	for i, num := range arr {
		for j := i + 1; j < len(arr); j++ {
			if num > arr[j] {
				sum += 1
			}
		}
		result = append(result, sum)
		sum = 0
	}
	fmt.Println(result)
	return result
}
