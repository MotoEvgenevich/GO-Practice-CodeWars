package kata

func Dup(arr []string) []string {
	result := []string{}
	for _, word := range arr {
		clearWord := ""
		for i := 0; i < len(word); i++ {
			if len(word)-1 == i {
				clearWord += string(word[i])
				break
			}
			if word[i] == word[i+1] {
				continue
			} else {
				clearWord += string(word[i])
			}
		}
		result = append(result, clearWord)
	}
	return result
}

/*
In this Kata, you will be given an array of strings and your task is to remove all consecutive duplicate letters from each string in the array.

For example:

dup(["abracadabra","allottee","assessee"]) = ["abracadabra","alote","asese"].

dup(["kelless","keenness"]) = ["keles","kenes"].

Strings will be lowercase only, no spaces. See test cases for more examples.

Good luck!
*/
