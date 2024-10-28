package main

import "fmt"

func main() {
	fmt.Println(ExpressionMatter(1, 1, 1))
}

func ExpressionMatter(a int, b int, c int) int {
	solOne := a * (b + c)
	solTwo := a * b * c
	solTree := a + b*c
	solFour := (a + b) * c
	solFive := a + b + c
	switch {
	case solOne >= solTwo && solOne >= solTree && solOne >= solFour && solOne >= solFive:
		return solOne
	case solTwo >= solTree && solTwo >= solFour && solTwo >= solOne && solTwo >= solFive:
		return solTwo
	case solTree >= solFour && solTree >= solTwo && solTree >= solOne && solTree >= solFive:
		return solTree
	case solFour >= solTree && solFour >= solTwo && solFour >= solOne && solFour >= solFive:
		return solFour
	case solFive >= solOne && solOne >= solTwo && solFive >= solTree && solOne >= solFour:
		return solFive
	default:
		return 5
	}

}

//solOne  = 1 * (2 + 3) = 5
//solTwo  = 1 * 2 * 3 = 6
//solTree = 1 + 2 * 3 = 7
//solFour = (1 + 2) * 3 = 9
