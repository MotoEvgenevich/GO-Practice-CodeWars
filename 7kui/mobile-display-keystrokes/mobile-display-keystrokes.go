package kata

var charMap = map[rune]int{
	'1': 1,
	'2': 1,
	'3': 1,
	'4': 1,
	'5': 1,
	'6': 1,
	'7': 1,
	'8': 1,
	'9': 1,
	'0': 1,
	'#': 1,
	'*': 1,
	'a': 2,
	'd': 2,
	'g': 2,
	'j': 2,
	'm': 2,
	'p': 2,
	't': 2,
	'w': 2,
	'b': 3,
	'e': 3,
	'h': 3,
	'k': 3,
	'n': 3,
	'q': 3,
	'u': 3,
	'x': 3,
	'c': 4,
	'f': 4,
	'i': 4,
	'l': 4,
	'o': 4,
	'r': 4,
	'v': 4,
	'y': 4,
	's': 5,
	'z': 5,
}

func MobileKeyboard(str string) int {
	sum := 0
	for _, char := range str {
		if val, ok := charMap[char]; ok {
			sum += val
		}
	}
	return sum
}

/*
Do you remember the old mobile display keyboards? Do you also remember how inconvenient it was to write on it?
Well, here you have to calculate how many keystrokes you have to do for a specific word.

This is the layout:
1 2 3
4 5 6
7 8 9
* 0 #

Given a string, return the number of keystrokes necessary to type it. You can assume that the input will entirely consist of characters included in the mobile layout (lowercase letters, digits, and the symbols * and #).

Examples

"*#"       =>  2 (1 + 1)
"123"      =>  3 (1 + 1 + 1)
"abc"      =>  9 (2 + 3 + 4)
"codewars" => 26 (4 + 4 + 2 + 3 + 2 + 2 + 4 + 5)
*/
