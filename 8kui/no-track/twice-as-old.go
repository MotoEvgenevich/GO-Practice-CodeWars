package main

import "fmt"

func main() {
	fmt.Println(TwiceAsOld(29, 0))
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	futureSum := 0
	const maxLife int = 100
	for i := 0; i < maxLife; i++ {
		if (dadYearsOld+i)-(sonYearsOld+i) == sonYearsOld+i {
			futureSum = i
		}
	}
	fmt.Println("FutureSum", futureSum)

	pastSum := 0
	for i := 0; i < maxLife; i++ {
		if (dadYearsOld-i)-(sonYearsOld-i) == sonYearsOld-i {
			pastSum = i
		}
	}
	fmt.Println("PastSum:", pastSum)

	if futureSum != 0 {
		return futureSum
	}
	if pastSum != 0 {
		return pastSum
	}
	return 0
}
