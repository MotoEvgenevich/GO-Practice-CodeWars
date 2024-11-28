package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	{
		result := Solution(4, 2)
		expect := []int{2, 3, 6, 7, 10, 11, 14, 15}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solution(6, 8)
		expect := []int{8, 9, 10, 11, 12, 13, 14, 15, 24, 25, 26, 27, 28, 29, 30, 31, 40, 41, 42, 43, 44, 45, 46, 47, 56, 57, 58, 59, 60, 61, 62, 63}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solution(5, 32)
		expect := []int{}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solution(6, 0)
		expect := []int{}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solution(0, 1)
		expect := []int{}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   dotest(4, 2, []int{2, 3, 6, 7, 10, 11, 14, 15})
   dotest(6, 8, []int{8, 9, 10, 11, 12, 13, 14, 15, 24, 25, 26, 27, 28, 29, 30, 31, 40, 41, 42, 43, 44, 45, 46, 47, 56, 57, 58, 59, 60, 61, 62, 63})
   dotest(5, 32, nil)
   dotest(6, 0, nil)
   dotest(0, 1, nil)
*/
