package kata

func Solution(mtrx [][]rune) bool {
	for _, line := range mtrx {
		var indexOfX, indexOfArrow int
		foundX, foundArrow := false, false

		// Проходим по строке
		for position, char := range line {
			switch char {
			case 'x':
				indexOfX = position
				foundX = true
			case '>':
				indexOfArrow = position
				foundArrow = true
			}
		}

		// Проверяем, что стрелка и цель есть в строке и стрелка слева от цели
		if foundX && foundArrow && indexOfArrow < indexOfX {
			return true
		}
	}

	// Если стрелка не направлена на цель в ни одной из строк
	return false
}

/*
Hit the target

given a matrix n x n (2-7), determine if the arrow is directed to the target (x).
There will be only 1 arrow '>' and 1 target 'x'
An empty spot will be denoted by a space " ", the target with a cross "x", and the scope ">"
Examples:

given matrix 4x4:
[
  [' ', ' ', ' ', ' '],
  [' ', ' ', ' ', ' '], --> return true
  [' ', '>', ' ', 'x'],
  [' ', ' ', ' ', ' ']
]
given matrix 4x4:
[
  [' ', ' ', ' ', ' '],
  [' ', '>', ' ', ' '], --> return false
  [' ', ' ', ' ', 'x'],
  [' ', ' ', ' ', ' ']
]
given matrix 4x4:
[
  [' ', ' ', ' ', ' '],
  [' ', 'x', '>', ' '], --> return false
  [' ', '', ' ', ' '],
  [' ', ' ', ' ', ' ']
]

In this example, only a 4x4 matrix was used, the problem will have matrices of dimensions from 2 to 7
Happy hacking as they say!

*/
