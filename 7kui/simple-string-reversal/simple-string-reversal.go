package kata

func Solve(s string) string {
	runes := []rune(s)
	result := make([]rune, len(runes))
	for i, r := range runes {
		if r == ' ' {
			result[i] = ' '
		}
	}
	j := len(runes) - 1
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			continue
		}
		for runes[j] == ' ' {
			j--
		}
		result[i] = runes[j]
		j--
	}

	return string(result)
}

/*
In this Kata, we are going to reverse a string while maintaining the spaces (if any) in their original place.

For example:

solve("our code") = "edo cruo"
-- Normal reversal without spaces is "edocruo".
-- However, there is a space at index 3, so the string becomes "edo cruo"

solve("your code rocks") = "skco redo cruoy".
solve("codewars") = "srawedoc"
More examples in the test cases. All input will be lower case letters and in some cases spaces.


*/
