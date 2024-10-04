package kata

func ScrabbleScore(st string) int {
	score := 0

	for _, v := range st {
		switch {
		case string(v) == "A" || string(v) == "E" || string(v) == "I" || string(v) == "O" || string(v) == "U" || string(v) == "L" || string(v) == "N" || string(v) == "R" || string(v) == "S" || string(v) == "T" || string(v) == "a" || string(v) == "e" || string(v) == "i" || string(v) == "o" || string(v) == "u" || string(v) == "l" || string(v) == "n" || string(v) == "r" || string(v) == "s" || string(v) == "t":
			score += 1
		case string(v) == "D" || string(v) == "G" || string(v) == "d" || string(v) == "g":
			score += 2
		case string(v) == "B" || string(v) == "C" || string(v) == "M" || string(v) == "P" || string(v) == "b" || string(v) == "c" || string(v) == "m" || string(v) == "p":
			score += 3
		case string(v) == "F" || string(v) == "H" || string(v) == "V" || string(v) == "W" || string(v) == "Y" || string(v) == "f" || string(v) == "h" || string(v) == "v" || string(v) == "w" || string(v) == "y":
			score += 4
		case string(v) == "K" || string(v) == "k":
			score += 5
		case string(v) == "J" || string(v) == "X" || string(v) == "j" || string(v) == "x":
			score += 8
		case string(v) == "Q" || string(v) == "Z" || string(v) == "q" || string(v) == "z":
			score += 10

		}

	}

	return score
}

/* Write a program that, given a word, computes the scrabble score for that word.

Letter Values

You'll need these:

Letter                           Value
A, E, I, O, U, L, N, R, S, T       1
D, G                               2
B, C, M, P                         3
F, H, V, W, Y                      4
K                                  5
J, X                               8
Q, Z                               10
There will be a preloaded map DictScores with all these values: DictScores["E"] == 1

Examples

"cabbage" --> 14
"cabbage" should be scored as worth 14 points:

3 points for C
1 point for A, twice
3 points for B, twice
2 points for G
1 point for E
And to total:

3 + 2*1 + 2*3 + 2 + 1 = 3 + 2 + 6 + 3 = 14

Empty string should return 0.
The string can contain spaces and letters (upper and lower case),
you should calculate the scrabble score only of the letters in that string.

""           --> 0
"STREET"    --> 6
"st re et"    --> 6
"ca bba g  e" --> 14 */
