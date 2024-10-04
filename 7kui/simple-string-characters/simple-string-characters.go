package kata

func Solve(s string) []int {
	slice := []byte{}
	resSlice := []int{}
	for _, v := range s {
		slice = append(slice, byte(v))
	}

	upperCase := 0
	lowerCase := 0
	numbers := 0
	specChar := 0
	for _, v := range slice {
		switch {
		case v > 64 && v < 91:
			upperCase += 1
		case v > 96 && v < 123:
			lowerCase += 1
		case v > 47 && v < 58:
			numbers += 1
		default:
			specChar += 1
		}
	}

	resSlice = append(resSlice, upperCase, lowerCase, numbers, specChar)

	return resSlice
}

/*
In this Kata, you will be given a string and your task will be to return a list of ints detailing
the count of uppercase letters, lowercase, numbers and special characters, as follows.

Solve("*'&ABCDabcde12345") = [4,5,5,3].
--the order is: uppercase letters, lowercase, numbers and special characters.
More examples in the test cases.

Good luck!

INPUT :
"P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H"
"*'&ABCDabcde12345"
*/
