package kata

func InAscOrder(numbers []int) bool {
	queue := true
	firstDigit := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if firstDigit < numbers[i] {
			queue = true
			firstDigit = numbers[i]
		} else {
			return false
		}
	}
	return queue
}

/* Are the numbers in order?

In this Kata, your function receives an array of integers as input.
Your task is to determine whether the numbers are in ascending order.
An array is said to be in ascending order if there are no two adjacent integers
where the left integer exceeds the right integer in value.

For the purposes of this Kata, you may assume that all inputs are valid, i.e. arrays containing only integers.

Note that an array of 0 or 1 integer(s) is automatically considered to be sorted
in ascending order since all (zero) adjacent pairs of integers satisfy the condition
that the left integer does not exceed the right integer in value.

*/
