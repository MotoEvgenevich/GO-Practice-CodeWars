package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	{
		result := Solve(8, 2)
		expect := []int{2, 6}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve(10, 3)
		expect := []int{-1, -1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve(12, 4)
		expect := []int{4, 8}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve(12, 5)
		expect := []int{-1, -1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve(50, 10)
		expect := []int{10, 40}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve(50, 4)
		expect := []int{-1, -1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve(10, 5)
		expect := []int{5, 5}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Solve(1000, 5)
		expect := []int{5, 995}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   dotest(8,2, []int{2,6})
   dotest(10,3, []int{-1,-1})
   dotest(12,4, []int{4,8})
   dotest(12,5, []int{-1,-1})
   dotest(50,10, []int{10,40})
   dotest(50,4, []int{-1,-1})
   dotest(10,5, []int{5,5})
   dotest(1000,5, []int{5,995})
*/
