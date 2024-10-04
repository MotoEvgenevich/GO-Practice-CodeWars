package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementer(t *testing.T) {

	{
		result := Incrementer([]int{1, 2, 3})
		var expect []int = []int{2, 4, 6}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := Incrementer([]int{4, 6, 7, 1, 3})
		var expect []int = []int{5, 8, 0, 5, 8}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := Incrementer([]int{3, 6, 9, 8, 9})
		var expect []int = []int{4, 8, 2, 2, 4}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}

	{
		result := Incrementer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 9, 8})
		var expect []int = []int{2, 4, 6, 8, 0, 2, 4, 6, 8, 9, 0, 1, 2, 2}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
  It("Sample tests", func() {
       Expect(Incrementer([]int{})).To(BeEmpty())
       Expect(Incrementer([]int{1, 2, 3})).To(Equal([]int{2, 4, 6}))
       Expect(Incrementer([]int{4, 6, 7, 1, 3})).To(Equal([]int{5, 8, 0, 5, 8}))
       Expect(Incrementer([]int{3, 6, 9, 8, 9})).To(Equal([]int{4, 8, 2, 2, 4}))
       Expect(Incrementer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9, 9, 8})).To(Equal([]int{2, 4, 6, 8, 0, 2, 4, 6, 8, 9, 0, 1, 2, 2}))
     })
*/
