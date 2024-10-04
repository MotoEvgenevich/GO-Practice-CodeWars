package kata

func Alternate(n int, firstValue string, secondValue string) []string {
	slice := []string{}
	if n == 0 {
		return slice
	}
	if n%2 == 0 {
		for i := 0; i < n/2; i++ {
			slice = append(slice, firstValue, secondValue)
		}
		return slice
	}
	if n%2 != 0 {
		for i := 0; i < (n-1)/2; i++ {
			slice = append(slice, firstValue, secondValue)
		}
		slice = append(slice, firstValue)
		return slice
	}
	return slice
}

/*
Given an integer n and two other values, build an array of size n filled with these two values alternating.

Examples

5, true, false     -->  [true, false, true, false, true]
10, "blue", "red"  -->  ["blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"]
0, "one", "two"    -->  []
*/
