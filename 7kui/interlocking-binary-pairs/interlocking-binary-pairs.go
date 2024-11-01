package kata

import (
	"strconv"
	"strings"
)

func Interlockable(a, b uint64) bool {
	// Проверяем, есть ли общие единицы
	return (a & b) == 0
}
func InterlockableAsString(a, b uint64) bool {

	binaryA := strconv.FormatUint(a, 2)
	binaryB := strconv.FormatUint(b, 2)

	if len(binaryA) > len(binaryB) {
		binaryB = strings.Repeat("0", len(binaryA)-len(binaryB)) + binaryB
	} else {
		binaryA = strings.Repeat("0", len(binaryB)-len(binaryA)) + binaryA
	}

	positions := make(map[int]bool)

	for i := len(binaryA) - 1; i >= 0; i-- {
		if binaryA[i] == '1' {
			positions[i] = true
		}
	}

	for i := len(binaryB) - 1; i >= 0; i-- {
		if binaryB[i] == '1' {
			if positions[i] {
				return false
			}
		}
	}

	return true
}

/*
Task

Write a function that checks if two non-negative integers make an "interlocking binary pair".

Interlock ?

numbers can be interlocked if their binary representations have no 1's in the same place
comparisons are made by bit position, starting from right to left (see the examples below)
when representations are of different lengths, the unmatched left-most bits are ignored
Examples

a = 9, b = 4

Stacking representations shows how they can interlock.

 9    1001
 4     100
Here, no 1's share any position, so the function returns true.


a = 3, b = 6

These representations do not interlock.

 3      11
 6     110
Finding they both have a 1 in the same position, the function returns false.

Input

Two non-negative integers.

Output

Boolean true or false whether or not these integers are interlockable.
*/
