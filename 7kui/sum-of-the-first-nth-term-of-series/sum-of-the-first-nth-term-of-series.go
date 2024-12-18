package kata

import (
	"fmt"
	"math"
)

func SeriesSum(n int) string {
	if n == 0 {
		return "0.00"
	}
	sum := 0.0
	for i := 0; i < n; i++ {
		divisor := 1 + 3*i
		sum += 1.0 / float64(divisor)
	}
	sum = math.Round(sum*100) / 100
	return fmt.Sprintf("%.2f", sum)
}

/*
Task

Your task is to write a function which returns the n-th term of the following series,
which is the sum of the first n terms of the sequence (n is the input parameter).

Series:1+1/4+1/7+1/10+1/13+1/16+...
Series:1+1/4+1/7+1/10+1/13+1/16+...
You will need to figure out the rule of the series to complete this.

Rules

You need to round the answer to 2 decimal places and return it as String.

If the given value is 0 then it should return "0.00".

You will only be given Natural Numbers as arguments.

Examples (Input --> Output)

n
1 --> 1 --> "1.00"
2 --> 1 + 1/4 --> "1.25"
5 --> 1 + 1/4 + 1/7 + 1/10 + 1/13 --> "1.57"
*/
