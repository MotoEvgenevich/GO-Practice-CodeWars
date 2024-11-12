package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoOldestAges(t *testing.T) {
	{
		result := TwoOldestAges([]int{1, 2, 10, 8})
		expected := [2]int{8, 10}
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
	{
		result := TwoOldestAges([]int{93, 35, 53, 67, 17, 23, 89, 75, 15, 53})
		expected := [2]int{89, 93}
		assert.Equal(t, result, expected, "value of result doesn't equal expected")
	}
}

/*
should return 89 and 93 for input []int{93,35,53,67,17,23,89,75,15,53}
*/
