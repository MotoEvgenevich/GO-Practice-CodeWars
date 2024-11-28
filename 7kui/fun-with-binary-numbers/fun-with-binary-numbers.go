package kata

import (
	"fmt"
	"math"
	"strconv"
)

func Solution(n, b int) []int {
	if b == 0 {
		return []int{}
	}

	slice := []int{}
	binaryB := strconv.FormatInt(int64(b), 2)
	for i := 1; i < int(math.Pow(2, float64(n))); i++ {
		binaryI := strconv.FormatInt(int64(i), 2)
		if HasSameBits(binaryI, binaryB) {
			slice = append(slice, i)
		}
	}

	return slice
}
func HasSameBits(binaryI, binaryB string) bool {
	diff := len(binaryI) - len(binaryB)
	if diff < 0 {
		binaryI = fmt.Sprintf("%0*s", len(binaryB), binaryI)
	} else if diff > 0 {
		binaryB = fmt.Sprintf("%0*s", len(binaryI), binaryB)
	}

	for i := 0; i < len(binaryB); i++ {
		if binaryB[i] == '1' && binaryI[i] != '1' {
			return false
		}
	}
	return true
}

/*
Your job is to list all of the numbers up to 2 ** n - 1 that contain a 1 at the same places/bits as the binary representation of b.b will be 1,2,4,8,etc.
For example: solution(4,2) would return [2,3,6,7,10,11,14,15].
The binary numbers from 1 to 16 are:
8 4 2 1 (place)
0 0 0 1
0 0 1 0
0 0 1 1
0 1 0 0
0 1 0 1
0 1 1 0
0 1 1 1
1 0 0 0
1 0 0 1
1 0 1 0
1 0 1 1
1 1 0 0
1 1 0 1
1 1 1 0
1 1 1 1
The numbers with a 1 in the 2's place are 2,3,6,7,10,11,14,and 15
Other examples:
solution(0,1) = []
solution(6,8) = [8, 9, 10, 11, 12, 13, 14, 15, 24, 25, 26, 27, 28, 29, 30, 31, 40, 41, 42, 43, 44, 45, 46, 47, 56, 57, 58, 59, 60, 61, 62, 63]
If b = 0,return empty list.

*/
