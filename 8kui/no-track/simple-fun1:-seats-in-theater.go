package main

import "fmt"

func main() {
	fmt.Println(seatsInTheater(1000, 1000, 1000, 1000))
}

func seatsInTheater(nCols int, nRows int, col int, row int) int {

	return (nCols - (col - 1)) * (nRows - row)
}
