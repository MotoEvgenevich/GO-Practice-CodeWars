package kata

func Pyramid(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		sublist := make([]int, i+1)
		for j := 0; j <= i; j++ {
			sublist[j] = 1
		}
		result = append(result, sublist)
	}
	return result
}

/*
Write a function that when given a number >= 0, returns an Array of ascending length subarrays.

pyramid(0) => [ ]
pyramid(1) => [ [1] ]
pyramid(2) => [ [1], [1, 1] ]
pyramid(3) => [ [1], [1, 1], [1, 1, 1] ]
Note: the subarrays should be filled with 1s
*/
