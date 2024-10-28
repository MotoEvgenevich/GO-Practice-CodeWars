package main

import "fmt"

func main() {
	fmt.Println(SmallestIntegerFinder([]int{34, -345, -1, 100}))
}

func SmallestIntegerFinder(numbers []int) int {
	small := numbers[0]

	for _, v := range numbers {
		if small > v {
			small = v
		}
	}
	return small
}
