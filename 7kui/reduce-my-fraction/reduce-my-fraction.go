package kata

func ReduceFraction(fraction [2]int) [2]int {
	result := [2]int{}
	a := fraction[0]
	b := fraction[1]
	for b != 0 {
		a, b = b, a%b
	}
	result[0] = fraction[0] / a
	result[1] = fraction[1] / a

	return result
}

/*
Write a function which reduces fractions to their simplest form! Fractions will be presented as an array/tuple
(depending on the language) of strictly positive integers, and the reduced fraction must be returned as an array/tuple:

input:   [numerator, denominator]
output:  [reduced numerator, reduced denominator]
example: [45, 120] --> [3, 8]
All numerators and denominators will be positive integers.

Note: This is an introductory Kata for a series... coming soon!
*/
