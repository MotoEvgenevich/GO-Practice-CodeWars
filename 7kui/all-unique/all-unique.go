package kata

func HasUniqueChar(str string) bool {
	digits := make(map[rune]int)

	for _, char := range str {
		digits[char]++
		if digits[char] > 1 {
			return false
		}
	}
	return true

}

/*
Write a program to determine if a string contains only unique characters. Return true if it does and false otherwise.

The string may contain any of the 128 ASCII characters. Characters are case-sensitive, e.g. 'a' and 'A' are considered different characters.
*/
