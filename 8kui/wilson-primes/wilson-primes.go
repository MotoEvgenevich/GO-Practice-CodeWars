package main

import "fmt"

func main() {
	fmt.Println(AmIWilson(563))
	fmt.Println(AmIWilson(5))
}

func AmIWilson(n int) bool {
	var sum float32
	var chk int
	p := 1

	for i := 1; i < n-1; i++ {

		p = p*i + 1

	}
	fmt.Println("p:", p)
	sum = float32((p + 1)) / float32((n * n))
	chk = int(sum)
	fmt.Println(chk)
	switch {
	case float32(chk) == sum:
		fmt.Println(float32(chk), " ", sum)
		return true
	case float32(chk) != sum:
		fmt.Println(float32(chk), " ", sum)
		return false
	}
	return false
}

/* func factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
} */

/*
Wilson primes satisfy the following condition. Let P represent a prime number.
Then,
((P-1)! + 1) / (P * P)
should give a whole number.
Your task is to create a function that returns true if the given number is a Wilson prime.
563
13
*/
