package kata

func bandNameGenerator(word string) (result string) {
	var sliceWord []byte
	var exWord string

	for _, v := range word {

		sliceWord = append(sliceWord, byte(v))

	}
	for i, v := range sliceWord {
		if v == 45 {
			sliceWord[i+1] = sliceWord[i+1] - 32
		}
	}
	for _, v := range sliceWord {
		exWord += string(v)
	}
	if exWord[0] != exWord[len(exWord)-1] {
		return "The " + string(exWord[0]-32) + exWord[1:]
	} else {
		return string(exWord[0]-32) + exWord[1:] + exWord[1:]
	}

}

/*
My friend wants a new band name for her band. She like bands that use the formula:
"The" + a noun with the first letter capitalized, for example:

"dolphin" -> "The Dolphin"

However, when a noun STARTS and ENDS with the same letter,
she likes to repeat the noun twice and connect them together with the first and last letter,
combined into one word (WITHOUT "The" in front), like this:

"alaska" -> "Alaskalaska"

Complete the function that takes a noun as a string, and returns her preferred band name written as a string.

INPUT:
The Step-daughter
tart
*/
