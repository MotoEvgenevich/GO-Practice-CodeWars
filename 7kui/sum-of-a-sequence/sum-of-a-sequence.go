package kata

func SequenceSum(start, end, step int) int {

	var totalSum int = 0
	if start > end {
		return 0
	}
	if start <= end {
		for i := start; i <= end; i += step {
			totalSum += i
		}
	}
	return totalSum
}

/*
Your task is to write a function which returns the sum of a sequence of integers.

The sequence is defined by 3 non-negative values: begin, end, step.

If begin value is greater than the end, your function should return 0.
If end is not the result of an integer number of steps, then don't add it to the sum. See the 4th example below.

Examples
2,2,2 --> 2
2,6,2 --> 12 (2 + 4 + 6)
1,5,1 --> 15 (1 + 2 + 3 + 4 + 5)
1,5,3  --> 5 (1 + 4)

INPUT:
(2, 6, 2, 12) (2 + 4 + 6)
(1, 5, 1, 15)
(1, 5, 3, 5)
(0, 15, 3, 45)
(16, 15, 3, 0)
(2, 24, 22, 26)
(2, 2, 2, 2)
(2, 2, 1, 2)
(1, 15, 3, 35) (1+4+7+10+13)
(15, 1, 3, 0)

*/
