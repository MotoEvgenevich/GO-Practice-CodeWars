package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMirror(t *testing.T) {
	{
		result := Mirror([]int{})
		expect := []int{}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Mirror([]int{1})
		expect := []int{1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Mirror([]int{2, 1})
		expect := []int{1, 2, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Mirror([]int{1, 3, 2})
		expect := []int{1, 2, 3, 2, 1}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Mirror([]int{-8, 42, 18, 0, -16})
		expect := []int{-16, -8, 0, 18, 42, 18, 0, -8, -16}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
   It("should test that the solution returns the correct value", func() {
       Expect(Mirror([]int{})).To(Equal([]int{}))
       Expect(Mirror([]int{1})).To(Equal([]int{1}))
       Expect(Mirror([]int{2, 1})).To(Equal([]int{1, 2, 1}))
       Expect(Mirror([]int{1, 3, 2})).To(Equal([]int{1, 2, 3, 2, 1}))
       Expect(Mirror([]int{-8, 42, 18, 0, -16})).To(Equal([]int{-16, -8, 0, 18, 42, 18, 0, -8, -16}))
       Expect(Mirror([]int{-5, 10, 8, 10, 2, -3, 10})).To(Equal([]int{-5, -3, 2, 8, 10, 10, 10, 10, 10, 8, 2, -3, -5}))
   })
*/
