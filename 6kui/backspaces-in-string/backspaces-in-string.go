package kata

func CleanString(s string) string {
	result := ""
	if len(s) == 0 {
		return ""
	}
	for i := 0; i < len(s); i++ {
		if string(s[i]) != "#" {
			result += string(s[i])
		} else {
			if i > 0 {
				result = result[result[i-1]:result[i+1]]
			}
		}
	}
	return result
}

/*
Assume "#" is like a backspace in string. This means that string "a#bc#d" actually is "bd"

Your task is to process a string with "#" symbols.

Examples

"abc#d##c"      ==>  "ac"
"abc##d######"  ==>  ""
"#######"       ==>  ""
""              ==>  ""
*/
