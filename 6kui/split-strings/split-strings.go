package kata

func Solution(str string) []string {
	result := make([]string, 0)
	if len(str)%2 == 0 {
		for i := 0; i < len(str); i += 2 {
			digits := str[i : i+2]
			result = append(result, digits)
		}
	}

	if len(str)%2 != 0 {
		for i := 0; i < len(str)-1; i += 2 {
			digits := str[i : i+2]
			result = append(result, digits)
		}
		result = append(result, string(str[len(str)-1])+"_")
	}
	return result
}

/*
Complete the solution so that it splits the string into pairs of two characters.
If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').

Examples:

* 'abc' =>  ['ab', 'c_']
* 'abcdef' => ['ab', 'cd', 'ef']
*/
