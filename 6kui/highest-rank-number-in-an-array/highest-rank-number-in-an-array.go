package main

import "fmt"

func main() {
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}))
}

func HighestRank(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v] += 1
	}
	maxValue := 0
	maxKey := 0
	for i, v := range m {
		if v > maxValue || (v == maxValue && i > maxKey) {
			maxValue = v
			maxKey = i
		}
	}
	return maxKey
}

/*
Complete the method which returns the number which is most frequent in the given input array.
If there is a tie for most frequent number, return the largest number among them.

Note: no empty arrays will be given.

Examples

[12, 10, 8, 12, 7, 6, 4, 10, 12]              -->  12
[12, 10, 8, 12, 7, 6, 4, 10, 12, 10]          -->  12
[12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10]  -->   3 */
