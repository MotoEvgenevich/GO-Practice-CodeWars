package kata

func Solve(arr []int) int {
	sum := 0
	for i := 2; i < len(arr); i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			sum += arr[i]
		}
	}
	return sum
}

/*
In this Kata, you will be given an integer array and your task is to return the sum of elements occupying prime-numbered indices.

The first element of the array is at index 0.

Good luck!
*/

// Функция для проверки, является ли число простым
