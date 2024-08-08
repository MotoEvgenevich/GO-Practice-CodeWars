package kata

func ReverseLetters(s string) string {
	reverseString := ""
	for i := len(s) - 1; i >= 0; i-- {
		reverseString += string(s[i])
	}
	clearReverseSring := ""
	for _, char := range reverseString {
		if char > 64 && char < 91 {
			clearReverseSring += string(char)
		}
		if char > 96 && char < 123 {
			clearReverseSring += string(char)
		}
	}
	return clearReverseSring
}

/*
Task

Given a string str, reverse it and omit all non-alphabetic characters.

Example

For str = "krishan", the output should be "nahsirk".

For str = "ultr53o?n", the output should be "nortlu".

Input/Output

[input] string str
A string consists of lowercase latin letters, digits and symbols.

[output] a string
*/
