package kata

import "math"

func IsNegativeZero(n float64) bool {
	return n == 0 && math.Copysign(1, n) == -1
}

/*
There exist two zeroes: +0 (or just 0) and -0.

Write a function that returns true if the input number is -0 and false otherwise (True and False for Python).

In JavaScript / TypeScript / Coffeescript the input will be a number.

In Python / Java / C / NASM / Haskell / the input will be a float.
*/
