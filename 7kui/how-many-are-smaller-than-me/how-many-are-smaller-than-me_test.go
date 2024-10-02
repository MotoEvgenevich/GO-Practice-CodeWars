package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmaller(t *testing.T) {

	{
		result := Smaller([]int{5, 4, 3, 2, 1})
		expected := []int{4, 3, 2, 1, 0}
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := Smaller([]int{1, 2, 3})
		expected := []int{0, 0, 0}
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := Smaller([]int{1, 2, 0})
		expected := []int{1, 1, 0}
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := Smaller([]int{1, 2, 1})
		expected := []int{0, 1, 0}
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := Smaller([]int{1, 1, -1, 0, 0})
		expected := []int{3, 3, 0, 0, 0}
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := Smaller([]int{5, 4, 7, 9, 2, 4, 4, 5, 6})
		expected := []int{4, 1, 5, 5, 0, 0, 0, 0, 0}
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
}
