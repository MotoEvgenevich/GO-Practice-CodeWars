package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {

	{
		result := Solve([]int{3, 4, 4, 3, 6, 3})
		expect := []int{4, 6, 3}
		assert.Equal(t, result, expect, "value of result doesn't equal expect")
	}
	{
		result := Solve([]int{1, 2, 1, 2, 1, 2, 3})
		expect := []int{1, 2, 3}
		assert.Equal(t, result, expect, "value of result doesn't equal expect")
	}
	{
		result := Solve([]int{1, 2, 3, 4})
		expect := []int{1, 2, 3, 4}
		assert.Equal(t, result, expect, "value of result doesn't equal expect")
	}
	{
		result := Solve([]int{1, 1, 4, 5, 1, 2, 1})
		expect := []int{4, 5, 2, 1}
		assert.Equal(t, result, expect, "value of result doesn't equal expect")
	}
	{
		result := Solve([]int{1, 2, 1, 2, 1, 1, 3})
		expect := []int{2, 1, 3}
		assert.Equal(t, result, expect, "value of result doesn't equal expect")
	}
	{
		result := Solve([]int{0, 4, 4, 3, 0, 3})
		expect := []int{4, 0, 3}
		assert.Equal(t, result, expect, "value of result doesn't equal expect")
	}

}
