package kata

func SortNumbers(numbers []int) []int {
	maxNmb := 0
	SortedNmb := []int{}

	for _, v := range numbers {
		if maxNmb <= v {
			maxNmb = v
		}
	}
	minNmb := maxNmb
	for _, v := range numbers {
		if minNmb >= v {
			minNmb = v

		}
	}

	for i := minNmb; i <= maxNmb; i++ {
		for _, v := range numbers {
			if i == v {
				SortedNmb = append(SortedNmb, v)
			}
		}
	}

	return SortedNmb
}

func SortNumbersV2(numbers []int) []int {

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[j] > numbers[i] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}

	}

	return numbers // your code here
}

/*
Finish the solution so that it sorts the passed in array of numbers.
If the function passes in an empty array or null/nil value then it should return an empty array.

For example:

solution(c(1, 2, 3, 10, 5)) # should return c(1, 2, 3, 5, 10)
solution(NULL)              # should return NULL
*/
