package kata

func CountRedBeads(n int) int {
	var countBed int
	switch {
	case n < 2:
		return 0
	case n == 2:
		return 2
	case n > 2:
		for i := 1; i < n; i++ {
			countBed += 2
		}
		return countBed
	}

	return 0 // your code here
}

/* Two red beads are placed between every two blue beads.
There are N blue beads.
After looking at the arrangement below work out the number of red beads.

@ @@ @ @@ @ @@ @ @@ @ @@ @

Implement count_red_beads(n) (in PHP count_red_beads($n);
in Java, Javascript, TypeScript, C, C++ countRedBeads(n))
so that it returns the number of red beads.
If there are less than 2 blue beads return 0. */
