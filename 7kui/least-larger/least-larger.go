package kata

func LeastLarger(a []int, i int) int {
	target := a[i]
	minLarger := 9223372036854775807 //maxint64
	minLargerIndex := -1

	for index, v := range a {
		if v > target && v < minLarger {
			minLarger = v
			minLargerIndex = index
		}
	}

	return minLargerIndex
}

/*
Task

Given an array of numbers and an index, return either the index of the smallest number that is larger than the element at the given index,
or -1 if there is no such index ( or, where applicable, Nothing or a similarly empty value ).

Notes

Multiple correct answers may be possible. In this case, return any one of them.
The given index will be inside the given array.
The given array will, therefore, never be empty.

Example

leastLarger( [] int {4, 1, 3, 5, 6}, 0 )  =  3
leastLarger( [] int {4, 1, 3, 5, 6}, 4 )  = -1
LeastLarger([]int{4, 1, 3, 5, 6}, 0)  // returns 3
LeastLarger([]int{4, 1, 3, 5, 6}, 4)  // returns -1

*/
