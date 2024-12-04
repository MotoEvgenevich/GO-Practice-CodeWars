package kata

func MultiplicationTable(size int) [][]int {
	result := [][]int{}
	for i := 1; i <= size; i++ {
		sublist := make([]int, size)
		for j := 1; j <= size; j++ {
			sublist[j-1] = i * j
		}
		result = append(result, sublist)
	}
	return result
}

/*
Your task, is to create N×N multiplication table, of size provided in parameter.

For example, when given size is 3:

1 2 3
2 4 6
3 6 9
For the given example, the return value should be:

[[1,2,3],[2,4,6],[3,6,9]]
*/
