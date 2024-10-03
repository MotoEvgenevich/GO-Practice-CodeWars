package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInAscOrder(t *testing.T) {

	{
		result := InAscOrder([]int{1, 2, 4, 7, 19})
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := InAscOrder([]int{1, 2, 3, 4, 5})
		expect := true
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := InAscOrder([]int{1, 6, 10, 18, 2, 4, 20})
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
	{
		result := InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1})
		expect := false
		assert.Equal(t, result, expect, "result doesn't equal expect")
	}
}

/*
For example:

InAscOrder([]int{1, 2, 4, 7, 19}) // returns True
InAscOrder([]int{1, 2, 3, 4, 5}) // returns True
InAscOrder([]int{1, 6, 10, 18, 2, 4, 20}) // returns False
InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}) // returns False because the numbers are in DESCENDING order
N.B. If your solution passes all fixed tests but fails at the random tests, make sure you aren't mutating the input array.
*/
