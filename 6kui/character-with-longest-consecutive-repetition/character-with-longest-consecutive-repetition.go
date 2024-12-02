package kata

type Result struct {
	C rune // character
	L int  // count
}

func LongestRepetition(text string) Result {
	if text == "" {
		return Result{}
	}

	var maxChar rune
	maxLen := 0
	var currentChar rune
	currentLen := 0

	for _, v := range text {
		if v == currentChar {
			currentLen++
		} else {
			currentChar = v
			currentLen = 1
		}

		if currentLen > maxLen {
			maxLen = currentLen
			maxChar = currentChar
		}
	}

	return Result{C: maxChar, L: maxLen}
}

/*
For a given string s find the character c (or C) with longest consecutive repetition and return:

type Result struct {
    C rune // character
    L int  // count
}
where l (or L) is the length of the repetition. If there are two or more characters with the same l return the first in order of appearance.

For empty string return:

Result{}
Happy coding! :)

*/
