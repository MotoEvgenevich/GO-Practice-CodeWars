package main

func main() {

}

func PositiveSum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		if value > 0 {
			sum += value
		}
	}
	return sum
}
