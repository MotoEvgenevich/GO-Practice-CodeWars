package main

import "fmt"

func main() {
	fmt.Println(HowMuchILoveYou(17))
}

func HowMuchILoveYou(i int) string {
	nmb := i
	for i := 0; nmb > 6; i++ {
		if nmb > 6 {
			nmb -= 6
		}
	}
	fmt.Println(nmb)

	switch nmb {
	case 1:
		return "I love you"
	case 2:
		return "a little"
	case 3:
		return "a lot"
	case 4:
		return "passionately"
	case 5:
		return "madly"
	case 6:
		return "not at all"
	default:
		return "wrong number"
	}

}
