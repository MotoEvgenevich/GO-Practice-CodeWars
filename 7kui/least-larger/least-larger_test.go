package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastLarger(t *testing.T) {
	{
		arr := []int{4, 1, 3, 5, 6}
		index := 0
		result := LeastLarger(arr, index)
		expect := 3
		assert.Equal(t, result, expect, "не соотвествует значения result & expect")
	}
	{
		arr := []int{4, 1, 3, 5, 6}
		index := 4
		result := LeastLarger(arr, index)
		expect := -1
		assert.Equal(t, result, expect, "не соотвествует значения result & expect")
	}
}

/*
LeastLarger([]int{4, 1, 3, 5, 6}, 0)  // returns 3
LeastLarger([]int{4, 1, 3, 5, 6}, 4)  // returns -1
*/
