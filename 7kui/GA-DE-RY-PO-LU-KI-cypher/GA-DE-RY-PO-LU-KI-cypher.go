package kata

func Encode(str string) string {
	result := ""
	for _, char := range str {
		switch char {
		case 'G':
			result += "A"
		case 'A':
			result += "G"
		case 'D':
			result += "E"
		case 'E':
			result += "D"
		case 'R':
			result += "Y"
		case 'Y':
			result += "R"
		case 'P':
			result += "O"
		case 'O':
			result += "P"
		case 'L':
			result += "U"
		case 'U':
			result += "L"
		case 'K':
			result += "I"
		case 'I':
			result += "K"
		case 'g':
			result += "a"
		case 'a':
			result += "g"
		case 'd':
			result += "e"
		case 'e':
			result += "d"
		case 'r':
			result += "y"
		case 'y':
			result += "r"
		case 'p':
			result += "o"
		case 'o':
			result += "p"
		case 'l':
			result += "u"
		case 'u':
			result += "l"
		case 'k':
			result += "i"
		case 'i':
			result += "k"
		default:
			result += string(char)
		}
	}
	return result
}

func Decode(str string) string {
	result := ""
	for _, char := range str {
		switch char {
		case 'A':
			result += "G"
		case 'G':
			result += "A"
		case 'E':
			result += "D"
		case 'D':
			result += "E"
		case 'Y':
			result += "R"
		case 'R':
			result += "Y"
		case 'O':
			result += "P"
		case 'P':
			result += "O"
		case 'U':
			result += "L"
		case 'L':
			result += "U"
		case 'I':
			result += "K"
		case 'K':
			result += "I"
		case 'a':
			result += "g"
		case 'g':
			result += "a"
		case 'e':
			result += "d"
		case 'd':
			result += "e"
		case 'y':
			result += "r"
		case 'r':
			result += "y"
		case 'o':
			result += "p"
		case 'p':
			result += "o"
		case 'u':
			result += "l"
		case 'l':
			result += "u"
		case 'i':
			result += "k"
		case 'k':
			result += "i"
		default:
			result += string(char)
		}
	}
	return result
}

func EncodeWithMap(str string) string {
	// Создаем карту замен
	replacements := map[rune]rune{
		'G': 'A', 'A': 'G',
		'D': 'E', 'E': 'D',
		'R': 'Y', 'Y': 'R',
		'P': 'O', 'O': 'P',
		'L': 'U', 'U': 'L',
		'K': 'I', 'I': 'K',
		'g': 'a', 'a': 'g',
		'd': 'e', 'e': 'd',
		'r': 'y', 'y': 'r',
		'p': 'o', 'o': 'p',
		'l': 'u', 'u': 'l',
		'k': 'i', 'i': 'k',
	}

	result := ""
	for _, char := range str {
		if replacement, ok := replacements[char]; ok {
			result += string(replacement)
		} else {
			result += string(char)
		}
	}
	return result
}

/*
Introduction

The GA DE RY PO LU KI is a simple substitution cypher used in scouting to encrypt messages. The encryption is based on short, easy to remember key.
The key is written as paired letters, which are in the cipher simple replacement.

The most frequently used key is "GA-DE-RY-PO-LU-KI".

 G => A
 g => a
 a => g
 A => G
 D => E
  etc.
The letters, which are not on the list of substitutes, stays in the encrypted text without changes.

Task

Your task is to help scouts to encrypt and decrypt thier messages. Write the Encode and Decode functions.

Input/Output

The input string consists of lowercase and uperrcase characters and white . The substitution has to be case-sensitive.

Example

 Encode("ABCD")          // => GBCE
 Encode("Ala has a cat") // => Gug hgs g cgt
 Encode("gaderypoluki")  // => agedyropulik
 Decode("Gug hgs g cgt") // => Ala has a cat
 Decode("agedyropulik")  // => gaderypoluki
 Decode("GBCE")          // => ABCD
*/
