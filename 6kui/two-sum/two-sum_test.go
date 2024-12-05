package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	{
		result := TwoSum([]int{1, 2, 3}, 4)
		expect := [2]int{0, 2}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := TwoSum([]int{1234, 5678, 9012}, 14690)
		expect := [2]int{1, 2}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
	{
		result := TwoSum([]int{2, 2, 3}, 4)
		expect := [2]int{0, 1}
		assert.Equal(t, expect, result, "result doesn't equal expect")
	}
}

/*
   Expect(TwoSum([]int{1, 2, 3}, 4)).To(Equal([2]int{0, 2}))
   Expect(TwoSum([]int{1234, 5678, 9012}, 14690)).To(Equal([2]int{1, 2}))
   Expect(TwoSum([]int{2, 2, 3}, 4)).To(Equal([2]int{0, 1}))
*/
