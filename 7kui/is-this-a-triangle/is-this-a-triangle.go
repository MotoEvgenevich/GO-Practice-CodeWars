package kata

func IsTriangle(a, b, c int) bool {
	if a <= 0 && b <= 0 && c <= 0 {
		return false
	}
	return a+b > c && a+c > b && b+c > a
}

/* Implement a function that accepts 3 integer values a, b, c.
The function should return true if a triangle can be built with
the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).

	Expect(IsTriangle(5, 1, 2)).To(Equal(false))
    Expect(IsTriangle(1, 2, 5)).To(Equal(false))
    Expect(IsTriangle(2, 5, 1)).To(Equal(false))
    Expect(IsTriangle(4, 2, 3)).To(Equal(true))
    Expect(IsTriangle(5, 1, 5)).To(Equal(true))
    Expect(IsTriangle(2, 2, 2)).To(Equal(true))
    Expect(IsTriangle(-1, 2, 3)).To(Equal(false))

*/
