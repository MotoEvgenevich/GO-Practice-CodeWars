package main

import "fmt"

func main() {
	fmt.Println(NearestSq(111))
}

func NearestSq(n int) int {
	squareBig := 1
	squareSmall := 1
	compMin := 0
	compMax := 0
	for i := 1; squareBig < n; i++ {
		squareBig = i * i
		squareSmall = (i - 1) * (i - 1)

	}

	compMax = squareBig - n
	compMin = n - squareSmall
	fmt.Println(compMax)
	fmt.Println(compMin)
	if compMax > compMin {
		return squareSmall
	} else {
		return squareBig
	}

}
