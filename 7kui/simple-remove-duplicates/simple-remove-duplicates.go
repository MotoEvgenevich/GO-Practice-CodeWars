package kata

func Solve(arr []int) []int {
	result := []int{}
	similar := false
	for i, digit := range arr {
		for i = i + 1; i < len(arr); i++ {
			if digit == arr[i] {
				similar = true
			}
		}
		if similar {
			similar = false
			continue

		} else {
			result = append(result, digit)
		}

	}
	return result
}

/*
Remove the duplicates from a list of integers, keeping the last ( rightmost ) occurrence of each element.

Example:

For input: [3, 4, 4, 3, 6, 3]

remove the 3 at index 0
remove the 4 at index 1
remove the 3 at index 3
Expected output: [4, 6, 3]

More examples can be found in the test cases.

Good luck!
*/
