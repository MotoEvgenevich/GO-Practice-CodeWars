package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	{
		result := Pyramid(0)
		expect := [][]int{}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Pyramid(1)
		expect := [][]int{{1}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Pyramid(2)
		expect := [][]int{{1}, {1, 1}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Pyramid(3)
		expect := [][]int{{1}, {1, 1}, {1, 1, 1}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   It("Testing for 0", func() { Expect(Pyramid(0)).To(Equal([][]int{})) })
   It("Testing for 1", func() { Expect(Pyramid(1)).To(Equal([][]int{{1}})) })
   It("Testing for 2", func() { Expect(Pyramid(2)).To(Equal([][]int{[]int{1}, []int{1, 1}})) })
   It("Testing for 3", func() { Expect(Pyramid(3)).To(Equal([][]int{[]int{1}, []int{1, 1}, []int{1, 1, 1}})) })
*/
