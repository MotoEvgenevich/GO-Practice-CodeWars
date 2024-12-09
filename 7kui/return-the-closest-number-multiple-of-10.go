/* Given a number return the closest number to it that is divisible by 10.

Example input:

22
25
37
Expected output:

20
30
40 */

package main

import "fmt"

func main() {
	fmt.Println(ClosestMultipleOf10(37))
}

func ClosestMultipleOf10(n uint32) uint32 {
	switch {
	case (n-1)%10 == 0:
		return n - 1
	case (n-2)%10 == 0:
		return n - 2
	case (n-3)%10 == 0:
		return n - 3
	case (n-4)%10 == 0:
		return n - 4
	case (n+5)%10 == 0:
		return n + 5
	case (n+4)%10 == 0:
		return n + 4
	case (n+3)%10 == 0:
		return n + 3
	case (n+2)%10 == 0:
		return n + 2
	case (n+1)%10 == 0:
		return n + 1
	}
	return n
}

func ClosestMultipleOf10v2(n uint32) uint32 {
	remainder := n % 10
	if remainder >= 5 {
		return n + (10 - remainder)
	}
	return n - remainder
}
