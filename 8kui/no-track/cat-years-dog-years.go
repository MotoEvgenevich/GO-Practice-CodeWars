package main

import "fmt"

func main() {
	fmt.Println(CalculateYears(10))
}

func CalculateYears(years int) (result [3]int) {
	firstYearAnimal := 15
	secondYearAnimal := 9
	totalCat := 0
	totalDog := 0

	switch {
	case years == 1:
		totalCat = firstYearAnimal
		totalDog = firstYearAnimal
		return [3]int{years, totalCat, totalDog}
	case years == 2:
		totalCat = firstYearAnimal + secondYearAnimal
		totalDog = firstYearAnimal + secondYearAnimal
		return [3]int{years, totalCat, totalDog}
	case years >= 3:
		totalCat = firstYearAnimal + secondYearAnimal + (4 * (years - 2))
		totalDog = firstYearAnimal + secondYearAnimal + (5 * (years - 2))
		return [3]int{years, totalCat, totalDog}
	}

	return
}
