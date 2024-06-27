package kata

import "fmt"

func Spacify(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	for _, char := range s {
		result += fmt.Sprint(string(char), " ")
	}
	return result[:len(result)-1]
}

/*
Modify the spacify function so that it returns the given string with spaces inserted between each character.

spacify=("hello world") -> # returns "h e l l o   w o r l d
*/
