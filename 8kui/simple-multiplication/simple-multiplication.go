package kata

func SimpleMultiplication(n int) int {
	if n%2 == 0 {
		return n * 8
	} else {
		return n * 9
	}
}

/*
This kata is about multiplying a given number by eight if it is an even number and by nine otherwise.
*/
