package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	{
		result := Transpose([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})
		expect := [][]int{[]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := Transpose([][]int{[]int{1, 2, 3}})
		expect := [][]int{[]int{1}, []int{2}, []int{3}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Transpose([][]int{[]int{1}})
		expect := [][]int{[]int{1}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Transpose([][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{0, 0, 1}, []int{0, 1, 0}, []int{1, 0, 0}})
		expect := [][]int{[]int{1, 0, 0, 0, 1}, []int{0, 1, 0, 1, 0}, []int{0, 0, 1, 0, 0}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
Expect(Transpose([][]int{[]int{1}})).To(Equal([][]int{[]int{1}}))
       Expect(Transpose([][]int{[]int{1, 2, 3}})).To(Equal([][]int{[]int{1}, []int{2}, []int{3}}))
       Expect(Transpose([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})).To(Equal([][]int{[]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9}}))
       Expect(Transpose([][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{0, 0, 1}, []int{0, 1, 0}, []int{1, 0, 0}})).To(Equal([][]int{[]int{1, 0, 0, 0, 1}, []int{0, 1, 0, 1, 0}, []int{0, 0, 1, 0, 0}}))
*/
