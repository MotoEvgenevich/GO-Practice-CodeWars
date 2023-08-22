package kata

func Capitalize(st string) []string {
	result := []string{}
	word1 := ""
	word2 := ""
	for i, v := range st {
		if i%2 == 0 {
			word1 += string(v - 32)
			word2 += string(v)
		}
		if i%2 != 0 {
			word1 += string(v)
			word2 += string(v - 32)
		}
	}
	result = append(result, word1, word2)

	return result
}

/* Given a string, capitalize the letters that occupy even indexes and odd indexes separately, and return as shown below. Index 0 will be considered even.

For example, capitalize("abcdef") = ['AbCdEf', 'aBcDeF']. See test cases for more examples.

The input will be a lowercase string with no spaces.

[]string{"CoDeWaRs", "cOdEwArS"})

Good luck! */
