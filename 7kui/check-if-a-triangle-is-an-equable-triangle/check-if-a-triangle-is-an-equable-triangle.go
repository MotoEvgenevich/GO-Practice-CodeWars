package kata

func EquableTriangle(a, b, c int) bool {
	perimeter := a + b + c

	area := 0
	p := (a + b + c) / 2

	sBfSquare := p * (p - a) * (p - b) * (p - c)

	for i := 1; i <= sBfSquare; i++ {
		area = i * i
		if area == sBfSquare {
			area = i

			break
		}

	}

	if area == perimeter {
		return true
	} else {
		return false
	}
}

/* A triangle is called an equable triangle if its area equals its perimeter.
Return true, if it is an equable triangle, else return false.
You will be provided with the length of sides of the triangle. Happy Coding!
*/
