package kata

func GetSum(a, b int) int {
	start := 0
	finish := 0
	slice := []int{}
	sum := 0
	switch {
	case a == b:
		return a
	case a <= b:
		start = a
		finish = b
	case a >= b:
		start = b
		finish = a
	}
	for i := start; i <= finish; i++ {
		slice = append(slice, i)
	}
	for _, v := range slice {
		sum += v
	}
	return sum
}

/* Given two integers a and b, which can be positive or negative,
find the sum of all the integers between and including them and return it.
If the two numbers are equal return a or b.

Note: a and b are not ordered!

Examples (a, b) --> output (explanation)

(1, 0) --> 1 (1 + 0 = 1)
(1, 2) --> 3 (1 + 2 = 3)
(0, 1) --> 1 (0 + 1 = 1)
(1, 1) --> 1 (1 since both are same)
(-1, 0) --> -1 (-1 + 0 = -1)
(-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)
Your function should only return a number,
not the explanation about how you get that number. */
