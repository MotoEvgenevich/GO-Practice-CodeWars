package kata

func Smaller(arr []int) []int {
	result := []int{}
	count := 0
	digit := 0
	for index, v := range arr {
		digit = v
		for i := index + 1; i < len(arr); i++ {
			if digit > arr[i] {
				count++
			}
		}
		result = append(result, count)
		count = 0
	}
	return result
}

/*

OLD VERSION SOLUTION

func Smaller(arr []int) []int {
	result := []int{}
	sum := 0
	for i, num := range arr {
		for j := i + 1; j < len(arr); j++ {
			if num > arr[j] {
				sum += 1
			}
		}
		result = append(result, sum)
		sum = 0
	}
	fmt.Println(result)
	return result
}

*/

/* Write a function that given, an array arr,
returns an array containing at each index i the amount of numbers that are smaller than arr[i] to the right.

For example:

* Input [5, 4, 3, 2, 1] => Output [4, 3, 2, 1, 0]
* Input [1, 2, 0] => Output [1, 1, 0] */
