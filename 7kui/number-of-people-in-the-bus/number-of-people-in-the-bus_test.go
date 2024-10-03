package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	{
		result := Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}})
		expect := 17
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Number([][2]int{{10, 0}, {3, 5}, {5, 8}})
		expect := 5
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Number([][2]int{{3, 0}, {9, 1}, {4, 8}, {12, 2}, {6, 1}, {7, 8}})
		expect := 21
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
	{
		result := Number([][2]int{{0, 0}})
		expect := 0
		assert.Equal(t, result, expect, "result value doesn't equal expect value")
	}
}

/*

   Expect(Number([][2]int{{3,0},{9,1},{4,8},{12,2},{6,1},{7,8}})).To(Equal(21))
   Expect(Number([][2]int{{0,0}})).To(Equal(0))

*/
