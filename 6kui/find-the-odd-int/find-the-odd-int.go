package kata

func FindOdd(seq []int) int {
	digits := make(map[int]int)
	for _, v := range seq {
		digits[v]++
	}

	for k, v := range digits {
		if v%2 == 1 {
			return k
		}
	}
	return seq[0]
}

/*
Given an array of integers, find the one that appears an odd number of times.

There will always be only one integer that appears an odd number of times.

Examples

[7] should return 7, because it occurs 1 time (which is odd).
[0] should return 0, because it occurs 1 time (which is odd).
[1,1,2] should return 2, because it occurs 1 time (which is odd).
[0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).
[1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it appears 1 time (which is odd).
*/
