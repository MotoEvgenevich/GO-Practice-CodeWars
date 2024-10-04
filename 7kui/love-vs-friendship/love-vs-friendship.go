package kata

func WordsToMarks(s string) int {
	slice := []byte{}
	sumOfword := 0

	for _, v := range s {
		slice = append(slice, byte(v))
	}

	for _, v := range slice {
		sumOfword += (int(v) - 96)
	}
	return sumOfword
}

/* If　a = 1, b = 2, c = 3 ... z = 26

Then l + o + v + e = 54

and f + r + i + e + n + d + s + h + i + p = 108

So friendship is twice as strong as love :-)

Your task is to write a function which calculates the value of a word based off the sum of the alphabet
positions of its characters.

The input will always be made of only lowercase letters and will never be empty. */
