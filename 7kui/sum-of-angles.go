/* Find the total sum of internal angles
(in degrees) in an n-sided simple polygon.
N will be greater than 2. */

package main

import "fmt"

func main() {
	fmt.Println(Angle(17))
}

func Angle(n int) int {

	return 180 * (n - 2)
}
