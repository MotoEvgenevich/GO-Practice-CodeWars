package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	{
		result := Solve([]int{})
		expect := 0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve([]int{1, 2, 3, 4})
		expect := 7
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve([]int{1, 2, 3, 4, 5, 6})
		expect := 13
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
		expect := 47
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   dotest([]int {}, 0)
   dotest([]int {1,2,3,4},7)
   dotest([]int {1,2,3,4,5,6}, 13)
   dotest([]int {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15},47)
*/
