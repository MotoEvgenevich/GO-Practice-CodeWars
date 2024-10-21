package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubblesortOnce(t *testing.T) {
	{
		result := BubblesortOnce([]int{9, 7, 5, 3, 1, 2, 4, 6, 8})
		expect := []int{7, 5, 3, 1, 2, 4, 6, 8, 9}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
	Expect(BubblesortOnce([]int{9, 7, 5, 3, 1, 2, 4, 6, 8})).To(Equal([]int{7, 5, 3, 1, 2, 4, 6, 8, 9}))
*/
