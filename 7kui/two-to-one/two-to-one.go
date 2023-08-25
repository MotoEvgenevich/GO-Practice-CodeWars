package kata

func TwoToOne(s1 string, s2 string) string {
	result := ""
	for i := 97; i <= 122; i++ {
		for _, v := range s1 + s2 {
			if rune(i) == v {
				result += string(v)
				break
			}

		}

	}
	return result
}

/* Take 2 strings s1 and s2 including only letters from a to z. Return a new sorted string,
the longest possible, containing distinct letters - each taken only once - coming from s1 or s2.

a = "xyaabbbccccdefww"
b = "xxxxyyyyabklmopq"
longest(a, b) -> "abcdefklmopqwxy"

a = "abcdefghijklmnopqrstuvwxyz"
longest(a, a) -> "abcdefghijklmnopqrstuvwxyz" */
