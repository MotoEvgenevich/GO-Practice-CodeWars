package kata

func LengthOfSequence(arr []int, key int) int {
	startIndex := 0
	endIndex := 0
	counter := 0
	for i, v := range arr {
		if v == key {
			startIndex = i + 1
			break
		}
	}
	for i := startIndex; i < len(arr); i++ {
		if arr[i] == key {
			endIndex = i + 1
			break
		}
	}
	for _, v := range arr {
		if v == key {
			counter++
		}
	}
	switch {
	case counter > 2:
		return 0
	case startIndex == 0 || endIndex == 0:
		return 0
	case startIndex != 0 && endIndex != 0:
		return endIndex - startIndex + 1

	}

	return 0
}

/*
 As part of this kata, you need to find the length of the sub-array that begins and ends with the specified number.

For example, if the array arr is [0, -3, 7, 4, 0, 3, 7, 9], finding the length of the sub-array that begins and ends with 7s would return 5.

If there are less or more than two occurrences of the number to search for, return 0.
*/
