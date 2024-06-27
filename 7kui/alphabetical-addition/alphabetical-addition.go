package kata

func AddLetters(letters []rune) rune {
	// your code here
	if len(letters) == 0 {
		return 'z'
	}
	sum := 0
	for _, letter := range letters {
		sum += int(letter - 'a' + 1)
	}
	//fmt.Println(sum)
	sum = (sum-1)%26 + 1
	return rune(sum + 'a' - 1)
}

/*
Your task is to add up letters to one letter.

The function will be given a slice of letters (runes), each one being a letter to add. Return a rune.

Notes:

Letters will always be lowercase.
Letters can overflow (see second to last example of the description)
If no letters are given, the function should return 'z'
Examples:

AddLetters([]rune{'a', 'b', 'c'}) = 'f'
AddLetters([]rune{'a', 'b'}) = 'c'
AddLetters([]rune{'z'}) = 'z'
AddLetters([]rune{'z', 'a'}) = 'a'
AddLetters([]rune{'y', 'c', 'b'}) = 'd' // notice the letters overflowing
AddLetters([]rune{}) = 'z'
*/
