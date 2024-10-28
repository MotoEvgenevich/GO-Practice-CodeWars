package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(PowersOfTwo(10))
}
func PowersOfTwo(n int) []uint64 {
	result := []uint64{}
	for i := 0; i < n; i++ {
		dig := math.Pow(2, float64(i))
		fmt.Println(dig)
		result = append(result, uint64(dig))
	}
	return result
}
