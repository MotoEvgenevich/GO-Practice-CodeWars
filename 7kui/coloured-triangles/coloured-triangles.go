package kata

func Triangle(row string) rune {
	result := ""

	if len(row) == 1 {
		return rune(row[0])
	}
	for i := 1; i < len(row); i++ {
		switch {
		case row[i-1] == row[i]:
			result += string(row[i])
		case row[i-1] == 'B' && row[i] == 'R':
			result += string('G')
		case row[i-1] == 'R' && row[i] == 'B':
			result += string('G')
		case row[i-1] == 'G' && row[i] == 'B':
			result += string('R')
		case row[i-1] == 'B' && row[i] == 'G':
			result += string('R')
		case row[i-1] == 'G' && row[i] == 'R':
			result += string('B')
		case row[i-1] == 'R' && row[i] == 'G':
			result += string('B')
		case row[i-1] == 'B' && row[i] == 'B':
			result += string('B')
		case row[i-1] == 'G' && row[i] == 'G':
			result += string('G')
		case row[i-1] == 'R' && row[i] == 'R':
			result += string('R')
		}

	}
	row = result

	return Triangle(row)
}

/*
If you finish this kata, you can try Insane Coloured Triangles by Bubbler, which is a much harder version of this one.

A coloured triangle is created from a row of colours, each of which is red, green or blue.
Successive rows, each containing one fewer colour than the last, are generated by considering the two touching colours in the previous row.
If these colours are identical, the same colour is used in the new row. If they are different, the missing colour is used in the new row.
This is continued until the final row, with only a single colour, is generated.

The different possibilities are:

Colour here:        G G        B G        R G        B R
Becomes colour:      G          R          B          G
With a bigger example:

R R G B R G B B
 R B R G B R B
  G G B R G G
   G R G B G
    B B R R
     B G R
      R B
       G
You will be given the first row of the triangle as a string and its your job to return the final colour which would appear in the bottom row as a string. In the case of the example above, you would the given RRGBRGBB you should return G.

The input string will only contain the uppercase letters R, G, B and there will be at least one letter so you do not have to test for invalid input.
If you are only given one colour as the input, return that colour.
*/