package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplicationTable(t *testing.T) {
	{
		result := MultiplicationTable(1)
		expect := [][]int{{1}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := MultiplicationTable(2)
		expect := [][]int{{1, 2}, {2, 4}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := MultiplicationTable(3)
		expect := [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*
	Expect(MultiplicationTable(1)).To(Equal([][]int{{1}}))
	Expect(MultiplicationTable(2)).To(Equal([][]int{{1, 2}, {2, 4}}))
	Expect(MultiplicationTable(3)).To(Equal([][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}))
*/
